package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/robotuimyhorobotuiotui/core/db"
	dbserver "github.com/robotuimyhorobotuiotui/server/db"
	"github.com/robotuimyhorobotuiotui/util/locale"
	"github.com/thommyho/robotui/api"
	"golang.org/x/text/language"
)

func csvResult(ctx context.Context, w http.ResponseWriter, res any, filename string) {
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", `attachment; filename="`+filename+`.csv"`)

	if ww, ok := res.(api.CsvWriter); ok {
		_ = ww.WriteCsv(ctx, w)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// sessionHandler returns the list of charging sessions
func sessionHandler(w http.ResponseWriter, r *http.Request) {
	if dbserver.Instance == nil {
		jsonError(w, http.StatusBadRequest, errors.New("database offline"))
		return
	}

	var res db.Sessions
	year := r.URL.Query().Get("year")
	month := r.URL.Query().Get("month")

	filename := "session"

	fmtYear := "%"
	fmtMonth := "%"

	if year != "" {
		fmtYear = year
		filename += "-" + fmtYear

		if month != "" {
			iMonth, err := strconv.Atoi(month)
			if err != nil {
				jsonError(w, http.StatusBadRequest, err)
				return
			}

			fmtMonth = fmt.Sprintf("%02d", iMonth)
			filename += "." + fmtMonth
		}
	}

	// TODO support other databases than Sqlite
	whereQuery := "charged_kwh>=0.05 AND strftime('%Y', created) LIKE ? AND strftime('%m', created) LIKE ?"
	if txn := dbserver.Instance.Where(whereQuery, fmtYear, fmtMonth).Order("created DESC").Find(&res); txn.Error != nil {
		jsonError(w, http.StatusInternalServerError, txn.Error)
		return
	}

	// prepare data
	for i, s := range res {
		if s.Odometer != nil {
			odo := math.Round(*s.Odometer*10) / 10
			res[i].Odometer = &odo
		}
	}

	if r.URL.Query().Get("format") == "csv" {
		lang := r.URL.Query().Get("lang")
		if lang == "" {
			// get request language
			lang = r.Header.Get("Accept-Language")
			if tags, _, err := language.ParseAcceptLanguage(lang); err == nil && len(tags) > 0 {
				lang = tags[0].String()
			}
		}

		ctx := context.WithValue(context.Background(), locale.Locale, lang)
		csvResult(ctx, w, &res, filename)
		return
	}

	jsonResult(w, res)
}

// deleteSessionHandler removes session in sessions table with given id
func deleteSessionHandler(w http.ResponseWriter, r *http.Request) {
	if dbserver.Instance == nil {
		jsonError(w, http.StatusBadRequest, errors.New("database offline"))
		return
	}

	var res db.Sessions

	vars := mux.Vars(r)
	id := vars["id"]

	if txn := dbserver.Instance.Table("sessions").Delete(&res, id); txn.Error != nil {
		jsonError(w, http.StatusBadRequest, txn.Error)
		return
	}

	jsonResult(w, res)
}

// updateSessionHandler updates the data of an existing session
func updateSessionHandler(w http.ResponseWriter, r *http.Request) {
	if dbserver.Instance == nil {
		jsonError(w, http.StatusBadRequest, errors.New("database offline"))
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	var session map[string]any
	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		jsonError(w, http.StatusBadRequest, err)
		return
	}

	if txn := dbserver.Instance.Table("sessions").Where("id = ?", id).Updates(&session); txn.Error != nil {
		jsonError(w, http.StatusBadRequest, txn.Error)
		return
	}
}
