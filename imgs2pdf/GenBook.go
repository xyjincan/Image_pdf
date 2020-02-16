package imgs2pdf

import (
	"fmt"
	"github.com/signintech/gopdf"
	"io/ioutil"
	"math"
	"path/filepath"
	"strings"
)

var use_path_as_name = true

const hdw = 842.0/595.0


type PdfBook struct {
	//保存文件名
	name string
	//图片数据路径
	imgs_path string
	//pdf实例
	pdf *gopdf.GoPdf

	dynamicPageSize *gopdf.Rect

	// 放大系数 （有效去除边框）
	ih float64 // 20
	iw float64

	//  上下左右微调
	move_down float64//23.0

	// 左右不对称数据处理 \
	// TODO
	move_left float64
	oddd bool
	startleft int
	countadd int
}


func (pb *PdfBook) initdata()  {
	pb.dynamicPageSize =  &gopdf.Rect{W: 595, H: 842};
	pb.countadd = 1

}
func NewPdfBook() *PdfBook {
	pdfb := &PdfBook{}
	pdfb.initdata()
	return pdfb
}

func NewPdfBookWithName(name, img_path string) *PdfBook {
	pdfb := &PdfBook{name: name, imgs_path: img_path}
	pdfb.initdata()
	return pdfb
}

func (pb *PdfBook) SetImgsPath(path string) bool {
	if IsDir(path) {
		pb.imgs_path = path
		return true
	}
	//fmt.Println("set imgspath fail: "+path)
	return false
}

func (pb *PdfBook) SetZoomin(num float64)  {
	pb.ih = num
	pb.iw = num
}

func (pb *PdfBook) SetUp_LeftMove(down ,left float64)  {
	pb.move_down = down
	pb.move_left = left
}

func (pb *PdfBook) checkdata()  {
	pb.countadd = 1
	if pb.ih < -100 || pb.ih > 100{
		pb.ih = 0
	}
	if pb.iw < -100 || pb.iw > 100{
		pb.iw = 0
	}
	if math.Abs(pb.move_down) >= 842{
		pb.move_down = 0
	}
	if math.Abs(pb.move_left) >= 595{
		pb.move_left = 0
	}
}

func (pb *PdfBook) AddPage(item string){
	var err error
	fmt.Println("add: " + item)
	//fmt.Println("add new page")
	pb.pdf.AddPage()
	var add_height = 0.0;
	var add_width = 0.0;
	add_height = (pb.ih/100.0)*842.0/2.0
	add_width = (pb.iw/100.0)*595.0/2.0
	pb.dynamicPageSize =  &gopdf.Rect{W: 595+add_width*2, H: 842+add_height*2};
	// 放大后尝试居中 & 上下左右调整
	err = pb.pdf.Image(item, (-add_width) + pb.move_left, (-add_height) + pb.move_down, pb.dynamicPageSize) //print image
	if err != nil {
		fmt.Println(err)
	}
}


//
func (pb *PdfBook) AddPagesWithFiles(files[] string) bool{

	pb.checkdata()
	if pb.pdf == nil {
		pb.pdf = &gopdf.GoPdf{}
		pb.pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	}

	if pb.imgs_path == "" {
		return false
	}

	for _, file := range files {
		str := file
		var item = pb.imgs_path + string(filepath.Separator) + str
		if IsDir(item) || !IsImageFile(item) {
			continue
		}
		pb.AddPage(item)
		pb.countadd++
	}
	return true

}



func (pb *PdfBook) AddPages() bool {

	pb.checkdata()
	if pb.pdf == nil {
		pb.pdf = &gopdf.GoPdf{}
		pb.pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	}

	if pb.imgs_path == "" {
		return false
	}

	var err error
	files, err := ioutil.ReadDir(pb.imgs_path)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		str := file.Name()
		var item = pb.imgs_path + string(filepath.Separator) + str
		if IsDir(item) || !IsImageFile(item) {
			continue
		}
		pb.AddPage(item)
		pb.countadd++
	}
	return true
}


func (pb *PdfBook) Save() bool {

	if pb.pdf == nil {
		return false
	}

	if use_path_as_name && strings.Index(pb.imgs_path, string(filepath.Separator)) != -1 {
		pb.name = filepath.Base(pb.imgs_path) + ".pdf"
	}
	if pb.name == "" {
		pb.name = "gen.pdf"
	}

	if IsExist(pb.name) {
		fmt.Println("(error) file exist: " + pb.name)
		return false
	}
	var err error
	err = pb.pdf.WritePdf(pb.name)
	if err != nil {
		fmt.Println(err)
		return false
	}
	pb.pdf.Close()
	pb.pdf = nil
	return true
}

func (pb *PdfBook) GetPdfName() string {
	return pb.name
}

func (pb *PdfBook) String() string {
	return "{ pdf: " + pb.name + ", imgs_path: " + pb.imgs_path + " }"
}

