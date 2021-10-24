package Handle

import (
	"PracticeFC/entities"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type FcService interface {
	CreateMatchCard (mat entities.MatchCard)error
	CreateTrueFalseCard (tf entities.TfCard)error
	CreateStudyCard(sc entities.InfoCard)error
	CreateMultipleCard(mc entities.McCard)error
	CreateQnACard(q entities.QaCard)error
}

type FCHandler struct {
	Svc FcService
}

func NewFcHandler(fs FcService) FCHandler{
	return FCHandler{
		 fs,
	}
}

func (fh FCHandler) CreateCards(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	card := vars["Type"]

	switch card {
	case "Matching":
		match:= entities.MatchCard{}
		err := json.NewDecoder(r.Body).Decode(&match)
		if err != nil {
			log.Fatalln(err)
		}
		err = fh.Svc.CreateMatchCard(match)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	case "Multiple":
		multiple := entities.McCard{}
		err := json.NewDecoder(r.Body).Decode(&multiple)
		if err != nil {
			log.Fatalln(err)
		}
		err = fh.Svc.CreateMultipleCard(multiple)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	case "TF":
		tf := entities.TfCard{}
		err := json.NewDecoder(r.Body).Decode(&tf)
		if err != nil {
			log.Fatalln(err)

		}
		err = fh.Svc.CreateTrueFalseCard(tf)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	case "Info":
		details := entities.InfoCard{}
		err := json.NewDecoder(r.Body).Decode(&details)
		if err != nil {
			log.Fatalln(err)
		}

		err = fh.Svc.CreateStudyCard(details)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	case "QA":
		qna := entities.QaCard{}
		err := json.NewDecoder(r.Body).Decode(&qna)
		if err != nil {
			log.Fatalln(err)
		}

		err = fh.Svc.CreateQnACard(qna)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}