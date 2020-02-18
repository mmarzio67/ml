package valvita

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/mmarzio67/ml/config"
	"github.com/mmarzio67/ml/session"
)

var us config.User

func Index(w http.ResponseWriter, r *http.Request) {
	if !session.AlreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	us = session.GetUser(w, r)
	vvs, err := AllVv(us.Id)
	if err != nil {
		http.Redirect(w, r, "/val/create", http.StatusSeeOther)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "valorivitali.html", vvs)
}

func Create(w http.ResponseWriter, r *http.Request) {
	if !session.AlreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/val", http.StatusSeeOther)
		return
	}

	us = session.GetUser(w, r)
	lvv, err := LastVv(&us.Id)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "createvv.html", lvv)
}

func CreateProcess(w http.ResponseWriter, r *http.Request) {
	if !session.AlreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/val", http.StatusSeeOther)
		return
	}

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	vv, err := PutVv(w, r)
	fmt.Println(vv)

	if err != nil {
		println("error in processing PutVv")
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}

	config.TPL.ExecuteTemplate(w, "createdvv.html", vv)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if !session.AlreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'Id' is missing")
		return
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.

	id, err := strconv.ParseInt(keys[0], 10, 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("id parameter reading accepted")
		fmt.Println(err)
	}

	vv, err := OneVv(id)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "updatevv.html", vv)
}

func UpdateProcess(w http.ResponseWriter, r *http.Request) {
	if !session.AlreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	vv, err := UpdateVv(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusBadRequest)
		return
	}

	config.TPL.ExecuteTemplate(w, "updatedvv.html", vv)
}

func DeleteProcess(w http.ResponseWriter, r *http.Request) {
	if !session.AlreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	err := DeleteVv(r)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/val", http.StatusSeeOther)
}
