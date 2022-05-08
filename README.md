# 时光记 Web

*最多是个查看器，没有什么特别的功能。*  
**严格地说**, 这是一个自托管的面向 时光记 的 WebDav 服务端程序。

## Features

* [x] WebDav 支持
* [x] 自定义数据目录 (你甚至可以使用挂载的 OneDrive 等其他网盘的目录)
* [x] WebDav 登录验证
* [x] WebUI (有且仅有日记数量统计)
* [ ] 查看日记内容
* [ ] ~~修改日记内容~~
* [ ] 根据分类查看日记
* [ ] 查看 Todo
* [ ] 导出日记为纯 Markdown 文件
* [ ] 其他乱七八糟的

## Install

*暂无 release*

### 自行编译

```shell
# 克隆 repo
git clone https://github.com/gggxbbb/timenoteWeb
cd timenoteWeb

# 确保安装 golang 1.18+
# 拉取依赖
go mod tidy
# 编译
go build
```

## Config

程序在运行时会自动创建配置文件，以下为配置示例:

```yaml
admin:
  password: admin123456 # 管理员账号
  username: admin # 管理员密码
dav:
  data_path: ./data # WebDav 工作目录
server:
  debug: false # 是否为调试模式
  listen: 0.0.0.0 # 监听地址
  port: 8080 # 监听端口
web:
  nickname: timenoteUser # WebUI 显示的用户名
  title: timenoteWeb # 暂时啥用没有
```

**注意**, 对于使用 WebDav, 时光记文件存储于 `WebDav根目录/timeNote`, 但使用其他存储方式是不一定如此。