package main

import (
	"./imgs2pdf"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// 用于读取程序内部相关配置文件
func GetMainFilePath() string {

	realpath, _ := os.Getwd()
	//fmt.Println("os. Getwd: (real path): " + realpath)
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	if strings.HasPrefix(dir, os.TempDir()) {
		//fmt.Println("in go run:")
		// 1 2 3 4 5 6 7
		// 注意包所在路径层次
		//_, filename, _, _ := runtime.Caller(2)
		filename, _ := os.Getwd()
		dir = filepath.Dir(filename)
	} else {
		//_, filename, _, ok := runtime.Caller(1)
		//fmt.Println(filename,ok)
		fmt.Println("todo:")
	}

	if strings.Contains(realpath, dir) {
		dir = realpath
	}
	dir = dir + string(filepath.Separator)

	return dir
}

// 用于获取相对目录文件的全路径
func GetBaePath() string {
	realpath, _ := os.Getwd()
	realpath = realpath + string(filepath.Separator)
	return realpath
}

// 将图片批量合成为一个pdf
func GenPdf(path string) string {

	var pdfbook *imgs2pdf.PdfBook
	pdfbook = imgs2pdf.NewPdfBook()
	pdfbook.SetImgsPath(path)
	pdfbook.AddPages()
	pdfbook.Save()
	//fmt.Println("pdfbook: ",pdfbook)
	return pdfbook.GetPdfName()
}

func main() {

	var path string
	for true {
		fmt.Print("请输入图片所在文件夹：")
		fmt.Scanln(&path)
		if strings.HasPrefix(path, ".") {
			fmt.Println("相对路径自动补全:" + path)
			path = GetBaePath() + path
		}
		path, _ = filepath.Abs(path)
		fmt.Println("图片路径:" + path)
		fmt.Println("文件：" + GenPdf(path))
		fmt.Println("已完成")
	}
}
