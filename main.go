package main

import (
	"net/http"
	"os"

	"github.com/mmarzio67/ml/daylevels"
	"github.com/mmarzio67/ml/mlogger"
	"github.com/mmarzio67/ml/session"
	"github.com/mmarzio67/ml/sleepsmart"
	"github.com/mmarzio67/ml/valvita"
)

func main() {

	port := getenv("PORT", "5900")

	logger := mlogger.GetInstance()
	logger.Println("Starting Medaliving Web Service")

	http.HandleFunc("/", daylevels.Index)

	// session handlers
	http.HandleFunc("/firstapi", session.FirstApi)
	http.HandleFunc("/signup", session.Signup)
	http.HandleFunc("/login", session.Login)
	http.HandleFunc("/logout", session.Logout)

	// daylevels handlers
	http.HandleFunc("/dls", daylevels.Index)
	http.HandleFunc("/dls/daylevels", daylevels.GetAllDLAPI)
	http.HandleFunc("/dls/create", daylevels.Create)
	http.HandleFunc("/dls/create/process", daylevels.CreateProcess)
	http.HandleFunc("/dls/update", daylevels.Update)
	http.HandleFunc("/dls/update/process", daylevels.UpdateProcess)
	http.HandleFunc("/dls/delete/process", daylevels.DeleteProcess)

	// valori vitali handlers
	http.HandleFunc("/val", valvita.Index)
	http.HandleFunc("/val/create", valvita.Create)
	http.HandleFunc("/val/create/process", valvita.CreateProcess)
	http.HandleFunc("/val/update", valvita.Update)
	http.HandleFunc("/val/update/process", valvita.UpdateProcess)
	http.HandleFunc("/val/delete/process", valvita.DeleteProcess)

	// valori igiene sonno
	http.HandleFunc("/slp", sleepsmart.Index)
	http.HandleFunc("/slp/create", sleepsmart.Create)
	http.HandleFunc("/slp/create/process", sleepsmart.CreateProcess)
	http.HandleFunc("/slp/update", sleepsmart.Update)
	http.HandleFunc("/slp/update/process", sleepsmart.UpdateProcess)
	http.HandleFunc("/slp/delete/process", sleepsmart.DeleteProcess)

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":"+port, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/dls", http.StatusSeeOther)
}

func getenv(k string, v string) string {
	if val := os.Getenv(k); val != "" {
		return val
	}
	os.Setenv(k, v)
	return v
}
