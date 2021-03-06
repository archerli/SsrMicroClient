package SsrDownload

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func get_ssr(path string) {
	file := path + "/shadowsocksr.zip" //源文件路径
	url := "https://github.com/asutorufg/shadowsocksr/archive/asutorufg.zip"
	fmt.Println("Downloading shadowsocksr.zip")
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)
}

func unzip_ssr(path string) {
	// 打开一个zip格式文件
	r, err := zip.OpenReader(path + "/shadowsocksr.zip")
	if err != nil {
		fmt.Println(err)
		return
	}
	var unzip_name string
	for num, k := range r.Reader.File {
		if k.FileInfo().IsDir() {
			if num == 0 {
				unzip_name = "/" + k.Name
			}
			err := os.MkdirAll(path+"/"+k.Name, 0755)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
		r, err := k.Open()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("unzip: ", k.Name)
		defer r.Close()
		NewFile, err := os.Create(path + "/" + k.Name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		io.Copy(NewFile, r)
		NewFile.Close()
	}

	err = os.Rename(path+unzip_name, path+"/shadowsocksr")
	if err != nil {
		log.Println(err)
		return
	}
}

func Get_ssr_python(path string) {
	get_ssr(path)
	unzip_ssr(path)

	err := os.Remove(path + "/shadowsocksr.zip")
	if err != nil {
		//如果删除失败则输出 file remove Error!
		log.Println("file remove Error!")
		//输出错误详细信息
		log.Println(err)
	}
}
