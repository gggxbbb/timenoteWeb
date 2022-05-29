package config

// ServerConfig 服务配置
type ServerConfig struct {
	// Listen 监听地址
	Listen string `json:"listen" mapstructure:"listen"`

	// Port 监听端口
	Port int `json:"port" mapstructure:"port"`

	// Debug 是否开启调试模式
	Debug bool `json:"debug" mapstructure:"debug"`

	// EnableWebDav 是否启用内建 WebDav 服务
	EnableWebDav bool `json:"enable_webdav" mapstructure:"enable_webdav"`
}

// DataConfig 数据存储配置
type DataConfig struct {
	// Root 数据根目录, WebDav 工作目录
	Root string `json:"root" mapstructure:"root"`

	// Dir 数据目录
	//
	// 对于使用内建 WebDav 进行备份, 保留默认值 "/timeNote/" 即可
	//
	// 对于使用 OneDrive 进行备份, 应修改为 "/应用/记时光/"
	Dir string `json:"dir" mapstructure:"dir"`
}

// AdminConfig 管理员配置, 唯一的用户
type AdminConfig struct {
	// Username 用户名
	Username string `json:"username" mapstructure:"username"`

	// Password 密码
	Password string `json:"password" mapstructure:"password"`
}

// WebConfig 网页配置
type WebConfig struct {
	// Nickname 网页显示的昵称
	Nickname string `json:"nickname" mapstructure:"nickname"`

	// Title 网页标题 (但似乎没有地方用到了这个值?)
	Title string `json:"title" mapstructure:"title"`
}

// MapConfig 地图服务配置
type MapConfig struct {
	// TokenApi 天地图 服务端 密钥
	TokenApi string `json:"token_api" mapstructure:"token_api"`

	// TokenWeb 天地图 网页端 密钥
	TokenWeb string `json:"token_web" mapstructure:"token_web"`
}

// Config 总配置
type Config struct {
	// Server 服务配置
	Server ServerConfig `json:"server" mapstructure:"server"`

	// Data 数据存储配置
	Data DataConfig `json:"data" mapstructure:"data"`

	// Admin 管理员配置
	Admin AdminConfig `json:"admin" mapstructure:"admin"`

	// Web 网页配置
	Web WebConfig `json:"web" mapstructure:"web"`

	// Map 地图服务配置
	Map MapConfig `json:"map" mapstructure:"map"`
}
