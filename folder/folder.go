package folder

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// D:\Temp
func Register() {
	// 开启用户输入
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("请输入要查找的路径：")
	scanner.Scan()

	text := scanner.Text()
	fmt.Println(text)
	_, err := os.Stat(text)
	if os.IsNotExist(err) {
		os.Exit(1)
	} else {
		defaulFolder(text)
		removeFileToFolder(text)
	}
	// D:\goStudy
	// os.Stat()
}

// 建立文件夹
func defaulFolder(basePath string) {
	extTypes := []string{"Image", "Video", "Txt"}
	for _, name := range extTypes {
		_, err := os.Stat(name)
		if os.IsNotExist(err) {
			os.Mkdir(filepath.Join(basePath, name), 0755)
		}
		// 字符串拼接
	}
}

func removeFileToFolder(basePath string) {
	files, _ := os.ReadDir(basePath)
	for _, file := range files {
		fileInfo, err := file.Info()
		if err != nil {
			continue
		}
		oldPath := filepath.Join(basePath, fileInfo.Name())
		ext := filepath.Ext(oldPath)
		fmt.Println(ext)
		switch ext {
		case ".gif", ".jpg":
			newPath := filepath.Join(basePath, "Image", fileInfo.Name())
			err := os.Rename(oldPath, newPath)
			if err != nil {
				fmt.Println(err)
			}
		case ".mp4", ".mkv":
			newPath := filepath.Join(basePath, "Video", fileInfo.Name())
			err := os.Rename(oldPath, newPath)
			if err != nil {
				fmt.Println(err)
			}
		case ".pdf", ".txt":
			newPath := filepath.Join(basePath, "Txt", fileInfo.Name())
			err := os.Rename(oldPath, newPath)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
