package imgs2pdf

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var imageSet map[string]bool

func init() {

	imageSet = map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
		".gif":  true,
		".bmp":  true,
		".jfif": true,
	}

}

func IsImageFile(file string) bool {

	fullFilename := strings.ToLower(file)
	var filenameWithSuffix string
	filenameWithSuffix = path.Base(fullFilename) //获取文件名带后缀
	var fileSuffix string
	fileSuffix = path.Ext(filenameWithSuffix) //获取文件后缀
	return imageSet[fileSuffix]
}

func GetMainFileDir() string {
	//fmt.Println("os. Getwd:")
	fmt.Println(os.Getwd())

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	if strings.HasPrefix(dir, os.TempDir()) {
		fmt.Println("in go run:(need add more)")
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
	dir = dir + string(filepath.Separator)

	return dir
}

func GetDefaultStartDir() string {

	realpath, _ := os.Getwd()
	fmt.Println("os. Getwd: (real path): " + realpath)

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
	dir = dir + string(filepath.Separator)

	return dir
}

func GetDIr() string {

	//获取当前路径
	// .
	// ..
	// abc
	
	return "."
}

func GetFullPath(endpath string) string {

	return GetDIr() + string(filepath.Separator) + endpath
}

func IsDir(f string) bool {
	fi, e := os.Stat(f)
	if e != nil {
		return false
	}
	return fi.IsDir()
}

func IsFile(f string) bool {
	fi, e := os.Stat(f)
	if e != nil {
		return false
	}
	return !fi.IsDir()
}

// IsExist checks whether a file or directory exists.
// It returns false when the file or directory does not exist.
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
