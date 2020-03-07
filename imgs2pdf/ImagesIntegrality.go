package imgs2pdf

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)


// 全局变量状态(非线程安全)

var g_natural = false
var g_prefix = ""
var g_suffix = ""
var g_minpage = -1
var g_maxpage = 0
var g_pagemap map[int]string



func CheckData() {

	files_path := ""
	for true {
		fmt.Println(files_path)
		files_path = strings.Trim(files_path, " ")
		if len(files_path) == 0 {
			files_path = `xxx`
		}
		fmt.Println(IsNaturalNum(files_path))
		files_path = ""
		fmt.Print("\n输入检查路径：")
		fmt.Scanln(&files_path) //TODO 空格
	}

}

/*
	是否包含数字，以及固定前后缀
*/

func IsNaturalNum(data_path string) (natural bool,imgfiles[] string){

	var count = 0
	var max_page = 0
	var min_page = -1
	var pagemap map[int]string = make(map[int]string)

	var base_prefix string
	var base_suffix string

	// 读取当前目录中的所有文件和子目录
	files, err := ioutil.ReadDir(data_path)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(files) == 0 {
		return
	}
	// 获取文件，并输出它们的名字
	fmt.Println("check files integrity：" + data_path)
	for _, file := range files {
		str := file.Name()
		if !IsImageFile(str) {
			continue
		}
		count++
		if count == 1 {
			base_prefix = file.Name()
			base_suffix = file.Name()
		}
		fmt.Println(file.Name())
		for base_prefix != "" {
			if strings.HasPrefix(str, base_prefix) {
				break
			} else {
				base_prefix = base_prefix[0 : len(base_prefix)-1]
			}
		}
		for base_suffix != "" {
			if strings.HasSuffix(str, base_suffix) {
				break
			} else {
				base_suffix = base_suffix[1:len(base_suffix)]
			}
		}
	}

	fmt.Println("images:", count)
	fmt.Println("base_prefix:", base_prefix)
	fmt.Println("base_suffix:", base_suffix)

	// check match 系统排序or自然数排序
	var matchcount = 0
	for _, file := range files {
		var str = file.Name()
		if !IsImageFile(str) {
			continue
		}
		var tpage = str[len(base_prefix) : len(str)-len(base_suffix)]
		page, err := strconv.Atoi(tpage)
		if err != nil {
			continue
		}
		//fmt.Println(str)
		//fmt.Println(page)
		if max_page < page {
			max_page = page
		}
		if min_page >= page || min_page == -1{
			min_page = page
		}
		pagemap[page] = str
		matchcount++
	}
	// if can find miss
	if matchcount > 0 {
		//
		// do
		fmt.Println("min: ", min_page)
		fmt.Println("max: ", max_page)
		fmt.Println("matchcount: ", matchcount)
		var hasmiss = false
		for i := min_page; i <= max_page; i++ {
			if pagemap[i] == "" {
				if(!hasmiss){
					fmt.Print("also need: ")
					hasmiss = true
				}
				fmt.Print(i, ",")
			}else {
				imgfiles = append(imgfiles, pagemap[i])
			}
		}
		if(hasmiss){
			fmt.Println( "")
		}
		g_natural = true
		return true,imgfiles
	} else {
		fmt.Println("use file system sequence")
		g_natural = false
		return  false,nil
	}

}

/*


// 全局变量状态(非线程安全)

var g_natural = false
var g_prefix = ""
var g_suffix = ""
var g_minpage = -1
var g_maxpage = 0
var g_pagemap map[int]string



*/
