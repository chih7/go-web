package main

import (
	"chih.me/go_web/ChitChat/conf"
	"chih.me/go_web/ChitChat/log"
	"chih.me/go_web/ChitChat/utils"
	"fmt"
	"net/http"
	"net/http/pprof"
	"reflect"
	"runtime"
	"time"
)

func handlelog(h http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		log.Info("Handler function called - " + name)
		h(writer, request)
	}
}

func headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
}

func body(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func main() {

	utils.P("ChitChat", utils.Version(), "started at", conf.Config.Address)

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(conf.Config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", handlelog(index))
	mux.HandleFunc("/err", err)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	server := &http.Server{
		Addr:           conf.Config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(conf.Config.ReadTimeOut * int64(time.Second)),
		WriteTimeout:   time.Duration(conf.Config.WriteTimeOut * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}

	//err := server.ListenAndServeTLS("./conf/cert.pem", "./conf/key.pem")
	//generateSSL()
	server.ListenAndServe()
	if err != nil {
		log.Danger("cannot opem server", err)
	}
}
