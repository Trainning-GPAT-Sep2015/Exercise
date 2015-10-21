package main

import (
	//"encoding/json"
	// "errors"
	"fmt"
	"io/ioutil"
)

type Blog struct {
	Title    string `json: "title"`
	Summary  string `json: "summary"`
	Body     string `json: "body"`
	FileName string `json: "filename"`
}

func showAllArticle(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var list_file_name []string
	for _, f := range files {
		list_file_name = append(list_file_name, f.Name())
	}
	return list_file_name, nil
}

func readArticle(filename string) {

}

func main() {
	l, _ := showAllArticle("../data")
	for _, v := range l {
		fmt.Println(v)
	}
}
