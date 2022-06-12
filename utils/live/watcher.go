package live

import (
	"github.com/fsnotify/fsnotify"
	"path/filepath"
	. "timenoteWeb/utils/config"
)

func WatchBackupDataPath() (watcher *fsnotify.Watcher) {
	var log = logging.WithField("源", "WatchBackupDataPath")
	dataPath := filepath.Join(AppConfig.Data.Root, AppConfig.Data.Dir)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.WithError(err).Warn("备份文件监听错误")
		return nil
	}
	err = watcher.Add(dataPath)
	if err != nil {
		log.WithError(err).Warn("备份文件监听错误")
		return nil
	}
	go func() {
		for {
			select {
			case ev := <-watcher.Events:
				{
					if ev.Op&fsnotify.Create == fsnotify.Create {
						//log.WithField("文件", ev.Name).Info("创建文件")
						continue
					}
					if ev.Op&fsnotify.Write == fsnotify.Write {
						log.WithField("文件", ev.Name).Info("写入文件")
					}
					if ev.Op&fsnotify.Remove == fsnotify.Remove {
						log.WithField("文件", ev.Name).Info("删除文件")
					}
					if ev.Op&fsnotify.Rename == fsnotify.Rename {
						log.WithField("文件", ev.Name).Info("重命名文件")
					}
					if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
						log.WithField("文件", ev.Name).Info("修改权限")
					}
					exportNotes()
				}
			case err := <-watcher.Errors:
				{
					log.WithError(err).Warn("监听异常")
				}
			}
		}
	}()
	return
}
