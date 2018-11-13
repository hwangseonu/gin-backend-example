package docs

import (
	"encoding/json"
	"github.com/swaggo/swag"
)

type License struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Contact struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Email string `json:"email"`
}

type APIInfo struct {
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	TermsOfService string  `json:"terms_of_service"`
	Contact        Contact `json:"contact"`
	License        License `json:"license"`
	Version        string  `json:"version"`
}

type APIDoc struct {
	Swagger  string  `json:"swagger"`
	Info     APIInfo `json:"info"`
	Host     string  `json:"host"`
	BasePath string  `json:"basePath"`
}

func (doc APIDoc) ReadDoc() string {
	b, e := json.MarshalIndent(doc, "", "  ")
	if e != nil {
		panic(e)
	}
	return string(b)
}

func init() {
	swag.Register(swag.Name, &APIDoc{
		Swagger: "2.0",
		Info: APIInfo{
			Title: "Gin Backend",
			Description: "This is backend server with gin",
			TermsOfService: "",
			Contact: Contact{
				Name: "hwangseonu",
				Url: "http://blog.mocha.ga",
				Email: "hwangseonu12@naver.com",
			},
			License: License{
				Name: "MIT",
				Url: "https://github.com/hwangseonu/gin-backend/blob/master/LICENSE",
			},
			Version: "1.0.0-SNAPSHOT",
		},
		Host: "https://gin.mocha.ga",
		BasePath: "",
	})
}
