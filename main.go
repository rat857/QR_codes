package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/tuotoo/qrcode"
)

func main() {
	// 初始化颜色输出
	color.NoColor = false

	// 删除已存在的output.txt文件
	os.Remove("output.txt")

	// 解析命令行参数
	filePath := flag.String("file", "", "要扫描的二维码图片文件路径")
	dirPath := flag.String("dir", "", "要扫描的目录路径（批量扫描）")
	verbose := flag.Bool("v", false, "详细输出模式")
	flag.Parse()

	// 检查是否提供了文件路径或目录路径
	if *filePath == "" && *dirPath == "" {
		printUsage()
		os.Exit(1)
	}

	// 单文件扫描模式
	if *filePath != "" {
		result, err := scanQRCode(*filePath, *verbose)
		if err != nil {
			log.Fatalf("扫描二维码失败: %v", err)
		}
		printSingleResult(*filePath, result)
		return
	}

	// 批量扫描模式
	if *dirPath != "" {
		results, err := scanDirectory(*dirPath, *verbose)
		if err != nil {
			log.Fatalf("批量扫描失败: %v", err)
		}
		printBatchResults(results)
		return
	}
}

// scanQRCode 扫描二维码并返回内容
func scanQRCode(filePath string, verbose bool) (string, error) {
	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", fmt.Errorf("文件不存在: %s", filePath)
	}

	// 获取文件的绝对路径
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return "", fmt.Errorf("获取文件绝对路径失败: %v", err)
	}

	if verbose {
		fmt.Printf("正在扫描文件: %s\n", absPath)
	}

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("打开文件失败: %v", err)
	}
	defer file.Close()

	// 读取二维码文件
	matrix, err := qrcode.Decode(file)
	if err != nil {
		return "", fmt.Errorf("解码二维码失败: %v", err)
	}

	if verbose {
		fmt.Printf("二维码维度: %dx%d\n", matrix.Size.Dx(), matrix.Size.Dy())
	}

	return matrix.Content, nil
}

// printUsage 打印使用说明
func printUsage() {
	fmt.Println("Go 二维码扫描工具")
	fmt.Println("")
	fmt.Println("使用方法:")
	fmt.Println("  单文件扫描: go run main.go -file <图片文件路径> [-v]")
	fmt.Println("  批量扫描:   go run main.go -dir <目录路径> [-v]")
	fmt.Println("")
	fmt.Println("参数:")
	fmt.Println("  -file string")
	fmt.Println("        要扫描的二维码图片文件路径 (单文件模式)")
	fmt.Println("  -dir string")
	fmt.Println("        要扫描的目录路径 (批量扫描模式)")
	fmt.Println("  -v")
	fmt.Println("        详细输出模式 (可选)")
	fmt.Println("")
	fmt.Println("示例:")
	fmt.Println("  单文件扫描:")
	fmt.Println("    go run main.go -file qrcode.png")
	fmt.Println("    go run main.go -file qrcode.jpg -v")
	fmt.Println("  批量扫描:")
	fmt.Println("    go run main.go -dir ./images")
	fmt.Println("    go run main.go -dir ./qrcodes -v")
	fmt.Println("")
	fmt.Println("支持的图片格式:")
	fmt.Println("  PNG, JPEG, GIF, BMP, TIFF 等常见格式")
	fmt.Println("")
	fmt.Println("注意事项:")
	fmt.Println("  1. 确保图片文件存在且可读")
	fmt.Println("  2. 确保图片中的二维码清晰可见")
	fmt.Println("  3. 支持多个二维码的图片，但只返回第一个扫描到的结果")
	fmt.Println("  4. 批量扫描会自动跳过不支持的文件格式")
	fmt.Println("  5. 扫描结果会包含文件名信息")
}

