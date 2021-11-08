package main

import (
	"fmt"
	"hello-world/handler"
	"hello-world/util"
	"net/http"
)

func main() {
	RunHttpServer()
	//RunShell("cat /etc/profile")
	//RunShell("tree /Users/zjw/Software")
}

func RunHttpServer() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler)
	http.HandleFunc("/", handler.TestHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server. err: %s\n", err)
	}
}



func RunShell(command string) {
	result, err := util.ExecLinuxShell(command)
	if err != nil {
		fmt.Printf("exec linux shell command failed: %s\n", err.Error())
		return
	}
	fmt.Println(result)
}
