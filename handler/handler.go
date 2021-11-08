package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 返回用于上传文件的页面
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			_, _ = io.WriteString(w, "internal server error")
			return
		}
		_, _ = io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		// 用于处理上传请求
		file, hand, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("Failed to get data: %s\n", err.Error())
			return
		}
		defer file.Close()

		newFile, err := os.Create("/tmp/" + hand.Filename)
		if err != nil {
			fmt.Printf("Failed to create file: %s\n", err.Error())
			return
		}

		defer newFile.Close()

		_, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("Failed to save data into local file: %s\n", err.Error())
			return
		}

		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = io.WriteString(w, "Upload file success.")
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	w.Header().Set("Content-Type", "application/json")
	var data = make(map[string]int)
	data["a"] = 1
	data["b"] = 2
	data["c"] = 3
	data["d"] = 4
	data["e"] = 5
	byteData, _ := json.Marshal(data)
	w.Write(byteData)
	//time.Sleep(time.Second)
	//io.WriteString(w, "internal server error\n")
	//time.Sleep(time.Second)
	//io.WriteString(w, "internal server error")
}
