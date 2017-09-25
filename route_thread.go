package main

import (
	"net/http"
	"chih.me/go_web/ChitChat/data"
	"fmt"
	"chih.me/go_web/ChitChat/utils"
	"chih.me/go_web/ChitChat/log"
)

func newThread(w http.ResponseWriter, r *http.Request) {
	_, err := utils.Session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", 302)
	} else {
		utils.GenerateHTML(w, nil, "layout", "private.navbar", "new.thread")
	}
}

func createThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := utils.Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			log.Danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			log.Danger(err, "Cannot get user from session")
		}
		topic := request.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			log.Danger(err, "Cannot create thread")
		}
		http.Redirect(writer, request, "/", 302)
	}
}

// GET /thread/read
// Show the details of the thread, including the posts and the form to write a post
func readThread(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	thread, err := data.ThreadByUUID(uuid)
	if err != nil {
		utils.Error_message(writer, request, "Cannot read thread")
	} else {
		_, err := utils.Session(writer, request)
		if err != nil {
			utils.GenerateHTML(writer, &thread, "layout", "public.navbar", "public.thread")
		} else {
			utils.GenerateHTML(writer, &thread, "layout", "private.navbar", "private.thread")
		}
	}
}

// POST /thread/post
// Create the post
func postThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := utils.Session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			log.Danger(err, "Cannot parse form")
		}
		user, err := sess.User()
		if err != nil {
			log.Danger(err, "Cannot get user from session")
		}
		body := request.PostFormValue("body")
		uuid := request.PostFormValue("uuid")
		thread, err := data.ThreadByUUID(uuid)
		if err != nil {
			utils.Error_message(writer, request, "Cannot read thread")
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			log.Danger(err, "Cannot create post")
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(writer, request, url, 302)
	}
}
