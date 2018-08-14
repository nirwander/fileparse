package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"time"
)

// Структура для хранения MAP statement
type repTable struct {
	srcOwner  []byte
	srcName   []byte
	tOwner    []byte
	tName     []byte
	extParams []byte
}

// Структура для хранения конфигурации, получаемой из json файла
type config struct {
	Home       string `json:"ggHome"`
	TableOwner string `json:"ggTableOwner"`
	TableName  string `json:"ggTableName"`
}

// Config - configuration parameters
var Config config

// Aliases - пары идентификаторов подключения к базе из credentialstore
var Aliases map[string]string

func main() {
	//fmt.Println("Hello World!")
	start := time.Now()

	//processReplicatReport(`C:\Users\wander\go\xfecr.txt`)
	getConfig()

	getCredStoreInfo()

	fmt.Printf("\n%s time spent", time.Since(start))
}

// Получаем данные credentialstore для вставки в нужную базу
func getCredStoreInfo() {
	fileBytes, _ := ioutil.ReadFile(`C:\Users\wander\go\cred.txt`)
	lines := bytes.Split(fileBytes, []byte("\n"))

	//Собираем пары alias-userid
	Aliases = make(map[string]string)
	var currAlias string
	for _, line := range lines {
		if bytes.Contains(line, []byte("Alias:")) {
			currAlias = string(bytes.TrimLeft(bytes.TrimSpace(line), string("Alias: ")))
			continue
		}
		if currAlias != "" {
			Aliases[currAlias] = string(bytes.TrimLeft(bytes.TrimSpace(line), string("Userid: ")))
			currAlias = ""
		}
	}
	//fmt.Println(Aliases)
}

func processReplicatReport(fName string) map[string]repTable {
	fileBytes, _ := ioutil.ReadFile(fName)

	lines := bytes.Split(fileBytes, []byte("\n"))
	re := regexp.MustCompile(`(?i)map[[:space:]]+([[:alnum:]_$]+)\.([[:alnum:]_$\?\*\-]+)[[:space:]]*,{0,1}[[:space:]]*target[[:space:]]+([[:alnum:]_$]+)\.([[:alnum:]_$\?\*\-]+)[[:space:]]*,{0,1}[[:space:]]*(.*);`)

	repTables := make(map[string]repTable)
	var c2 int
	var c3 int
	for _, line := range lines {
		// Ищем предложения MAP OWNER.NAME TARGET OWNER.NAME [params] ;
		//fmt.Printf("%d: %s", i, line)
		matches := re.FindSubmatch(line)
		c2++
		if len(matches) > 0 {
			//fmt.Printf("%q\n", matches)
			fmt.Printf("\t%s.%s >> %s.%s, tail: %s\n", matches[1], matches[2], matches[3], matches[4], matches[5])
			repTables[string(matches[3])+"."+string(matches[4])] = repTable{matches[1], matches[2], matches[3], matches[4], matches[5]}
			//str := string(matches[3]) + "." + string(matches[4])
			//fmt.Printf("\t%s\n", str)
			c3++
		}

		if bytes.Contains(line, []byte("Run Time Messages")) {
			break
		}
	}

	fmt.Printf("%s exists\n", repTables["FE_STG.LIMIT_MEASURES"].srcOwner)
	fmt.Printf("%s not exists\n", repTables["FE_STG.LIMIT_MEASURES2"].srcOwner)
	if repTables["FE_STG.LIMIT_MEASURES2"].srcOwner == nil {
		fmt.Println("not exists")
	}
	fmt.Printf("\n%d lines in file\n%d lines matched\n", c2, c3)
	fmt.Printf("%d tables in map", len(repTables))

	return repTables
}

func getConfig() {
	fileBytes, err := ioutil.ReadFile(`C:\Users\wander\go\src\github.com\nirwander\fileparse\config.json`)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", fileBytes)

	err = json.Unmarshal(fileBytes, &Config)
	if err != nil {
		panic(err)
	}

	fmt.Println(Config)

}
