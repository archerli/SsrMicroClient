package ssr_init

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	//"runtime"
	"sync"
	//"path/filepath"
	//"database/sql"
	"io/ioutil"

	"path/filepath"

	"../config"
	SsrDownload "../shadowsocksr"
	"../subscription"
)

//判断目录是否存在返回布尔类型
func Path_exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		} else {
			return false
		}
	} else {
		return true
	}
}

func Init(config_path, sql_db_path string) {
	//判断目录是否存在 不存在则创建
	if !Path_exists(config_path) {
		err := os.MkdirAll(config_path, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}

	if !Path_exists(sql_db_path) {
		var wg sync.WaitGroup

		wg.Add(1)
		go subscription.Subscription_link_init(sql_db_path, &wg)
		wg.Add(1)
		go subscription.Init_config_db(sql_db_path, &wg)
		wg.Add(1)
		go subscription.Ssr_server_node_init(sql_db_path, &wg)
		Auto_create_config(config_path)

		wg.Wait()
	}

	if !Path_exists(config_path + "/shadowsocksr") {
		SsrDownload.Get_ssr_python(config_path)
	}
}

func Menu_init(path string) {
	//获取当前可执行文件目录
	file, _ := exec.LookPath(os.Args[0])
	path2, _ := filepath.Abs(file)
	//fmt.Println(path2)
	rst := filepath.Dir(path2)
	//fmt.Println(rst)

	fmt.Println("当前配置文件目录:" + path)
	fmt.Println("当前可执行文件目录:" + rst)
}

func Auto_create_config(path string) {
	in_line := "\n"
	deamon := "deamon" + in_line
	config_path := path + "/ssr_config.conf"
	ssr_path := "#" + path + "/shadowsocksr/shadowsocks/local.py #ssr路径" + in_line
	pid_file := "pid-file " + path + "/shadowsocksr.pid" + in_line
	log_file := "log-file /dev/null" + in_line
	fast_open := "fast-open" + in_line
	workers := "workers 8" + in_line
	local_address := "#local_address 127.0.0.1" + in_line
	local_port := "#local_port 1080" + in_line
	connect_verbose_info := "#connect-verbose-info" + in_line
	// acl := "#acl " + path + "/aacl-none.acl" + in_line
	acl := ""
	python_path := "#python_path " + config.Get_python_path() + "#python路径" + in_line

	if runtime.GOOS == "windows" {
		in_line = "\r\n"
		deamon = "#deamon" + in_line
		config_path = path + `\ssr_config.conf`
		ssr_path = "#" + path + "\\shadowsocksr\\shadowsocks\\local.py #ssr路径" + in_line
		pid_file = ""
		log_file = ""
	}

	config_conf := python_path + ssr_path + pid_file + log_file + fast_open + deamon + workers + local_address + local_port + connect_verbose_info + acl
	fmt.Println(config_conf)
	ioutil.WriteFile(config_path, []byte(config_conf), 0644)
}
