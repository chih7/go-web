package data

import (
	"chih.me/go_web/ChitChat/conf"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

type OrmPost struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	Comments  []Comment
	CreatedAt time.Time
}

type Comment struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	OrmPostId int    `sql:"index"`
	CreatedAt time.Time
}

var OrmDb *gorm.DB

func init() {
	var err error
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		conf.Config.DBusername, conf.Config.DBpassword, conf.Config.DBname)
	OrmDb, err = gorm.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	OrmDb.AutoMigrate(&OrmPost{}, &Comment{})
}

func ormtest() {
	post := OrmPost{Content: "Hello World!", Author: "Sau Sheong"}
	fmt.Println(post)

	OrmDb.Create(&post)
	fmt.Println(post)

	comment := Comment{Content: "Good post!", Author: "Joe"}
	OrmDb.Model(&post).Association("Comments").Append(comment)

	var readPost OrmPost
	OrmDb.Where("author = $1", "Sau Sheong").First(&readPost)
	var comments []Comment
	OrmDb.Model(&readPost).Related(&comments)
	fmt.Println(comments[0])
}
