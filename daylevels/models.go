package daylevels

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/mmarzio67/ml/session"

	"github.com/mmarzio67/ml/config"
)

// DayLevel ... list of physiology levels for a day daylevel"
type DayLevel struct {
	ID              int64
	Focus           int64
	FischioOrecchie int64
	PowerEnergy     int64
	Dormito         int64
	PR              int64
	Ansia           int64
	Arrabiato       int64
	Irritato        int64
	Depresso        int64
	CinqueTib       bool
	Meditazione     bool
	CreatedOn       time.Time
}

// AllDL ... selects all the daylevels from the Database
func AllDL(uid int64) (*[]DayLevel, error) {

	queryAllDL := `SELECT id, 
					focus, 
					fischio_orecchie, 
					power_energy,
					dormito,
					pr,
					ansia,
					arrabiato, 
					irritato,
					depresso, 
					cinque_tibetani,
					meditazione,
					createdon
					FROM daylevels
					WHERE uid=$1`

	rows, err := config.DB.Query(queryAllDL, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	dl := DayLevel{}
	dls := make([]DayLevel, 0)
	for rows.Next() {
		err := rows.Scan(
			&dl.ID,
			&dl.Focus,
			&dl.FischioOrecchie,
			&dl.PowerEnergy,
			&dl.Dormito,
			&dl.PR,
			&dl.Ansia,
			&dl.Arrabiato,
			&dl.Irritato,
			&dl.Depresso,
			&dl.CinqueTib,
			&dl.Meditazione,
			&dl.CreatedOn)

		if err != nil {
			return nil, err
		}
		dls = append(dls, dl)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &dls, nil
}

// OneDL ... selects one daylevel from the Database passed as argument
func OneDL(id int64) (*DayLevel, error) {

	var err error
	dl := DayLevel{}
	fmt.Println(id)

	oneQueryDL := `SELECT id,
	focus, 
	fischio_orecchie, 
	power_energy,
	dormito,
	pr,
	ansia,
	arrabiato, 
	irritato,
	depresso, 
	cinque_tibetani,
	meditazione
	FROM daylevels
	WHERE id=$1`

	row := config.DB.QueryRow(oneQueryDL, id)

	err = row.Scan(
		&dl.ID,
		&dl.Focus,
		&dl.FischioOrecchie,
		&dl.PowerEnergy,
		&dl.Dormito,
		&dl.PR,
		&dl.Ansia,
		&dl.Arrabiato,
		&dl.Irritato,
		&dl.Depresso,
		&dl.CinqueTib,
		&dl.Meditazione)

	if err != nil {
		fmt.Println(err)
		return &dl, err
	}
	return &dl, nil
}

// LastDL ... selects the last daylevel from the Database passed as argument
func LastDL(usrid *int64) (*DayLevel, error) {
	var err error
	ldl := DayLevel{}
	fmt.Println(usrid)

	lastQueryDL := `SELECT id,
	focus, 
	fischio_orecchie, 
	power_energy,
	dormito,
	pr,
	ansia,
	arrabiato, 
	irritato,
	depresso, 
	cinque_tibetani,
	meditazione
	FROM daylevels
	WHERE uid=$1
	ORDER BY id DESC 
	LIMIT 1`

	row := config.DB.QueryRow(lastQueryDL, usrid)

	err = row.Scan(
		&ldl.ID,
		&ldl.Focus,
		&ldl.FischioOrecchie,
		&ldl.PowerEnergy,
		&ldl.Dormito,
		&ldl.PR,
		&ldl.Ansia,
		&ldl.Arrabiato,
		&ldl.Irritato,
		&ldl.Depresso,
		&ldl.CinqueTib,
		&ldl.Meditazione)

	if err != nil {
		fmt.Println(err)
		return &ldl, err
	}
	return &ldl, nil
}

// PutDL ... reads the values from the webform (POST) and create a new entry in the DB
func PutDL(w http.ResponseWriter, r *http.Request) (*DayLevel, error) {
	var err error
	// get form values
	dl := DayLevel{}

	dl.Focus, err = strconv.ParseInt(r.FormValue("focus"), 10, 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("focus entry not accepted")
		fmt.Println(err)
	}

	dl.FischioOrecchie, err = strconv.ParseInt(r.FormValue("fischio_orecchie"), 10, 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("Fischio orecchie entry not accepted")
		fmt.Println(err)
	}

	dl.PowerEnergy, err = strconv.ParseInt(r.FormValue("power_energy"), 10, 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("Power energy entry not accepted")
		fmt.Println(err)
	}

	dl.Dormito, err = strconv.ParseInt(r.FormValue("dormito"), 10, 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("Dormito entry not accepted")
		fmt.Println(err)
	}

	dl.PR, err = strconv.ParseInt(r.FormValue("pr"), 10, 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("Public relations entry not accepted")
		fmt.Println(err)
	}

	dl.Ansia, err = strconv.ParseInt(r.FormValue("ansia"), 10, 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("Ansia entry not accepted")
		fmt.Println(err)
	}

	dl.Arrabiato, err = strconv.ParseInt(r.FormValue("arrabiato"), 10, 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("Arrabiato entry not accepted")
		fmt.Println(err)
	}

	dl.Irritato, err = strconv.ParseInt(r.FormValue("irritato"), 10, 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("irritato entry not accepted")
		fmt.Println(err)
	}

	dl.Depresso, err = strconv.ParseInt(r.FormValue("depresso"), 10, 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("Depresso entry not accepted")
		fmt.Println(err)
	}

	dl.CinqueTib, err = strconv.ParseBool(r.FormValue("cinque_tibetani"))
	if err != nil {
		// handle the error in some way
		fmt.Println("Cinque tibetani entry not accepted")
		fmt.Println(err)
	}

	dl.Meditazione, err = strconv.ParseBool(r.FormValue("meditazione"))
	if err != nil {
		// handle the error in some way
		fmt.Println("Meditazione entry not accepted")
		fmt.Println(err)
	}

	dl.CreatedOn = time.Now()
	fmt.Println(dl.CreatedOn)

	// insert values
	queryDL := `INSERT INTO daylevels (
			focus, 
			fischio_orecchie,
			power_energy,
			dormito,
			pr,
			ansia,
			arrabiato,
			irritato,
			depresso,
			cinque_tibetani,
			meditazione,
			uid) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11,$12)`

	us = session.GetUser(w, r)

	_, err = config.DB.Exec(queryDL,
		&dl.Focus,
		&dl.FischioOrecchie,
		&dl.PowerEnergy,
		&dl.Dormito,
		&dl.PR,
		&dl.Ansia,
		&dl.Arrabiato,
		&dl.Irritato,
		&dl.Depresso,
		&dl.CinqueTib,
		&dl.Meditazione,
		&us.Id)
	if err != nil {
		return &dl, errors.New("500. Internal Server Error." + err.Error())
	}
	//return the newly created ID
	fmt.Println("sono le 9 e tutto va bene")
	lastidrow := config.DB.QueryRow("SELECT max(id) FROM daylevels")
	err = lastidrow.Scan(&dl.ID)
	if err != nil {
		return &dl, err
	}
	fmt.Println(&dl)
	return &dl, nil
}

// UpdateDL ... reads the values from the webform (POST) and updates the daylevel values in the DB
func UpdateDL(r *http.Request) (*DayLevel, error) {
	// get form values
	var err error

	dl := DayLevel{}

	dl.ID, err = strconv.ParseInt(r.FormValue("Id"), 10, 64)

	dl.Focus, err = strconv.ParseInt(r.FormValue("focus"), 10, 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("focus entry not accepted")
		fmt.Println(err)
	}

	dl.FischioOrecchie, err = strconv.ParseInt(r.FormValue("fischio_orecchie"), 10, 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("Fischio orecchie entry not accepted")
		fmt.Println(err)
	}

	dl.PowerEnergy, err = strconv.ParseInt(r.FormValue("power_energy"), 10, 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("Power energy entry not accepted")
		fmt.Println(err)
	}

	dl.Dormito, err = strconv.ParseInt(r.FormValue("dormito"), 10, 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("Dormito entry not accepted")
		fmt.Println(err)
	}

	dl.PR, err = strconv.ParseInt(r.FormValue("pr"), 10, 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("Public relations entry not accepted")
		fmt.Println(err)
	}

	dl.Ansia, err = strconv.ParseInt(r.FormValue("ansia"), 10, 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("Ansia entry not accepted")
		fmt.Println(err)
	}

	dl.Arrabiato, err = strconv.ParseInt(r.FormValue("arrabiato"), 10, 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("Arrabiato entry not accepted")
		fmt.Println(err)
	}

	dl.Irritato, err = strconv.ParseInt(r.FormValue("irritato"), 10, 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("irritato entry not accepted")
		fmt.Println(err)
	}

	dl.Depresso, err = strconv.ParseInt(r.FormValue("depresso"), 10, 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("Depresso entry not accepted")
		fmt.Println(err)
	}

	dl.CinqueTib, err = strconv.ParseBool(r.FormValue("cinque_tibetani"))
	if err != nil {
		// handle the error in some way
		fmt.Println("Cinque tibetani entry not accepted")
		fmt.Println(err)
	}

	dl.Meditazione, err = strconv.ParseBool(r.FormValue("meditazione"))
	if err != nil {
		// handle the error in some way
		fmt.Println("Meditazione entry not accepted")
		fmt.Println(err)
	}

	t := time.Now()
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	dl.CreatedOn = t
	updateQuery := `UPDATE daylevels 
			  SET focus= $1,
			  fischio_orecchie=$2,
			  power_energy=$3,
			  dormito=$4,
			  pr=$5,
			  ansia=$6,
			  arrabiato=$7,
			  irritato=$8,
			  depresso=$9,
			  cinque_tibetani=$10,
			  meditazione=$11
			  WHERE id=$12`

	// insert values
	_, err = config.DB.Exec(updateQuery,
		&dl.Focus,
		&dl.FischioOrecchie,
		&dl.PowerEnergy,
		&dl.Dormito,
		&dl.PR,
		&dl.Ansia,
		&dl.Arrabiato,
		&dl.Irritato,
		&dl.Depresso,
		&dl.CinqueTib,
		&dl.Meditazione,
		&dl.ID)

	if err != nil {
		return &dl, err
	}
	return &dl, nil
}

// DeleteDL ... deletes the daylevel on the DB based on a passed ID
func DeleteDL(r *http.Request) error {
	id := r.FormValue("id")
	if id == "" {
		return errors.New("400. Bad Request")
	}

	_, err := config.DB.Exec("DELETE FROM daylevels WHERE id=$1;", id)
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}
