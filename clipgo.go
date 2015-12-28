package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	templateList = kingpin.Command("list", "利用可能なテンプレートの一覧が表示されます")
	copyCommand  = kingpin.Command("copy", "指定したタグのテンプレートをクリップボードにコピーします").Default()
	tempalteTag  = copyCommand.Flag("tag", "クリップボードにコピーするテンプレートを指定します").Short('t').String()
)

func main() {
	kingpin.Version("0.0.1")
	switch kingpin.Parse() {
	case "list":
		showTemplates()
	case "copy":
		copyTemplate(*tempalteTag)
	}
}

func copyTemplate(tag string) {
	bytes, err := ioutil.ReadFile("./.clipgo/templates/" + tag + ".md")
	if err != nil {
		notExist, _ := regexp.MatchString("no such file or directory", err.Error())
		if notExist {
			fmt.Println("指定したファイルは存在しません")
		} else {
			fmt.Println(err.Error())
		}
	}
	s := string(bytes[:len(bytes)])
	clipboard.WriteAll(s)
	fmt.Println("クリップボードに[" + tag + ".md] の内容をコピーしました")
}

func showTemplates() {
	files, err := ioutil.ReadDir("./.clipgo/templates")
	if err != nil {
		fmt.Println(err)
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
