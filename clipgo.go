package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

var (
	templateList = kingpin.Command("list", "help string")
	// tempalteTag    = kingpin.Flag("tag", "Select server id to speedtest").Short('t').String()
)

// ./.clipgo/clipgo.json

func main() {
	kingpin.Version("0.0.1")
	switch kingpin.Parse() {
	case "list":
		showTemplates()

	}
}

func showTemplates() {
	files, err := ioutil.ReadDir("./.clipgo/templates")
	if err != nil {
		log.Println(err)
	}
	fmt.Println("登録されているテンプレートは...")
	for _, file := range files {
		fileName := trimExtension(file.Name())
		fmt.Println(fileName)
	}
}

func trimExtension(filename string) string {
	extension := filepath.Ext(filename)
	withoutExtension := strings.TrimRight(filename, extension)
	return withoutExtension
}
