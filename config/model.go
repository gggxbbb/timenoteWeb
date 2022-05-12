package config

type ServerConfig struct {
	Listen string `json:"listen" mapstructure:"listen"`
	Port   int    `json:"port" mapstructure:"port"`
	Debug  bool   `json:"debug" mapstructure:"debug"`
}

type DavConfig struct {
	DataPath string `json:"dataPath" mapstructure:"data_path"`
}

type DataConfig struct {
	Dir string `json:"dir" mapstructure:"dir"`
}

type AdminConfig struct {
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
}

type WebConfig struct {
	Nickname string `json:"nickname" mapstructure:"nickname"`
	Title    string `json:"title" mapstructure:"title"`
}

type MapConfig struct {
	TokenApi string `json:"token_api" mapstructure:"token_api"`
	TokenWeb string `json:"token_web" mapstructure:"token_web"`
}

type Config struct {
	Server ServerConfig `json:"server" mapstructure:"server"`
	Dav    DavConfig    `json:"dav" mapstructure:"dav"`
	Data   DataConfig   `json:"data" mapstructure:"data"`
	Admin  AdminConfig  `json:"admin" mapstructure:"admin"`
	Web    WebConfig    `json:"web" mapstructure:"web"`
	Map    MapConfig    `json:"map" mapstructure:"map"`
}
