package app

import (
	"net/http"

	"test/go/src/github.com/lambogreny/gonews/pkg/view"
)

func adminLogin(w http.ResponseWriter, r *http.Request) {
	view.AdminLogin(w, nil)
}

func adminList(w http.ResponseWriter, r *http.Request) {

}

func adminCreate(w http.ResponseWriter, r *http.Request) {

}

func adminEdit(w http.ResponseWriter, r *http.Request) {

}
