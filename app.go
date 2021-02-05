package main

import (
	"mypackages/myTime"
	"os"
	"strings"
	"text/template"
	"time"
)

var tpl *template.Template

func myFn (name string) string {
	return name[:3]
}

var funcMap = template.FuncMap{
	"myFn": myFn,
	"upper": strings.ToUpper,
	"date": myTime.ToDateString,
}


func init() {
	tpl = template.Must(template.New("").Funcs(funcMap).ParseFiles("templates/Page.gohtml"))
}


func main() {
	
	peoples := []string{"Rasel", "Rahim", "Karim", "Alex", "name"}
	data := map[string]interface{}{
		"persons": peoples,
		"timestamp": time.Now(),
	}
	
	_ = tpl.ExecuteTemplate(os.Stdout, "Page.gohtml",	data)
	
}


