# GoHttp

A http test tool.

## Feature
- built on golang,full platform support
- text-base data file,for vcs. with easy use gui
- environment , variable support
- easy use for authentication
- script support,do something before and after request

## Getting Start

- run locally
go run .

# TODO 
[ ] status code display
[ ] tls data and certificate (response has data)
[ ] resposne view mode (json highlight etc.)
[ ] request but save response to file
[ ] after modify , before close , prompt if save the http file
[ ] about page
[ ] 添加保存快捷键
[ ] 环境变量列表管理
[ ] 更换日志框架
## Roadmap

- Prototype version
implement basic feature
- Standard version
refactor code
- ssl debug
- support command line parameter ,e.g. start dir
# Stack

- golang
- fyne 2
- https://pkg.go.dev/github.com/patrickmn/go-cache#section-readme
- https://github.com/valyala/fasttemplate