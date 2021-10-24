package Repo

import (
	"PracticeFC/entities"
	"encoding/json"
	"io/ioutil"
)

type FlashcardDb struct {
	Flashcard []interface{}
}

type Repo struct {
	Filename string
}

func NewFcRepo(f string) Repo {
	return Repo{
		Filename: f,
	}
}

func (r Repo) CreateMatchCard(mat entities.MatchCard) error {
	mtdb := FlashcardDb{}

	file, err := ioutil.ReadFile(r.Filename) // reads the file moviedb.json
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &mtdb) // pointer indicates change and points to address of the instance of the Flashcard Struct
	if err != nil {
		return err
	}

	mtdb.Flashcard = append(mtdb.Flashcard, mat) //appends slice of flashcards to the instance of mvDB struct

	fcBytes, err := json.MarshalIndent(mtdb, "", "	") // takes in the bytes and converts back to json
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, fcBytes, 0644) //writes the information to the database file
	if err != nil {
		return err
	}

	return nil

}

func (r Repo) CreateTrueFalseCard (tf entities.TfCard) error {
	tdb := FlashcardDb{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &tdb)
	if err != nil {
		return err
	}

	tdb.Flashcard = append(tdb.Flashcard, tf)

	fcBytes, err := json.MarshalIndent(tdb, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, fcBytes, 0644)
	if err != nil {
		return err
	}

	return nil

}

func (r Repo) CreateStudyCard(sc entities.InfoCard) error {
	sdb := FlashcardDb{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &sdb)
	if err != nil {
		return err
	}

	sdb.Flashcard = append(sdb.Flashcard, sc)

	fcBytes, err := json.MarshalIndent(sdb, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, fcBytes, 0644)
	if err != nil {
		return err
	}

	return nil

}

func (r Repo) CreateMultipleCard(mc entities.McCard) error {
	mdb := FlashcardDb{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &mdb)
	if err != nil {
		return err
	}

	mdb.Flashcard = append(mdb.Flashcard, mc)

	fcBytes, err := json.MarshalIndent(mdb, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, fcBytes, 0644)
	if err != nil {
		return err
	}

	return nil

}

func (r Repo) CreateQnACard(q entities.QaCard) error {
	odb := FlashcardDb{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &odb)
	if err != nil {
		return err
	}

	odb.Flashcard = append(odb.Flashcard, q)

	fcBytes, err := json.MarshalIndent(odb, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, fcBytes, 0644)
	if err != nil {
		return err
	}

	return nil

}

