# Go 二维码扫描工具

这是一个用Go语言编写的命令行二维码扫描工具，可以从图片文件中读取并解析二维码内容。

## 功能特性

- 支持多种图片格式（PNG、JPEG、GIF、BMP、TIFF等）
- 命令行界面，易于使用
- 详细输出模式，提供更多信息
- 错误处理和用户友好的提示信息

## 安装依赖

```bash
go mod init qrcodescanner
go get github.com/tuotoo/qrcode
```

## 编译

```bash
go build -o qrcodescanner main.go
```

## 使用方法

### 基本用法

```bash
# 使用Go运行
go run main.go -file <图片文件路径>

# 使用编译后的可执行文件
./qrcodescanner -file <图片文件路径>
```

### 详细输出模式

```bash
go run main.go -file <图片文件路径> -v
```

### 示例

```bash
# 扫描PNG格式的二维码
go run main.go -file qrcode.png

# 扫描JPEG格式的二维码并显示详细信息
go run main.go -file qrcode.jpg -v
```

## 参数说明

- `-file string`: 要扫描的二维码图片文件路径（必需）
- `-v`: 详细输出模式（可选）

## 支持的图片格式

- PNG (.png)
- JPEG (.jpg, .jpeg)
- GIF (.gif)
- BMP (.bmp)
- TIFF (.tiff, .tif)

## 输出示例

### 普通模式
```
扫描结果:
文件: qrcode.png
内容: https://www.example.com
```

### 详细模式
```
正在扫描文件: /full/path/to/qrcode.png
二维码维度: 29x29
扫描结果:
文件: qrcode.png
内容: https://www.example.com
```

## 错误处理

工具会检查以下常见错误并提供相应的错误信息：

- 文件不存在
- 文件无法读取
- 二维码解码失败
- 不支持的文件格式

## 注意事项

1. 确保图片文件存在且可读
2. 确保图片中的二维码清晰可见
3. 支持多个二维码的图片，但只返回第一个扫描到的结果
4. 图片质量会影响扫描成功率

## 技术实现

- 使用 `github.com/tuotoo/qrcode` 库进行二维码解码
- 使用Go标准库进行文件操作和命令行参数解析
- 支持通过 `io.Reader` 接口读取各种格式的图片文件

## 许可证

本项目采用MIT许可证。
