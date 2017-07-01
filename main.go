package main

import (
	"github.com/go-martini/martini"
	"net/http"
	"./models"
	"encoding/json"
	"log"
	"strings"
	"strconv"
	"os"
)

func main() {
	m := martini.Classic()
	m.Group("/api", func(r martini.Router) {
		r.Post("/join",join)
		r.Get("/get",get)
	})
	m.Run()
}

func join(res http.ResponseWriter, req *http.Request, log *log.Logger) {
	req.ParseForm()
	s := models.Student{
		Name:      req.Form.Get("name"),
		School:    req.Form.Get("school"),
		Grade:     req.Form.Get("grade"),
		Class:     req.Form.Get("class"),
		Telephone: req.Form.Get("telephone"),
		Address:   req.Form.Get("address"),
		Text:      req.Form.Get("text"),
	}
	if req.Form.Get("sex") == "0" {
		s.Sex = "男"
	} else {
		s.Sex = "女"
	}
	if req.Form.Get("type") == "1" {
		s.Type = "文科"
	} else {
		s.Type = "理科"
	}
	m := make(map[string]string)
	err := json.Unmarshal([]byte(req.Form.Get("score")), &m)
	if err != nil {
		res.WriteHeader(500)
		log.Fatal(err)
	}
	for key,value:=range m{
		if strings.Index(value,"/")!= -1{
			score,_ := strconv.Atoi(value[:strings.Index(value,"/")])
			fullScore ,_:=strconv.Atoi(value[strings.Index(value,"/")+1:])
			s.AddScore(key,score,fullScore)
		}
	}
	err = s.Insert()
	if err != nil {
		res.WriteHeader(500)
		log.Fatal(err)
	}
	res.Header().Set("Access-Control-Allow-Origin","*")
	res.WriteHeader(200)
}

func get(res http.ResponseWriter,log *log.Logger){
	data,err := models.GetAllStudent()
	if err!=nil{
		res.WriteHeader(500)
		log.Fatal(err)
	}
	j ,_:= json.Marshal(data)
	res.Header().Set("Access-Control-Allow-Origin","*")
	res.WriteHeader(200)
	res.Write(j)
}