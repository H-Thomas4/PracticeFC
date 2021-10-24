package entities

import "github.com/google/uuid"

type FlashCard struct {
	Matching []MatchCard
	TrueOrFalse []TfCard
	Info []InfoCard
	Multiple []McCard
	QA []QaCard
}

type MatchCard struct {
	Id string
	Category string
	Type string
	Question map[string]interface{}
	Options map[string]interface{}
	Answer map[string]interface{}
}

type TfCard struct {
	Id string
	Category string
	Type string
	Question string
	Answer bool
}

type InfoCard struct {
	Id string
	Category string
	Type string
	Details string
}

type McCard struct {
	Id string
	Category string
	Type string
	Question string
	Options map[string]interface{}
	Answer string
}

type QaCard struct {
	Id string
	Category string
	Question string
	Answer string
}

func (match *MatchCard) SetId() {
	match.Id = uuid.New().String()
}

func (tf *TfCard) SetId() {
	tf.Id = uuid.New().String()
}

func (ic *InfoCard) SetId() {
	ic.Id = uuid.New().String()
}

func (mc *McCard) SetId() {
	mc.Id = uuid.New().String()
}

func (q *QaCard) SetId() {
	q.Id = uuid.New().String()
}