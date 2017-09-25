package data

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"crypto/rand"
	"fmt"
	"crypto/sha1"
	"chih.me/go_web/ChitChat/conf"
)

var Db *sql.DB

func init() {
	var err error
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		conf.Config.DBusername, conf.Config.DBpassword, conf.Config.DBname)
	fmt.Println(dbinfo)
	Db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("cannot generate UUID", err)
	}
	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

func Encrypt(plaintest string) (cryptest string) {
	cryptest = fmt.Sprintf("%x", sha1.Sum([]byte(plaintest)))
	return
}
