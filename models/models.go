package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {
	engine, _ = xorm.NewEngine("mysql", "test1:8mS2cXH8LM@tcp(47.94.91.118:3306)/test1?charset=utf8")
	engine.Sync(new(Student))
}

type Student struct {
	Id        int `xorm:"autoincr"`
	Name      string `xorm:"char(5)"`
	Sex       string `xorm:"char(1)"`
	Telephone string `xorm:"char(13)"`
	Address   string `xorm:"varchar(50)"`
	School    string `xorm:"varchar(20)"`
	Grade     string `xorm:"char(5)"`
	Class     string `xorm:"char(5)"`
	Type      string `xorm:"char(2)"`
	Text      string `xorm:"varchar(255)"`
	Date      string `xorm:"DATE created"`
	Score     []Score `xorm:"json"`
}

type Score struct {
	Subject   string
	Score     int
	FullScore int
}

func (s *Student) Insert() (err error) {
	_, err = engine.Insert(s)
	return
}

func (s *Student) Delete() (err error) {
	_, err = engine.Delete(s)
	return
}

func (s *Student) AddScore(subject string, score, fullScore int) {
	sc := Score{
		Subject:   subject,
		Score:     score,
		FullScore: fullScore,
	}
	s.Score = append(s.Score, sc)
}

func GetAllStudent() (data []Student, err error) {
	err = engine.Find(&data)
	return
}
