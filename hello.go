package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
)

var URL string = "http://127.0.0.1:9200"

type Employment struct {
	name1 string
	name2 string
	age int
	saying string
	loves []string
}

func createindex() {
	Client, err := elastic.NewClient(elastic.SetSniff(false),elastic.SetURL(URL))
	fmt.Println(Client, err)
	name := "people2"
	Client.CreateIndex(name).Do(context.Background())
}

func insertdata() {
	Client, err := elastic.NewClient(elastic.SetSniff(false),elastic.SetURL(URL))
	fmt.Println(Client, err)
	name := "people2"
	data := `{
	"name": "wali",
		"country": "Chian",
		"age": 30,
		"date": "1987-03-07"
}`
	_, err = Client.Index().Index(name).Type("man1").Id("1").BodyJson(data).Do(context.Background())
	if err != nil {
		panic(err)
	}
}

func searchbyid() {
	Client, err := elastic.NewClient(elastic.SetSniff(false),elastic.SetURL(URL))
	fmt.Println(Client, err)
	name := "people2"
	getResult, err := Client.Get().Index(name).Type("man1").Id("1").Do(context.Background())
	fmt.Println(getResult, err)
}

func update() {
	Client, err := elastic.NewClient(elastic.SetSniff(false),elastic.SetURL(URL))
	response, err := Client.Update().Index("people2").Type("man1").Id("1").
		Doc(map[string]interface{}{"age": 88}).Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	fmt.Println("update age %s\n", response.Result)
}

func deletedate()  {
	Client, err := elastic.NewClient(elastic.SetSniff(false),elastic.SetURL(URL))
	e1 := Employment{"Jane", "Smith", 32, "I like to collect rock albums", []string{"music"}}
	Client.Index().Index("megacorp").Type("employee").Id("1").BodyJson(e1).Do(context.Background())
	if err!=nil {
		panic(err)
	}
	getResult, err := Client.Get().Index("megacorp").Type("employee").Id("1").Do(context.Background())
	fmt.Println(getResult,err)
}

func main()  {
	update()
}