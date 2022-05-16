package web

import (
	"github.com/gin-gonic/gin"
	"sort"
	"strconv"
	"time"
	. "timenoteWeb/config"
	"timenoteWeb/loader"
)

func TimelinePage(c *gin.Context) {

	timenoteData, success := loader.LoadLastDataFile()
	if !success {
		var data errorPageData
		data.Title = "时间线"
		data.Nickname = AppConfig.Web.Nickname
		data.Error = errNoDataFile
		c.HTML(errNoDataFile.Code, "error.html", data)
		return
	}
	tempData := make(map[int64]timelineEvent)

	for _, v := range timenoteData.Notes {
		t := time.Unix(v.ID/1000, 0)
		tempData[v.ID] = timelineEvent{
			Year:  t.Year(),
			Month: int(t.Month()),
			Day:   t.Day(),
			ID:    v.ID,
			Title: "日记: " + v.Title,
			Url:   "/notes/" + strconv.FormatInt(v.ID, 10),
		}
	}
	for _, v := range timenoteData.Todos {
		t := time.Unix(v.ID/1000, 0)
		tempData[v.ID] = timelineEvent{
			Year:  t.Year(),
			Month: int(t.Month()),
			Day:   t.Day(),
			ID:    v.ID,
			Title: "Todo: " + v.Title,
			Url:   "/todos/" + strconv.FormatInt(v.ID, 10),
		}
	}
	for _, v := range timenoteData.Categories {
		if v.ID == -1 {
			continue
		}
		t := time.Unix(v.ID/1000, 0)
		tempData[v.ID] = timelineEvent{
			Year:  t.Year(),
			Month: int(t.Month()),
			Day:   t.Day(),
			ID:    v.ID,
			Title: "分类: " + v.CategoryName,
			Url:   "/categories/" + strconv.FormatInt(v.ID, 10),
		}
	}

	sortedData := make(map[int]map[int]map[int][]timelineEvent)
	for k := range tempData {
		d := tempData[k]
		if sortedData[d.Year] == nil {
			sortedData[d.Year] = make(map[int]map[int][]timelineEvent)
		}
		if sortedData[d.Year][d.Month] == nil {
			sortedData[d.Year][d.Month] = make(map[int][]timelineEvent)
		}
		if sortedData[d.Year][d.Month][d.Day] == nil {
			sortedData[d.Year][d.Month][d.Day] = make([]timelineEvent, 0)
		}
		sortedData[d.Year][d.Month][d.Day] = append(sortedData[d.Year][d.Month][d.Day], d)
	}

	var timelineData []timelineYear

	var yearKeys []int
	for k := range sortedData {
		yearKeys = append(yearKeys, k)
	}
	sort.Slice(yearKeys, func(i, j int) bool {
		return yearKeys[i] > yearKeys[j]
	})

	for _, year := range yearKeys {
		months := sortedData[year]
		var monthKeys []int
		for k := range months {
			monthKeys = append(monthKeys, k)
		}
		sort.Slice(monthKeys, func(i, j int) bool {
			return monthKeys[i] > monthKeys[j]
		})
		var monthData []timelineMonth
		for _, month := range monthKeys {
			days := months[month]
			var dayKeys []int
			for k := range days {
				dayKeys = append(dayKeys, k)
			}
			sort.Slice(dayKeys, func(i, j int) bool {
				return dayKeys[i] > dayKeys[j]
			})
			var dayData []timelineDay
			for _, day := range dayKeys {
				events := days[day]
				sort.Slice(events, func(i, j int) bool {
					return events[i].ID > events[j].ID
				})
				dayData = append(dayData, timelineDay{
					Year:   year,
					Month:  month,
					Day:    day,
					Events: events,
				})
			}
			monthData = append(monthData, timelineMonth{
				Year:  year,
				Month: month,
				Days:  dayData,
			})
		}
		timelineData = append(timelineData, timelineYear{
			Year:   year,
			Months: monthData,
		})
	}

	var data timelinePageData
	data.Title = "时间线"
	data.Nickname = AppConfig.Web.Nickname
	data.Years = timelineData
	data.Source = timenoteData.Source

	c.HTML(200, "timeline.html", data)
}
