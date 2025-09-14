# Go 二维码扫描工具

这是一个用Go语言编写的命令行二维码扫描工具，可以从图片文件中读取并解析二维码内容。支持单文件扫描和批量扫描，具有彩色终端输出和文件保存功能。

## 功能特性

- ✅ 支持多种图片格式（PNG、JPEG、GIF、BMP、TIFF等）
- ✅ 单文件扫描模式
- ✅ 批量扫描模式（扫描整个目录）
- ✅ 彩色终端输出，提升用户体验
- ✅ 自动保存扫描结果到output.txt文件
- ✅ 详细的错误处理和用户友好的提示信息
- ✅ 扫描结果包含文件名信息
- ✅ 程序启动时自动清理旧的output.txt文件

## 安装依赖

```bash
go mod init qrcodescanner
go get github.com/tuotoo/qrcode
go get github.com/fatih/color
```

## 编译

```bash
go build -o qrcodescanner main.go
```

## 使用方法

### 单文件扫描

```bash
# 使用Go运行
go run main.go -file <图片文件路径>

# 使用编译后的可执行文件
./qrcodescanner -file <图片文件路径>

# 启用详细输出模式
go run main.go -file <图片文件路径> -v
```

### 批量扫描

```bash
# 扫描整个目录
go run main.go -dir <目录路径>

# 批量扫描并显示详细信息
go run main.go -dir <目录路径> -v
```

### 示例

```bash
# 单文件扫描
go run main.go -file qrcode.png
go run main.go -file qrcode.jpg -v

# 批量扫描
go run main.go -dir ./images
go run main.go -dir ./qrcodes -v
```

## 参数说明

- `-file string`: 要扫描的二维码图片文件路径（单文件模式）
- `-dir string`: 要扫描的目录路径（批量扫描模式）
- `-v`: 详细输出模式（可选）

## 支持的图片格式

- PNG (.png)
- JPEG (.jpg, .jpeg)
- GIF (.gif)
- BMP (.bmp)
- TIFF (.tiff, .tif)

## 输出特性

### 彩色终端输出

工具使用不同颜色来区分不同类型的信息：
- 🔵 青色：标题和摘要信息
- 🟡 黄色：文件名和路径信息
- 🟢 绿色：成功状态和二维码内容
- 🔴 红色：错误信息和失败状态
- ⚪ 白色：分隔符和次要信息
- 🟣 紫色：重要分隔线

### 文件输出

- 每次运行程序时自动删除已存在的output.txt文件
- 扫描结果以无颜色格式保存到output.txt文件
- 文件内容与终端显示内容一致，但不包含颜色代码
- 支持单文件和批量扫描结果的保存

## 输出示例

### 单文件扫描 - 终端输出（彩色）
```
扫描结果:
文件: qrcode.png
内容: https://www.example.com
----------------------------------------
```

### 批量扫描 - 终端输出（彩色）
```
批量扫描结果:
总共扫描文件数: 5
========================================
文件: ./images/qrcode1.png
状态: 成功
内容: https://example.com/page1
----------------------------------------
文件: ./images/qrcode2.jpg
状态: 成功
内容: Hello World!
----------------------------------------
文件: ./images/invalid.png
状态: 失败
错误: 解码二维码失败: cannot find qr code
----------------------------------------
扫描完成: 成功 2 个，失败 1 个
```

### output.txt文件内容（无颜色）
```
批量扫描结果:
总共扫描文件数: 5
========================================
文件: ./images/qrcode1.png
状态: 成功
内容: https://example.com/page1
----------------------------------------
文件: ./images/qrcode2.jpg
状态: 成功
内容: Hello World!
----------------------------------------
文件: ./images/invalid.png
状态: 失败
错误: 解码二维码失败: cannot find qr code
----------------------------------------
扫描完成: 成功 2 个，失败 1 个
```

## 错误处理

工具会检查以下常见错误并提供相应的错误信息：

- 文件不存在
- 文件无法读取
- 二维码解码失败
- 不支持的文件格式
- 目录不存在
- 文件写入失败（output.txt）

## 注意事项

1. 确保图片文件存在且可读
2. 确保图片中的二维码清晰可见
3. 支持多个二维码的图片，但只返回第一个扫描到的结果
4. 批量扫描会自动跳过不支持的文件格式
5. 每次运行程序时会自动删除旧的output.txt文件
6. 扫描结果会包含完整的文件名信息
7. 图片质量会影响扫描成功率
8. 终端颜色输出可能在某些环境中不可用

## 技术实现

- 使用 `github.com/tuotoo/qrcode` 库进行二维码解码
- 使用 `github.com/fatih/color` 库实现彩色终端输出
- 使用Go标准库进行文件操作和命令行参数解析
- 支持通过 `io.Reader` 接口读取各种格式的图片文件
- 使用 `strings.Builder` 高效构建文件输出内容

## 项目结构

```
批量二维码扫描工具/
├── main.go          # 主程序文件
├── go.mod           # Go模块文件
├── go.sum           # 依赖校验文件
├── README.md        # 说明文档
├── qrcodescanner    # 编译后的可执行文件
└── output.txt       # 扫描结果输出文件（运行时生成）
```

## 许可证

本项目采用MIT许可证。
