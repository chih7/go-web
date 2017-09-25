package utils

import (
	"net/http"
	"fmt"
	"html/template"
	"chih.me/go_web/ChitChat/data"
	"errors"
	"strings"
)

func P(a ...interface{}) {
	fmt.Println(a)
}

func Error_message(w http.ResponseWriter, r *http.Request, msg string) {
	url := []string{
		"err?msg=",
		msg,
	}
	http.Redirect(w, r, strings.Join(url, ""), 302)
}

func Session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {

	cookie, err := r.Cookie("_cookie")
	if err == nil {
		sess = data.Session{
			Uuid: cookie.Value,
		}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

func ParseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

func GenerateHTML(w http.ResponseWriter, data interface{}, filenames ... string) {
	var files []string
	// variadic
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, "layout", data)
}

// version
func Version() string {
	return "0.1"
}
