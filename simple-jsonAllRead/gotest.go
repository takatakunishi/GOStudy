package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bitly/go-simplejson"
	"github.com/mitchellh/mapstructure"
)

func main() {
	getAllData()
}

func readFile(filename string) ([]byte, error) {
	byte, err := ioutil.ReadFile(filename)
	return byte, err
}

func getAllData() {
	filePath := "works3.json"

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	rawData, err := simplejson.NewJson(bytes)

	data := rawData.Get("Id")

	works := make([]ID, 0)
	for _, v := range data.MustMap() {
		fake, _ := json.Marshal(v)
		if err != nil {
			fmt.Println(err)
			break
		}
		var box ID
		err = json.Unmarshal(fake, &box)
		if err != nil {
			fmt.Println(err)
			break
		}
		works = append(works, box)
	}

	fmt.Printf("xxx %+v", works)
}

func writeFile(filename string, bytes []byte) (err error) {
	return ioutil.WriteFile(filename, bytes, os.ModePerm)
}

func makeJSON(bytes []byte, body []byte) (ids Data, err error) {
	var data map[string]interface{}
	var Datas Data

	if err := json.Unmarshal(bytes, &data); err != nil {
		return Datas, err
	}
	//mapにエンコード

	var work ID
	if body != nil {
		fmt.Println("in")
		// numはsliceの大きさを指定するため。
		//なお大きさはユーザーの作品数に依存する
		if err := json.Unmarshal(body, &work); err != nil {
			return Datas, err
		}
		// err = mapstructure.Decode(data, &work)
		if err != nil {
			return Datas, err
		}
	}

	works := make([]ID, 1)
	err = mapstructure.Decode(data["Id"], &works)
	if err != nil {
		return Datas, err
	}

	if body != nil {
		works = append(works, work)
	}

	Datas.Id = works
	return Datas, nil
}

// ID is each works data
type ID struct {
	WorkTag     string   `json:"WorkTag"`
	Title       string   `json:"Title"`
	Auth        string   `json:"Auth"`
	Corporator  []string `json:"Corporator"`
	Date        string   `json:"Date"`
	URL         []string `json:"Url"`
	Description string   `json:"Description"`
	Tags        []string `json:"Tags"`
	//Likes is
	Likes struct {
		Amount int `json:"Amount"`
		// Users is about other User of this user
		Users []string `json:"Users"`
	} `json:"Likes"`
}

// Data is each person data
type Data struct {
	Id []ID `json:"Id"`
}
