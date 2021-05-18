package session

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mmarzio67/ml/mlogger"

	_ "github.com/lib/pq"
	"github.com/mmarzio67/ml/config"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var Cusr *config.User

func Bar(w http.ResponseWriter, req *http.Request) {
	u := GetUser(w, req)
	if !AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	showSessions() // for demonstration purposes
	config.TPL.ExecuteTemplate(w, "bar.html", u)
}

func Signup(w http.ResponseWriter, req *http.Request) {

	if AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var us config.User
	var id int64

	// POST request JSON Data
	decoder := json.NewDecoder(req.Body)
	sww := decoder.Decode(&us)
	if sww != nil {
		panic(sww)
	}
	log.Println(us.UserName)

	un := us.UserName
	p := us.Password
	f := us.First
	l := us.Last
	r := us.Role

	bs := []byte(p)
	us = config.User{id, un, bs, f, l, r}

	usertaken := config.SignupAuth(&us)

	if usertaken != nil {
		fmt.Println(usertaken)
		return
	}

	// create session
	sID := uuid.NewV4()
	c := &http.Cookie{
		Name:  "session",
		Value: sID.String(),
	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	config.DbSessions[c.Value] = config.Session{un, time.Now()}
	// store User in dbUsers
	bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	us = config.User{id, un, bs, f, l, r}
	config.DbUsers[un] = us
}

func Login(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	if AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	var pc config.Credentials
	var errAuth error

	// process form submission
	if req.Method == http.MethodPost {
		uf := req.FormValue("Username")
		p := req.FormValue("password")

		pc = config.Credentials{p, uf}
		Cusr, errAuth = pc.LoginCred()
		if errAuth != nil {
			http.Error(w, "Something wrong with the user authentication", http.StatusForbidden)
			return
		}

		config.DbUsers[uf] = *Cusr
		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(Cusr.Password, []byte(p))
		if err != nil {
			fmt.Println(Cusr.Password)
			http.Error(w, "Username and/or password do not match, dude", http.StatusForbidden)
			return
		}

		//log the successful login

		logger := mlogger.GetInstance()
		logger.Printf("Login successful for %s\n", Cusr.UserName)

		// create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		c.MaxAge = sessionLength
		http.SetCookie(w, c)
		config.DbSessions[c.Value] = config.Session{uf, time.Now()}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return

	}
	showSessions() // for demonstration purposes
	config.TPL.ExecuteTemplate(w, "login.html", Cusr)
}

func Logout(w http.ResponseWriter, req *http.Request) {
	if !AlreadyLoggedIn(w, req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session")
	// delete the session
	delete(config.DbSessions, c.Value)
	// remove the cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	// clean up dbSessions
	if time.Now().Sub(DbSessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}

func FirstApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Welcome to the first pure flapigo API"}`))
}

func AlreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := config.DbSessions[c.Value]
	if ok {
		s.LastActivity = time.Now()
		config.DbSessions[c.Value] = s
	}
	_, ok = config.DbUsers[s.Un]
	// refresh session
	c.MaxAge = sessionLength
	http.SetCookie(w, c)
	return ok
}
