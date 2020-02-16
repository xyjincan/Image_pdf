package main

import (
	"./imgs2pdf"
	"bufio"
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

func main() {

	///imgs2pdf.CheckData()
	var debug = true
	//imgs2pdf.WipeMargin("./images/abc.jpg",10,10,10,10)

	if debug{
		//return
	}
	var path string
	stdinReader := bufio.NewReader(nil)
	for true {
		fmt.Print("请输入图片所在文件夹：")
		fmt.Scanln(&path)
		stdinReader.Reset(os.Stdin)
		if strings.HasPrefix(path, ".") {
			fmt.Println("相对路径自动补全:" + path)
			path = GetBaePath() + path
		}
		path, _ = filepath.Abs(path)
		// 开始合成
		fmt.Println("图片路径:" + path)
		//检查
		isnature,files := imgs2pdf.IsNaturalNum(path)
		fmt.Println("输入调整选项（回车默认0不调整）")
		fmt.Print("调整图片百分比值：")
		var ih float64
		fmt.Scanln(&ih)
		stdinReader.Reset(os.Stdin)
		fmt.Println("get:",ih)
		fmt.Print("进行上下和左右调整（0 0）：")
		var idown,ileft float64
		fmt.Scanln(&idown,&ileft)
		stdinReader.Reset(os.Stdin)
		fmt.Println("get:",idown,ileft)
		stdinReader.Reset(os.Stdin)

		var pdfbook *imgs2pdf.PdfBook
		pdfbook = imgs2pdf.NewPdfBook()
		pdfbook.SetZoomin(ih)
		pdfbook.SetUp_LeftMove(idown,ileft)
		pdfbook.SetImgsPath(path)
		if isnature{
			pdfbook.AddPagesWithFiles(files)
		}else {
			pdfbook.AddPages()
		}
		if pdfbook.Save(){
			fmt.Println("文件：" + pdfbook.GetPdfName())
		}
		// 输入
		fmt.Println("已完成")
	}
}
