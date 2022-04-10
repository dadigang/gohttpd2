package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var port = flag.String("port", ":8080", "绑定本地端口")
var root = flag.String("root", "./", "默认文件路径")

func main() {

	flag.Parse()
	fmt.Println(getCurrentDirectory())
	fmt.Println(*port)
	fmt.Println(*root)

	os.Mkdir(*root, 0777)
	http.Handle("/", http.FileServer(http.Dir(*root)))

	e := http.ListenAndServe(*port, nil)

	fmt.Println(e)
}
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func printip(){
	
}