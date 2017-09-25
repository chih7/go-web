package data

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type CsvPost struct {
	Id      int
	Content string
	Author  string
}

func csvtest() {
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []CsvPost{
		CsvPost{Id: 1, Content: "Hello World!", Author: "Sau Sheong"},
		CsvPost{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"},
		CsvPost{Id: 3, Content: "Hola Mundo!", Author: "Pedro"},
		CsvPost{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"},
	}

	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{
			strconv.Itoa(post.Id),
			post.Content,
			post.Author,
		}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []CsvPost
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := CsvPost{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}
	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}
