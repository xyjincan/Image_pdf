package imgs2pdf

import (
	"fmt"
	"github.com/signintech/gopdf"
	"io/ioutil"
	"path/filepath"
	"strings"
)

var use_path_as_name = true

type PdfBook struct {
	//保存文件名
	name string
	//图片数据路径
	imgs_path string
	//pdf实例
	pdf *gopdf.GoPdf
}

func NewPdfBook() *PdfBook {
	return &PdfBook{}
}

func NewPdfBookWithName(name, img_path string) *PdfBook {
	return &PdfBook{name: name, imgs_path: img_path}
}

func (pbook *PdfBook) SetImgsPath(path string) bool {
	if IsDir(path) {
		pbook.imgs_path = path
		return true
	}
	//fmt.Println("set imgspath fail: "+path)
	return false
}

func (pbook *PdfBook) AddPages() bool {

	if pbook.pdf == nil {
		pbook.pdf = &gopdf.GoPdf{}
		pbook.pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	}

	if pbook.imgs_path == "" {
		return false
	}
	//fmt.Println("AddPages: " + pbook.imgs_path)
	var err error

	files, err := ioutil.ReadDir(pbook.imgs_path)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(files)
	for _, file := range files {
		str := file.Name()
		var item = pbook.imgs_path + string(filepath.Separator) + str
		if IsDir(item) || !IsImageFile(item) {
			continue
		}
		fmt.Println("add: " + item)
		//fmt.Println("add new page")
		pbook.pdf.AddPage()
		err = pbook.pdf.Image(item, 0, 0, gopdf.PageSizeA4) //print image
		if err != nil {
			fmt.Println(err)
		}
	}

	return true
}

func (pbook *PdfBook) Save() bool {

	if pbook.pdf == nil {
		return false
	}

	if use_path_as_name && strings.Index(pbook.imgs_path, string(filepath.Separator)) != -1 {
		pbook.name = filepath.Base(pbook.imgs_path) + ".pdf"
		//fmt.Println("as name:" + pbook.name)
	}
	if pbook.name == "" {
		pbook.name = "gen.pdf"
	}

	if IsExist(pbook.name) {
		fmt.Println("file exist: " + pbook.name)
		return false
	}
	var err error
	err = pbook.pdf.WritePdf(pbook.name)
	if err != nil {
		fmt.Println(err)
	}
	pbook.pdf.Close()
	pbook.pdf = nil
	return true
}

func (pbook *PdfBook) GetPdfName() string {
	return pbook.name
}

func (pbook *PdfBook) String() string {
	return "{ pdf: " + pbook.name + ", imgs_path: " + pbook.imgs_path + " }"
}
