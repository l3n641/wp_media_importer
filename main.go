package main

import (
	"flag"
	"io/ioutil"
	"path"
	"strings"
	"sync"
)

func main() {

	// 图片目录路径
	var imageDirPath string
	// 网站源码根目录
	var webRootPath string
	var coroutineQuantity int
	var images []string
	var wg sync.WaitGroup

	// StringVar用指定的名称、控制台参数项目、默认值、使用信息注册一个string类型flag，并将flag的值保存到p指向的变量
	flag.StringVar(&imageDirPath, "i", "", "图片目录路径")
	flag.StringVar(&webRootPath, "w", "", "网站源码根目录")
	flag.IntVar(&coroutineQuantity, "q", 8, "网站源码根目录")
	flag.Parse()

	tasks := make(chan string, coroutineQuantity)
	imageDirPath = strings.TrimSpace(imageDirPath)
	if imageDirPath == "" || !IsDir(imageDirPath) {
		panic("图片路径不能为空或者不存在")
	}

	webRootPath = strings.TrimSpace(webRootPath)
	if webRootPath == "" || !IsDir(webRootPath) {
		panic("网站路径不能为空者不存在")
	}

	if !IsCommandExist("ls") {
		panic("wp 命令不存在")
	}

	files, err := ioutil.ReadDir(imageDirPath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		wg.Add(1)
		println(file.Name())
		filePath := path.Join(imageDirPath, file.Name())
		images = append(images, filePath)
		tasks <- filePath
		go importMedia(tasks, webRootPath, &wg)
	}
	close(tasks)
	wg.Wait()

}
