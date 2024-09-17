package main

import (
	"os"
	"text/template"
)

func main() {
	type Curse struct {
		Name  string
		Hours int
	}

	type Curses []Curse

	temp := template.Must(template.New("template.html").ParseFiles([]string{"template.html"}...))

	myCurses := Curses{
		{"Aprenda Go", 1000},
		{"Aprenda Python", 1500},
		{"Aprenda Java", 2000},
	}

	if err := temp.Execute(os.Stdout, myCurses); err != nil {
		panic(err)
	}
}
