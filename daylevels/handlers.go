package daylevels

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/mmarzio67/ml/config"
	"github.com/mmarzio67/ml/session"
)

var us config.User

// Index ... lists all the daylevels thanks to the function AllDL
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

	dls, err := AllDL(us.Id)
	if err != nil {
		http.Redirect(w, r, "/dls/create", http.StatusSeeOther)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "daylevels.html", dls)
}

// Create ... lists the last entry daylevel created thanks to the function LastDL
func Create(w http.ResponseWriter, r *http.Request) {
	if !session.AlreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	us = session.GetUser(w, r)
	ldl, err := LastDL(&us.Id)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "create.html", ldl)
}

// CreateProcess ... creates a new daylevel entry thanks to the function LastDL
func CreateProcess(w http.ResponseWriter, r *http.Request) {
	if !session.AlreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	dl, err := PutDL(w, r)
	fmt.Println(dl)

	if err != nil {
		println("error in processing PutDL")
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}

	config.TPL.ExecuteTemplate(w, "created.html", dl)
}

// Update ... selects a single daylevel that was selected in the webpage thanks to the function OneDL
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

	dl, err := OneDL(id)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	config.TPL.ExecuteTemplate(w, "update.html", dl)
}

// UpdateProcess ... updates daylevels values from a daylevel passed argument thanks to the function UpdateDL
func UpdateProcess(w http.ResponseWriter, r *http.Request) {
	if !session.AlreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	dl, err := UpdateDL(r)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusBadRequest)
		return
	}

	config.TPL.ExecuteTemplate(w, "updated.html", dl)
}

// DeleteProcess ... deletes a daylevels from DB thanks to the function DeleteDL
func DeleteProcess(w http.ResponseWriter, r *http.Request) {
	if !session.AlreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	err := DeleteDL(r)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/daylevels", http.StatusSeeOther)
}

func GetAllDLAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	us = session.GetUser(w, r)

	// For test puposes: write some tests dude!!!!
	//dls, err := AllDL(2)

	// PROD Code
	dls, err := AllDL(us.Id)
	if err != nil {
		http.Redirect(w, r, "/dls/create", http.StatusSeeOther)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	jsonDL, err := json.MarshalIndent(dls, "", "   ")

	w.Write(jsonDL)
}
