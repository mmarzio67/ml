package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const longForm = "Jan 2, 2006 at 3:04pm (MST)"
const hourMinuteForm = "15:04"
const dateForm = "2006-01-02"

type User struct {
	Id       int64
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

// Create a struct that models the structure of a user, both in the request body, and in the DB
type Credentials struct {
	Password string
	Username string
}

type Session struct {
	Un           string
	LastActivity time.Time
}

func SignupAuth(u *User) error {
	// Parse and decode the request body into a new `Credentials` instance
	Password := u.Password

	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as 8 (this value can be more or less, depending on the computing power you wish to utilize)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), 8)

	// Next, insert the username, along with the hashed password into the database
	if _, err = DB.Query("insert into users (user_name, user_pwd, first_name, last_name, idrole) values ($1, $2,$3,$4,$5)", u.UserName, string(hashedPassword), u.First, u.Last, 1); err != nil {
		// If there is any issue with inserting into the database, return a 500 error
		return err
	}

	// We reach this point if the credentials we correctly stored in the database, and the default status of 200 is sent back
	return err
}

func (creds *Credentials) LoginCred() (u *User, e error) {

	var err error
	result := DB.QueryRow("select id, first_name, last_name, user_name, user_pwd, idrole from users where user_name=$1", creds.Username)
	if err != nil {
		fmt.Println("something wrong with the query to the credentials persistance")
		return nil, err
	}
	// We create another instance of `Credentials` to store the credentials we get from the database
	su := &User{}
	// Store the obtained password in `storedCreds`
	err = result.Scan(&su.Id, &su.First, &su.Last, &su.UserName, &su.Password, &su.Role)
	if err != nil {
		// If an entry with the username does not exist, send an "Unauthorized"(401) status
		if err == sql.ErrNoRows {
			fmt.Println("username does not exit")
			return nil, err
		}
		// If the error is of any other type, send a 500 status
		fmt.Println("something wrong with the query to the credentials persistance")
		return nil, err
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword([]byte(su.Password), []byte(creds.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		fmt.Println("password seem do not match")
	}

	// If we reach this point, that means the users password was correct, and that they are authorized
	// The default 200 status is sent
	return su, nil
}

func SigninAuth(w http.ResponseWriter, r *http.Request) {

	// Parse and decode the request body into a new `Credentials` instance
	creds := &Credentials{}
	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		// If there is something wrong with the request body, return a 400 status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Get the existing entry present in the database for the given username
	result := DB.QueryRow("select password from users where username=$1", creds.Username)
	if err != nil {
		// If there is an issue with the database, return a 500 error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// We create another instance of `Credentials` to store the credentials we get from the database
	storedCreds := &Credentials{}
	// Store the obtained password in `storedCreds`
	err = result.Scan(&storedCreds.Password)
	if err != nil {
		// If an entry with the username does not exist, send an "Unauthorized"(401) status
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// If the error is of any other type, send a 500 status
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(creds.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
	}

	// If we reach this point, that means the users password was correct, and that they are authorized
	// The default 200 status is sent
}

func GetHourMinuteSecond(t *time.Time) *string {

	layout := "15:04:05"
	hsm := t.Format(layout)
	return &hsm

}

// candidate to be embedded in a interface (nicer)
func NiceDate(d time.Time) string {

	nd := d.Format(dateForm)
	fmt.Println(nd)
	return nd
}

// candidate to be embedded in a interface (nicer)
func NiceTime(t time.Time) string {

	nt := t.Format(hourMinuteForm)
	fmt.Println(nt)
	return nt

}
