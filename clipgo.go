package main

import (
	"encoding/json"
	"fmt"
	"github.com/atotto/clipboard"
	"gopkg.in/alecthomas/kingpin.v2"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	inspect      = kingpin.Command("inspect", "現在の設定を表示します")
	initial      = kingpin.Command("init", "クリップボードにコピーするテンプレートの初期設定をします")
	templateList = kingpin.Command("list", "利用可能なテンプレートの一覧が表示されます")
	copyCommand  = kingpin.Command("copy", "指定したタグのテンプレートをクリップボードにコピーします")
	tempalteTag  = copyCommand.Flag("tag", "クリップボードにコピーするテンプレートを指定します").Short('t').String()
)

type conf struct {
	Path string `json:"templatePath"`
}

func main() {
	kingpin.Version("0.0.1")
	switch kingpin.Parse() {
	case "inspect":
		inspectConfiguration()
	case "init":
		initialize()
	case "list":
		showTemplates()
	case "copy":
		copyTemplate(*tempalteTag)
	}
}

func inspectConfiguration() {
	userInfo, _ := user.Current()
	configurationFile := userInfo.HomeDir + "/.clipgo.json"
	bytes, err := ioutil.ReadFile(configurationFile)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(configurationFile)
	fmt.Println(string(bytes[:len(bytes)]))
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func initialize() {
	userInfo, _ := user.Current()
	configurationFile := userInfo.HomeDir + "/.clipgo.json"
	if exists(configurationFile) {
		fmt.Println("既に設定ファルが存在します")
	} else {
		file, err := os.Create(configurationFile)
		if err != nil {
			fmt.Println(err)
		}
		configurationBody := `{ "templatePath": "` + userInfo.HomeDir + `/.clipgoTemplate/" }`
		file.WriteString(configurationBody)
		os.Mkdir(userInfo.HomeDir+"/.clipgoTemplate", 0777)
		fmt.Println(userInfo.HomeDir + "/.clipgo.json を作成しました")
		fmt.Println(userInfo.HomeDir + "/.clipgoTemplate にクリップボードにコピーしたいテンプレートを保存して下さい")
	}
}

func getTemplateDir() string {
	userInfo, _ := user.Current()
	configurationFile := userInfo.HomeDir + "/.clipgo.json"

	bytes, err := ioutil.ReadFile(configurationFile)
	if err != nil {
		fmt.Println(err)
	}

	var c conf
	json.Unmarshal(bytes, &c)

	return c.Path
}

func copyTemplate(tag string) error {
	templateDir := getTemplateDir()
	bytes, err := ioutil.ReadFile(templateDir + tag + ".md")
	if err != nil {
		notExist, _ := regexp.MatchString("no such file or directory", err.Error())
		if notExist {
			fmt.Println("指定したファイルは存在しません")
		} else {
			fmt.Println(err.Error())
		}
		return err
	}
	s := string(bytes[:len(bytes)])
	clipboard.WriteAll(s)
	fmt.Println("クリップボードに[" + tag + ".md] の内容をコピーしました")
	return nil
}

func showTemplates() {
	templateDir := getTemplateDir()
	files, err := ioutil.ReadDir(templateDir)
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
