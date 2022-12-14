package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/melinaco4/company-manager-mongo/controllers"
	"gopkg.in/mgo.v2"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewCompanyController(getSession())
	r.GET("/company/:id", uc.GetCompany)
	r.POST("/company", uc.CreateCompany)
	r.DELETE("/company/:id", uc.DeleteCompany)
	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session {

	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}
