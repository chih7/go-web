package data

import "fmt"

type MemPost struct {
	Id      int
	Content string
	Author  string
}

var PostById map[int]*MemPost
var PostsByAuthtor map[string][]*MemPost

func store(post MemPost) {
	PostById[post.Id] = &post
	PostsByAuthtor[post.Author] = append(PostsByAuthtor[post.Author], &post)
}

func memtest() {
	PostById = make(map[int]*MemPost)
	PostsByAuthtor = make(map[string][]*MemPost)

	post1 := MemPost{Id: 1, Content: "Hello World", Author: "chih"}
	post2 := MemPost{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"}
	post3 := MemPost{Id: 3, Content: "Hola Mundo!", Author: "Pedro"}
	post4 := MemPost{Id: 4, Content: "Greetings Earthlings!", Author: "chih"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthtor["chih"] {
		fmt.Println(post)
	}

	for _, post := range PostsByAuthtor["Pedro"] {
		fmt.Println(post)
	}

}
