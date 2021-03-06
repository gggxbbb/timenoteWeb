# 记时光 Web

[![Build](https://github.com/gggxbbb/timenoteWeb/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/gggxbbb/timenoteWeb/actions/workflows/build.yml) | [![Release](https://github.com/gggxbbb/timenoteWeb/actions/workflows/release.yml/badge.svg)](https://github.com/gggxbbb/timenoteWeb/actions/workflows/release.yml)  

*最多是个查看器，没有什么特别的功能。*  
**严格地说**, 这是一个自托管的面向 记时光 的 WebDav 服务端程序。

## Features

* [x] WebDav 支持
* [x] 自定义数据目录 (你甚至可以使用挂载的 OneDrive 等其他网盘的目录)
* [x] WebDav 登录验证
* [x] WebUI (有且仅有日记数量统计)
* [x] 查看日记内容
* [ ] ~~修改日记内容~~
* [x] 根据分类查看日记
* [ ] 查看 Todo
* [x] 导出日记为纯 Markdown 文件
  * 目前仅支持自动导出全部
  * 后续可能会支持将修改同步至备份文件
* [ ] 其他乱七八糟的

## Install

### Release

前往 [Release](https://github.com/gggxbbb/timenoteWeb/releases) 下载

### Dev release

前往 [自动化 Build](https://github.com/gggxbbb/timenoteWeb/actions/workflows/build.yml) 下载

### 自行编译

```shell
# 克隆 repo
git clone https://github.com/gggxbbb/timenoteWeb --recurse-submodules
cd timenoteWeb

# 确保安装 golang 1.18
# 拉取依赖
go mod tidy
# 编译
go build
```

## Config

程序在运行时会自动创建配置文件，以下为配置示例:

```yaml
server:
  debug: false # 是否为调试模式
  enable_webdav: true # 是否启用 WebDav
  listen: 0.0.0.0 # 监听地址
  port: 8080 # 监听端口
admin:
  password: admin123456 # 管理员账号
  username: admin # 管理员密码
web:
  nickname: timenoteUser # WebUI 显示的用户名
  title: timenoteWeb # 暂时啥用没有
data:
  root: ./data # 根数据目录, 也即 WebDav 工作目录
  dir: /timeNote/ # 根数据目录下存放记时光备份文件的文件夹
map:
  token_api: "" #天地图 服务器端 密钥
  token_web: "" #天地图 浏览器端 密钥
live:
  data_dir: /timenoteDoc/ # 跟数据目录下存放 markdown 文件的文件夹
  enable: true #是否启用自动导出 markdown
```

**注意**, 对于使用 WebDav, 记时光文件存储于 `WebDav根目录/timeNote`, 但使用其他存储方式是不一定如此。

### 如果想和 OneDrive 备份配合使用

1. 将 `data -> root` 设置为 OneDrive 在你本地的路径, 如 `C:/Users/gggxbbb/OneDrive/`
2. 将 `data -> dir` 设置为 `/应用/记时光/`
3. (可选) 将 `server -> enable_webdav` 设置为 `false`

逻辑上程序将正常读取数据。  
**注意**, 由于 WebDav 备份和 OneDrive 备份数据存储路径不同, 记时光App 中仍应使用 OneDrive 模式进行备份。  
因此此时内置 WebDav 服务将失去作用。并且此时可通过 WebDav 直接访问你 OneDrive 中的一切数据。为了安全，建议禁用内置 WebDav 服务。
