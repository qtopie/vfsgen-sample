package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/artificerpi/vfsgen-sample/data"
)

var version = "unset"

func main() {

	log.Println("Program version:", version)

	file, err := data.Assets.Open("hello.txt")
	if err != nil {
		log.Fatal(err)
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("The content of file hello.txt is:", string(content))
	defer file.Close()

	http.Handle("/", http.FileServer(data.Assets))
	log.Fatal(http.ListenAndServe(":7070", nil))
}
