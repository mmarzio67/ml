package session

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mmarzio67/ml/config"
	"github.com/satori/go.uuid"
)

var DbSessionsCleaned time.Time
var u config.User

const sessionLength int = 300

func GetUser(w http.ResponseWriter, req *http.Request) config.User {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}

	}
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	// if the User exists already, get User
	if s, ok := config.DbSessions[c.Value]; ok {
		s.LastActivity = time.Now()
		config.DbSessions[c.Value] = s
		u = config.DbUsers[s.Un]
	}
	return u
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

func cleanSessions() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	showSessions()              // for demonstration purposes
	for k, v := range config.DbSessions {
		if time.Now().Sub(v.LastActivity) > (time.Second * 30) {
			delete(config.DbSessions, k)
		}
	}
	DbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	showSessions()             // for demonstration purposes
}

// for demonstration purposes
func showSessions() {
	fmt.Println("********")
	for k, v := range config.DbSessions {
		fmt.Println(k, v.Un)
	}
	fmt.Println("")
}