// isSupportedFormat 检查文件格式是否受支持
func isSupportedFormat(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	supportedFormats := []string{".png", ".jpg", ".jpeg", ".gif", ".bmp", ".tiff", ".tif"}

	for _, format := range supportedFormats {
		if ext == format {
			return true
		}
	}
	return false
}

// printSingleResult 打印单个文件的扫描结果
func printSingleResult(filePath, content string) {
	// 彩色终端输出
	color.Cyan("扫描结果:")
	color.Yellow("文件: %s", filePath)
	color.Green("内容: %s", content)
	color.White("----------------------------------------")

	// 保存到文件
	saveToFile(fmt.Sprintf("扫描结果:\n文件: %s\n内容: %s\n----------------------------------------\n", filePath, content))
}

// scanDirectory 扫描目录中的所有图片文件
func scanDirectory(dirPath string, verbose bool) ([]ScanResult, error) {
	// 检查目录是否存在
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("目录不存在: %s", dirPath)
	}

	var results []ScanResult

	// 遍历目录中的所有文件
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 跳过目录
		if info.IsDir() {
			return nil
		}

		// 检查文件格式是否支持
		if !isSupportedFormat(path) {
			if verbose {
				fmt.Printf("跳过不支持的文件格式: %s\n", path)
			}
			return nil
		}

		if verbose {
			fmt.Printf("正在扫描文件: %s\n", path)
		}

		// 扫描二维码
		content, err := scanQRCode(path, verbose)
		if err != nil {
			if verbose {
				fmt.Printf("扫描失败: %s - 错误: %v\n", path, err)
			}
			// 添加失败的结果
			results = append(results, ScanResult{
				FileName: path,
				Content:  "",
				Success:  false,
				Error:    err.Error(),
			})
			return nil
		}

		// 添加成功的结果
		results = append(results, ScanResult{
			FileName: path,
			Content:  content,
			Success:  true,
			Error:    "",
		})

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("遍历目录失败: %v", err)
	}

	return results, nil
}

// printBatchResults 打印批量扫描结果
func printBatchResults(results []ScanResult) {
	// 彩色终端输出
	color.Cyan("批量扫描结果:")
	color.Yellow("总共扫描文件数: %d", len(results))
	color.Magenta("========================================")

	var fileContent strings.Builder
	fileContent.WriteString(fmt.Sprintf("批量扫描结果:\n总共扫描文件数: %d\n========================================\n", len(results)))

	successCount := 0
	for _, result := range results {
		if result.Success {
			color.Yellow("文件: %s", result.FileName)
			color.Green("状态: 成功")
			color.Green("内容: %s", result.Content)
			successCount++
		} else {
			color.Yellow("文件: %s", result.FileName)
			color.Red("状态: 失败")
			color.Red("错误: %s", result.Error)
		}
		color.White("----------------------------------------")

		// 添加到文件内容
		fileContent.WriteString(fmt.Sprintf("文件: %s\n", result.FileName))
		if result.Success {
			fileContent.WriteString("状态: 成功\n")
			fileContent.WriteString(fmt.Sprintf("内容: %s\n", result.Content))
		} else {
			fileContent.WriteString("状态: 失败\n")
			fileContent.WriteString(fmt.Sprintf("错误: %s\n", result.Error))
		}
		fileContent.WriteString("----------------------------------------\n")
	}

	color.Cyan("扫描完成: 成功 %d 个，失败 %d 个", successCount, len(results)-successCount)
	fileContent.WriteString(fmt.Sprintf("扫描完成: 成功 %d 个，失败 %d 个\n", successCount, len(results)-successCount))

	// 保存到文件
	saveToFile(fileContent.String())
}

// saveToFile 保存内容到output.txt文件
func saveToFile(content string) {
	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("无法打开output.txt文件: %v", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		log.Printf("无法写入output.txt文件: %v", err)
	}
}

// ScanResult 定义扫描结果结构
type ScanResult struct {
	FileName string
	Content  string
	Success  bool
	Error    string
}
