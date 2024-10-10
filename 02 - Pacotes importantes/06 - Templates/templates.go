package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{"Go", 50}
	t := template.Must(template.New("CursoTemplate").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}"))
	// template := template.New("CursoTemplate")
	// template, _ = template.Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}")
	err := t.Execute(os.Stdout, curso)
	if err != nil {
		panic(err)
	}

}
