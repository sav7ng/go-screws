![](https://file.cdn.tanchi.shop/dev/1/system/61c60b8a0c1602ad6619b6cc52335bd2)

# Go 螺丝钉

go-screws `/ɡəʊˈskruːs/`

`go-screws` 是一个多功能的 Go 语言工具包，提供了一系列实用的工具和辅助函数，旨在简化你的 Go 开发流程。就像一盒螺丝一样，这个包涵盖了多种工具，可以组合使用，高效地解决常见开发任务。无论你是在构建小型脚本还是大规模应用，`go-screws` 都能为你提供合适的工具。

主要特性包括：

    提供多种字符串操作、文件处理等常用工具函数。
    针对常见 Go 任务进行性能优化。
    模块化设计，易于集成，提升开发效率。
    持续扩展的 Go 工具集合，助你应对不同开发场景。

立即开始使用 go-screws，让你的 Go 代码更加严谨高效！

---

`go-screws` is a versatile Go library that provides a collection of essential utilities and helper functions, designed to streamline your Go development. Like a box of screws, this package offers various tools that can be combined and reused to solve common tasks efficiently. Whether you're building a small script or a large-scale application, `go-screws` equips you with the right tools for the job.

Features include:

    A variety of utility functions for string manipulation, file handling, and more.
    Performance optimizations for common Go tasks.
    Modular and easy-to-integrate tools that enhance productivity.
    Continuously growing collection of useful Go utilities.

Get started with go-screws to tighten up your Go code today!

## Install & Update Version

```
go get -u github.com/sav7ng/go-screws.git
```

## Usage

```
import goscrews "github.com/sav7ng/go-screws.git"

base64, err := goscrews.QRCodeTool.GenerateQrcodeToBytes("saving", 512)
```

## Tools

| 工具名称          | 描述         |
|:--------------|:-----------|
| AesTool       | AES加密工具    |
| ArrayTool     | 数组切片工具     |
| CatchTool     | 捕抓工具       |
| ConvertTool   | 转换器工具      |
| EmojiTool     | emoji表情包工具 |
| HttpTool      | http请求工具   |
| IPTool        | IP工具       |
| QRCodeTool    | 二维码工具      |
| RandomTool    | 随机数工具      |
| TimeTool      | 时间工具       |
| ValidatorTool | 验证器工具      |
| geo           | 地理信息工具     |