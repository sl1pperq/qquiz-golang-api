package main

import (
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"io"
	"net/http"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type Response struct {
	Status  string `json:"status" xml:"status"`
	Message string `json:"message" xml:"message"`
}

func main() {
	e := echo.New()
	e.GET("/", MainPage)

	e.Logger.Fatal(e.Start(":1323"))
}

func MainPage(c echo.Context) error {
	u := &Response{
		Status:  "200",
		Message: "Сервер работает хорошо",
	}
	return c.JSON(http.StatusOK, u)
}
