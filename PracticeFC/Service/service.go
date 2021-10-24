package Service

import (
	"PracticeFC/Repo"
	"PracticeFC/entities"
	"github.com/google/uuid"
)

type FcRepo interface {
	CreateMatchCard (mat entities.MatchCard)error
	CreateTrueFalseCard (tf entities.TfCard)error
	CreateStudyCard(sc entities.InfoCard)error
	CreateMultipleCard(mc entities.McCard)error
	CreateQnACard(q entities.QaCard)error
}

type FcService struct {
	Repo FcRepo
}

func NewFcService(f Repo.Repo) FcService {
	return FcService{
		Repo: f,
	}
}

func (fc FcService) CreateMatchCard(mat entities.MatchCard) error {
	mat.Id = uuid.New().String()


	err := fc.Repo.CreateMatchCard(mat)
	if err != nil {
		return err
	}
	return nil
}


func (fc FcService) CreateTrueFalseCard(tf entities.TfCard) error {
	tf.Id = uuid.New().String()

	err := fc.Repo.CreateTrueFalseCard(tf)
	if err != nil {
		return err
	}
	return nil
}


func (fc FcService) CreateStudyCard(sc entities.InfoCard) error {
	sc.Id = uuid.New().String()

	err := fc.Repo.CreateStudyCard(sc)
	if err != nil {
		return err
	}
	return nil
}


func (fc FcService) CreateMultipleCard(mc entities.McCard) error {
	mc.Id = uuid.New().String()

	err := fc.Repo.CreateMultipleCard(mc)
	if err != nil {
		return err
	}
	return nil
}


func (fc FcService) CreateQnACard(q entities.QaCard) error {
	q.Id = uuid.New().String()

	err := fc.Repo.CreateQnACard(q)
	if err != nil {
		return err
	}
	return nil
}

