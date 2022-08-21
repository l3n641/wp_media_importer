package main

import (
	"log"
	"os"
	"os/exec"
	"sync"
)

func IsCommandExist(command string) bool {
	_, err := exec.LookPath(command)
	if err != nil {
		return false
	} else {
		return true
	}
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// wp media import ./benetton-2022-08-03/*.*  --allow-root --path=/www/wwwroot/www.vova.show/

func importMedia(ch chan string, lock chan int, webRootPath string, wg *sync.WaitGroup) {
	for val := range ch {
		//wp media import ./benetton-2022-08-03/*.*  --allow-root --path=/www/wwwroot/www.vova.show/
		path := "--path=" + webRootPath
		cmd := exec.Command("wp", "media", "import", val, "--allow-root", path)
		data, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("failed to call Output(): %s", data)
		}
		log.Printf("output: %s", data)
		wg.Done()
		<-lock

	}
}
