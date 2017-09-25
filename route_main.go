package main

import (
	"net/http"
	"chih.me/go_web/ChitChat/data"
	"chih.me/go_web/ChitChat/utils"
)

func err(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	_, err := utils.Session(w, r)
	if err != nil {
		utils.GenerateHTML(w, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		utils.GenerateHTML(w, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func index(w http.ResponseWriter, r *http.Request) {

	threads, err := data.Threads()
	if err != nil {
		utils.Error_message(w, r, "Cannot get threads")
	} else {
		_, err := utils.Session(w, r)

		if err != nil {
			utils.GenerateHTML(w, threads, "layout", "public.navbar", "index")
		} else {
			utils.GenerateHTML(w, threads, "layout", "private.navbar", "index")
		}
	}
}
