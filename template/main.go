package main

import (
	"fmt"
	"os"
	"text/template"
)

type Curse struct {
	Name  string
	Hours int
}

func main() {
	curse := Curse{"golang", 300}

	tmp := template.New("curse")
	tmp, _ = tmp.Parse("the course called '{{.Name}}' has {{.Hours}} hours.")

	if err := tmp.Execute(os.Stdout, curse); err != nil {
		fmt.Println(err)
	}

	// utilizando template.must

	t := template.Must(tmp.Parse("--O curso chamado {{.Name}} tem {{.Hours}} horas de duração"))

	if err := t.Execute(os.Stdout, curse); err != nil {
		fmt.Println(err)
	}
}
