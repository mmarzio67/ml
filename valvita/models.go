package valvita

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/mmarzio67/ml/config"
	"github.com/mmarzio67/ml/session"
)

type ValVita struct {
	Id          int64
	Peso        float64
	Girovita    float64
	CreatedOn   time.Time
	Grassocorpo float64
	Imc         float64
	Pulse       int64
}

var uid int64

func AllVv(uid int64) (*[]ValVita, error) {

	queryAllVv := `SELECT id, 
					peso, 
					girovita,
					createdon,
					grassocorpo,
					imc,
					pulse
					FROM misurevitali
					WHERE uid=$1`

	rows, err := config.DB.Query(queryAllVv, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	vv := ValVita{}
	vvs := make([]ValVita, 0)
	for rows.Next() {
		err := rows.Scan(
			&vv.Id,
			&vv.Peso,
			&vv.Girovita,
			&vv.CreatedOn,
			&vv.Grassocorpo,
			&vv.Imc,
			&vv.Pulse)

		if err != nil {
			return nil, err
		}
		vvs = append(vvs, vv)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &vvs, nil
}

func OneVv(id int64) (*ValVita, error) {

	var err error
	vv := ValVita{}
	fmt.Println(id)

	oneQueryVv := `SELECT id,
	            	peso, 
					girovita,
					grassocorpo,
					imc,
					pulse
	            	FROM misurevitali
	            	WHERE id=$1`

	row := config.DB.QueryRow(oneQueryVv, id)

	err = row.Scan(
		&vv.Id,
		&vv.Peso,
		&vv.Girovita,
		&vv.Grassocorpo,
		&vv.Imc,
		&vv.Pulse)

	if err != nil {
		fmt.Println(err)
		return &vv, err
	}
	return &vv, nil
}

func LastVv(usrid *int64) (*ValVita, error) {
	var err error
	lvv := ValVita{}

	lastQueryVv := `SELECT id,
	peso, 
	girovita,
	grassocorpo,
	imc,
	pulse
	FROM misurevitali
	WHERE uid=$1
	ORDER BY id DESC 
	LIMIT 1`

	row := config.DB.QueryRow(lastQueryVv, usrid)

	err = row.Scan(
		&lvv.Id,
		&lvv.Peso,
		&lvv.Girovita,
		&lvv.Grassocorpo,
		&lvv.Imc,
		&lvv.Pulse)

	if err != nil {
		fmt.Println(err)
		return &lvv, err
	}
	return &lvv, nil
}

func PutVv(w http.ResponseWriter, r *http.Request) (*ValVita, error) {
	var err error
	// get form values
	vv := ValVita{}

	vv.Peso, err = strconv.ParseFloat(r.FormValue("peso"), 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("misura peso immessa non accettata")
		fmt.Println(err)
	}

	vv.Girovita, err = strconv.ParseFloat(r.FormValue("girovita"), 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("misura girovita immessa non accettata")
		fmt.Println(err)
	}

	vv.Grassocorpo, err = strconv.ParseFloat(r.FormValue("grassocorpo"), 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("grasso corporeo entry not accepted")
		fmt.Println(err)
	}

	vv.Imc, err = strconv.ParseFloat(r.FormValue("imc"), 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("imc entry not accepted")
		fmt.Println(err)
	}

	vv.Pulse, err = strconv.ParseInt(r.FormValue("pulse"), 10, 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("pulse entry not accepted")
		fmt.Println(err)
	}

	vv.CreatedOn = time.Now()
	fmt.Println(vv.CreatedOn)

	// insert values
	queryVv := `INSERT INTO misurevitali (
			peso, 
			girovita,
			grassocorpo,
			imc,
			pulse,
			uid) 
			VALUES ($1, $2, $3, $4, $5, $6)`

	us = session.GetUser(w, r)
	_, err = config.DB.Exec(queryVv,
		&vv.Peso,
		&vv.Girovita,
		&vv.Grassocorpo,
		&vv.Imc,
		&vv.Pulse,
		&us.Id)

	if err != nil {
		return &vv, errors.New("500. Internal Server Error." + err.Error())
	}
	//return the newly created ID
	fmt.Println("sono le 9 e tutto va bene")
	lastidrow := config.DB.QueryRow("SELECT max(id) FROM misurevitali")
	err = lastidrow.Scan(&vv.Id)
	if err != nil {
		return &vv, err
	}
	fmt.Println(&vv)
	return &vv, nil
}

func UpdateVv(r *http.Request) (*ValVita, error) {
	// get form values
	var err error

	vv := ValVita{}

	vv.Id, err = strconv.ParseInt(r.FormValue("Id"), 10, 64)

	vv.Peso, err = strconv.ParseFloat(r.FormValue("peso"), 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("peso entry not accepted")
		fmt.Println(err)
	}

	vv.Girovita, err = strconv.ParseFloat(r.FormValue("girovita"), 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("girovita entry not accepted")
		fmt.Println(err)
	}

	vv.Grassocorpo, err = strconv.ParseFloat(r.FormValue("grassocorpo"), 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("grasso corporeo entry not accepted")
		fmt.Println(err)
	}

	vv.Imc, err = strconv.ParseFloat(r.FormValue("imc"), 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("imc entry not accepted")
		fmt.Println(err)
	}

	vv.Pulse, err = strconv.ParseInt(r.FormValue("pulse"), 10, 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("pulse entry not accepted")
		fmt.Println(err)
	}

	t := time.Now()
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	vv.CreatedOn = t
	updateQueryVv := `UPDATE misurevitali
			  SET peso= $1,
			  girovita=$2,
			  grassocorpo=$3,
			  imc=$4,
			  pulse=$5
			  WHERE id=$6`

	// insert values
	_, err = config.DB.Exec(updateQueryVv,
		&vv.Peso,
		&vv.Girovita,
		&vv.Grassocorpo,
		&vv.Imc,
		&vv.Pulse,
		&vv.Id)

	if err != nil {
		return &vv, err
	}
	return &vv, nil
}

func DeleteVv(r *http.Request) error {
	var err error
	var id int64
	id, err = strconv.ParseInt(r.FormValue("id"), 10, 64)

	if err != nil {
		// handle the error in some way
		fmt.Println("misura girovita immessa non accettata")
		fmt.Println(err)
	}

	_, err = config.DB.Exec("DELETE FROM misurevitali WHERE id=$1;", id)
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}
