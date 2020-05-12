package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/bitly/go-simplejson"
	"github.com/mitchellh/mapstructure"
	"github.com/rs/xid"
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

	sends := `{
			"WorkTag": "yaaaaashdhsadjasfas",
			"Title": "The girl of the night",
			"Auth": "mmmaker",
			"Corporator": [],
			"Date": "20190922",
			"Url": [
				"https://pbs.twimg.com/media/EFEV9KxVAAANguo?format=jpg&name=small"
			],
			"Description": "皆様の目の片隅にでも止まっていただければ幸いです",
			"Tags": [
				"オリジナル"
			],
			"Likes": {
				"Amount": 3,
				"Users": [
					"userIdx",
					"userIdy"
				]
			}
		}`

	var body = []byte(sends)

	var CreatedWorkTag string = (xid.New()).String()
	fmt.Println(CreatedWorkTag)
	DesignationUserID := "Id"

	var sendData map[string]interface{}
	if err := json.Unmarshal(body, &sendData); err != nil {
		fmt.Println(59, err)
	}

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	data, err := simplejson.NewJson(bytes)

	_, tof := data.Get(DesignationUserID).CheckGet(CreatedWorkTag)

	if tof {
		var i = 0
		for {
			if i < 5 && tof {
				_, tof = data.CheckGet(CreatedWorkTag)
				if i == 4 {
					fmt.Println("over")
				}
			}
			i++
		}
	} else {
		fmt.Println("ok")
	}
	data.Get(DesignationUserID).SetPath([]string{CreatedWorkTag}, sendData)
	data.Get(DesignationUserID).Get(CreatedWorkTag).Set("WorkTag", CreatedWorkTag)
	fmt.Printf("%+v\n", data.Get(DesignationUserID).Get(CreatedWorkTag))

	works := make([]ID, 0)
	for _, v := range data.MustMap() {
		fake, _ := json.Marshal(v)
		var box ID
		err = json.Unmarshal(fake, &box)
		if err != nil {
			fmt.Println(err)
			break
		}
		works = append(works, box)
	}

	o, _ := data.EncodePretty()
	err = writeFile(filePath, o)
	if err != nil {
		fmt.Println(err)
	}

	fake, err := data.Get(DesignationUserID).Get(CreatedWorkTag).MarshalJSON()
	if err != nil {
		log.Fatal(141)
		return
	}
	var resultJSON ID
	err = json.Unmarshal(fake, &resultJSON)
	fmt.Println("xxx", resultJSON)
	if err != nil {
		log.Fatal(147)
		return
	}

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
