package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
	"sync"
)

func main1() {
	var wg sync.WaitGroup
	srcDir := flag.String("srcDir", "/Users/zjw/Downloads/Golang语言企业级项目-仿百度网盘（完整资料）", "文件目录")

	//subStr := flag.String("subStr", `（最新）【公司实战】vue后台管理系统-新人入职必看.*[0-9]\.`, "需要去掉的字串")
	//subStr := flag.String("subStr", `【樱花论坛 www.sakuraaaa.com】`, "需要去掉的字串")
	subStr := flag.String("subStr", `【樱花论坛 www.sakuraaaa.com】`, "需要去掉的字串")

	//srcDir := flag.String("srcDir", "/tmp", "文件目录")
	//subStr := flag.String("subStr", "[tmp]", "需要去掉的字串")
	flag.Parse()
	wg.Add(1)
	go RenameFiles(*srcDir, *subStr, &wg)
	wg.Wait()
	fmt.Printf("目录「%v」文件名称修改完成\n", *srcDir)
}

func RenameFiles(dir string, subStr string, wg *sync.WaitGroup) {
	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		if file.IsDir() {
			wg.Add(1)
			go RenameFiles(path.Join(dir, file.Name()), subStr, wg)
		}
		oldFileName := path.Join(dir, file.Name())
		tmpFileName := RemoveSubstr(file.Name(), subStr)
		newFileName := path.Join(dir, tmpFileName)
		os.Rename(oldFileName, newFileName)
	}
	wg.Done()
}

func RemoveSubstr(fileName string, subStr string) string {
	reg, err := regexp.Compile(subStr)
	if err !=nil {
		panic(err)
	}
	if reg.MatchString(fileName) {
		splits := reg.Split(fileName, -1)
		// splits := strings.Split(fileName, subStr)
		return strings.Join(splits, "")
	} else {
		return fileName
	}
}
