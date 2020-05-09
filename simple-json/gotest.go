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

	raws := `{
			"WorkTag": "x",
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

	var tes map[string]interface{}
	if err := json.Unmarshal([]byte(raws), &tes); err != nil {
		fmt.Println(59, err)
	}

	// fmt.Printf("%+v\n", tes)
	fmt.Println()

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	js, err := simplejson.NewJson(bytes)

	js.Get("Id").SetPath([]string{"WorkId3"}, tes)

	ids := js.Get("Id")
	fmt.Println(ids)

	var i = 1
	// ss, tf := js.Get("Id").CheckGet("WorkId3")
	works := make([]ID, 0)
	for _, v := range ids.MustMap() {
		fake, _ := json.Marshal(v)
		var box ID
		err = json.Unmarshal(fake, &box)
		if err != nil {
			fmt.Println(err)
			break
		}
		works = append(works, box)
		// fmt.Println()
		i++
	}
	fmt.Println()

	fmt.Printf("xxx %+v", works)

	// fmt.Println(tf)
	// if !tf {
	// 	fmt.Println("none")
	// } else {
	// 	fmt.Printf("this %+v\n", ss)
	// }

	// fmt.Printf("%+v\n", js)

	w, err := os.Create("./result.json")
	// defer w.Close()
	// o, _ := js.MarshalJSON()
	// o, _ := js.Encode()
	// o, _ := js.EncodePretty()
	//oはバイナリ
	// fmt.Printf("%+v\n", o)
	o, err := json.Marshal(works)
	w.Write(o)
	// err = writeFile("works2.json", o)
	// if err != nil {
	// 	fmt.Println(err)
	// }
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
