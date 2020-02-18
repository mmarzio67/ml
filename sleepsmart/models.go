package sleepsmart

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/mmarzio67/ml/config"
	"github.com/mmarzio67/ml/session"
)

const longForm = "Jan 2, 2006 at 3:04pm (MST)"
const hourMinuteForm = "15:04"
const dateForm = "2006-01-02"

type SleepSmart struct {
	ID                int64
	DSleep            time.Time
	HSleep            time.Time
	DWake             time.Time
	HWake             time.Time
	HAsleep           float64
	NrCycles          float64
	CyclesInterrupted bool
	Alcool            bool
	Coffee            int64
	LastCoffee        time.Time
	Diner             int64
	PrepSleep         int64
	BlueLight         int64
	MelatoninLevel    int64
	CreatedOn         time.Time
}

var uid int64

func AllSs(uid int64) (*[]SleepSmart, error) {

	queryAllSs := `SELECT id,
					datadormire,
					hsleep,
					datasveglia,
    				hwake,
					hasleep,
					nrCycles,
					cyclesInterrupted,
					alcool,
					coffee,
					lastCoffe,
					diner,
					prepSleep,
					blueLight,
					melatoninLevel
					FROM sleepsmart
					WHERE uid=$1`

	rows, err := config.DB.Query(queryAllSs, uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ss := SleepSmart{}
	sss := make([]SleepSmart, 0)

	for rows.Next() {
		err := rows.Scan(
			&ss.ID,
			/*
				config.NiceDate(ss.DSleep),
				config.NiceTime(ss.HSleep),
				config.NiceDate(ss.DWake),
				config.NiceTime(ss.HWake),
			*/
			&ss.DSleep,
			&ss.HSleep,
			&ss.DWake,
			&ss.HWake,
			&ss.HAsleep,
			&ss.NrCycles,
			&ss.CyclesInterrupted,
			&ss.Alcool,
			&ss.Coffee,
			&ss.LastCoffee,
			&ss.Diner,
			&ss.PrepSleep,
			&ss.BlueLight,
			&ss.MelatoninLevel,
		)

		if err != nil {
			return nil, err
		}
		sss = append(sss, ss)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &sss, nil
}

func OneSs(id int64) (*SleepSmart, error) {

	var err error
	ss := SleepSmart{}
	fmt.Println(id)

	oneQuerySs := `SELECT id,
					datadormire,
					hsleep,
					datasveglia,
					hwake,
					hasleep,
					nrCycles,
					cyclesInterrupted,
					alcool,
					coffee,
					lastCoffe,
					diner,
					prepSleep,
					blueLight,
					melatoninLevel
					FROM sleepsmart
	            	WHERE id=$1`

	row := config.DB.QueryRow(oneQuerySs, id)

	err = row.Scan(
		&ss.ID,
		&ss.DSleep,
		&ss.HSleep,
		&ss.DWake,
		&ss.HWake,
		&ss.HAsleep,
		&ss.NrCycles,
		&ss.CyclesInterrupted,
		&ss.Alcool,
		&ss.Coffee,
		&ss.LastCoffee,
		&ss.Diner,
		&ss.PrepSleep,
		&ss.BlueLight,
		&ss.MelatoninLevel,
	)

	// formats the dates without time and time without dates
	// before to send the struct to the template

	formatDateHours(&ss)

	if err != nil {
		fmt.Println(err)
		return &ss, err
	}
	return &ss, nil
}

func LastSs(usrid *int64) (*SleepSmart, error) {
	var err error
	ss := SleepSmart{}

	lastQuerySs := `SELECT id,
					datadormire,
					hsleep,
					datasveglia,
					hwake,
					hasleep,
					nrCycles,
					cyclesInterrupted,
					alcool,
					coffee,
					lastCoffe,
					diner,
					prepSleep,
					blueLight,
					melatoninLevel
					FROM sleepsmart
					WHERE uid=$1
					ORDER BY id DESC 
					LIMIT 1`

	row := config.DB.QueryRow(lastQuerySs, usrid)

	// format HH:MM:SS for a type time.Time ready to insert into the database

	err = row.Scan(
		&ss.ID,
		&ss.DSleep,
		&ss.HSleep,
		&ss.DWake,
		&ss.HWake,
		&ss.HAsleep,
		&ss.NrCycles,
		&ss.CyclesInterrupted,
		&ss.Alcool,
		&ss.Coffee,
		&ss.LastCoffee,
		&ss.Diner,
		&ss.PrepSleep,
		&ss.BlueLight,
		&ss.MelatoninLevel,
	)

	if err != nil {
		fmt.Println(err)
		return &ss, err
	}
	return &ss, nil
}

func PutSs(w http.ResponseWriter, r *http.Request) (*SleepSmart, error) {

	var err error
	// get form values
	ss := SleepSmart{}
	marshalFormEntries(&ss, r)
	// insert values
	querySs := `INSERT INTO sleepsmart (
				datadormire,
				hsleep,
				datasveglia,
				hwake,
				hasleep,
				nrCycles,
				cyclesInterrupted,
				alcool,
				coffee,
				lastCoffe,
				diner,
				prepSleep,
				blueLight,
				melatoninLevel,
				uid) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`

	us = session.GetUser(w, r)

	_, err = config.DB.Exec(querySs,
		&ss.DSleep,
		&ss.HSleep,
		&ss.DWake,
		&ss.HWake,
		&ss.HAsleep,
		&ss.NrCycles,
		&ss.CyclesInterrupted,
		&ss.Alcool,
		&ss.Coffee,
		&ss.LastCoffee,
		&ss.Diner,
		&ss.PrepSleep,
		&ss.BlueLight,
		&ss.MelatoninLevel,
		&us.Id)

	if err != nil {
		return &ss, errors.New("500. Internal Server Error." + err.Error())
	}
	//return the newly created ID
	fmt.Println("sono le 9 e tutto va bene")
	lastidrow := config.DB.QueryRow("SELECT max(id) FROM sleepsmart")
	err = lastidrow.Scan(&ss.ID)
	if err != nil {
		return &ss, err
	}
	fmt.Println(&ss)
	return &ss, nil
}

func UpdateSs(r *http.Request) (*SleepSmart, error) {

	var err error
	ss := SleepSmart{}
	marshalFormEntries(&ss, r)

	updateQuerySs := `UPDATE sleepsmart
			  SET datadormire=$1,
				hsleep=$2,
				datasveglia=$3, 
				hwake=$4,
				hasleep=$5,
				nrCycles=$6,
				cyclesInterrupted=$7,
				alcool=$8,
				coffee=$9,
				lastCoffe=$10,
				diner=$11,
				prepSleep=$12,
				blueLight=$13,
				melatoninLevel=$14
			  	WHERE id=$15`

	// insert values
	_, err = config.DB.Exec(updateQuerySs,
		&ss.DSleep,
		&ss.HSleep,
		&ss.DWake,
		&ss.HWake,
		&ss.HAsleep,
		&ss.NrCycles,
		&ss.CyclesInterrupted,
		&ss.Alcool,
		&ss.Coffee,
		&ss.LastCoffee,
		&ss.Diner,
		&ss.PrepSleep,
		&ss.BlueLight,
		&ss.MelatoninLevel,
		&ss.ID)

	if err != nil {
		return &ss, err
	}
	return &ss, nil
}

func DeleteSs(r *http.Request) error {
	var err error
	var id int64
	id, err = strconv.ParseInt(r.FormValue("id"), 10, 64)

	if err != nil {
		// handle the error in some way
		fmt.Println("impossibile cancellare il dato")
		fmt.Println(err)
	}

	_, err = config.DB.Exec("DELETE FROM sleepsmart WHERE id=$1;", id)
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}

func calcSleepCycles(hd time.Time, hs time.Time) float64 {

	oreSonno, _ := time.ParseDuration(oreDormite(hd, hs))
	var durataCicloMin float64
	durataCicloMin = 90
	nrCicli := oreSonno.Minutes() / durataCicloMin
	return nrCicli
}

func oreDormite(hd time.Time, hs time.Time) string {

	diff := hs.Sub(hd)
	out := (time.Time{}.Add(diff)).Format("15:04:05")
	return out
}

func marshalFormEntries(ss *SleepSmart, r *http.Request) {
	// get form values
	var err error

	ss.DSleep, err = time.Parse(dateForm, r.FormValue("dsleep"))
	if err != nil {
		// handle the error in some way
		fmt.Println("data di andare a dormire non valida")
		fmt.Println(err)
	}

	ss.HSleep, err = time.Parse(hourMinuteForm, r.FormValue("hsleep"))
	if err != nil {
		// handle the error in some way
		fmt.Println("ora di andare a dormire non valida")
		fmt.Println(err)
	}

	ss.DWake, err = time.Parse(dateForm, r.FormValue("dwake"))
	if err != nil {
		// handle the error in some way
		fmt.Println("data di sveglia non valida")
		fmt.Println(err)
	}

	ss.HWake, err = time.Parse(hourMinuteForm, r.FormValue("hwake"))
	if err != nil {
		// handle the error in some way
		fmt.Println("ora sveglia non valida")
		fmt.Println(err)
	}

	ss.HAsleep, _ = strconv.ParseFloat(r.FormValue("hasleep"), 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("orario non accettato")
		fmt.Println(err)
	}

	// questa funzione deve essere trasformata in go routine e il valore di ritorno in channel
	//ss.NrCycles = calcSleepCycles(ss.HSleep, ss.HWake)
	ss.NrCycles = 4

	ss.CyclesInterrupted, err = strconv.ParseBool(r.FormValue("cyclesInterrupted"))
	if err != nil {
		// handle the error in some way
		fmt.Println("valore cicli sonno non accettato")
		fmt.Println(err)
	}

	ss.Alcool, err = strconv.ParseBool(r.FormValue("alcool"))
	if err != nil {
		// handle the error in some way
		fmt.Println("valore alcool non accettato")
		fmt.Println(err)
	}

	ss.Coffee, err = strconv.ParseInt(r.FormValue("coffee"), 10, 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("imc entry not accepted")
		fmt.Println(err)
	}

	ss.LastCoffee, err = time.Parse(hourMinuteForm, r.FormValue("lastcoffee"))
	if err != nil {
		// handle the error in some way
		fmt.Println("ora di andare a dormire non valida")
		fmt.Println(err)
	}

	ss.Diner, err = strconv.ParseInt(r.FormValue("diner"), 10, 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("diner entry not accepted")
		fmt.Println(err)
	}

	ss.PrepSleep, err = strconv.ParseInt(r.FormValue("prepSleep"), 10, 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("prepSleep level entry not accepted")
		fmt.Println(err)
	}

	ss.BlueLight, err = strconv.ParseInt(r.FormValue("blueLight"), 10, 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("bluelight level entry not accepted")
		fmt.Println(err)
	}

	ss.MelatoninLevel, err = strconv.ParseInt(r.FormValue("melatoninLevel"), 10, 64)
	if err != nil {
		// handle the error in some way
		fmt.Println("melatonin level entry not accepted")
		fmt.Println(err)
	}

	t := time.Now()
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	ss.CreatedOn = t
}

func formatDateHours(ss *SleepSmart) *SleepSmart {
	/* take the date and time values of the scanned struct
	   and format them without unnecessary values like hours
	   in the dates and dates in the hours and return them
	   before to send the struct to the template
	*/
	ss.DSleep, _ = time.Parse(dateForm, ss.DSleep.Format(dateForm))
	ss.DWake, _ = time.Parse(dateForm, ss.DWake.Format(dateForm))
	ss.HSleep, _ = time.Parse(hourMinuteForm, ss.HSleep.Format(hourMinuteForm))
	ss.HWake, _ = time.Parse(hourMinuteForm, ss.HWake.Format(hourMinuteForm))
	ss.LastCoffee, _ = time.Parse(hourMinuteForm, ss.LastCoffee.Format(hourMinuteForm))
	return ss
}
