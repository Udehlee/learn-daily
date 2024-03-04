package main

import (
	"os"
	"text/template"
)

// Animal represents an animal with its sound
type Animal struct {
	Name  string
	Sound string
}

func main() {

	dog := Animal{
		"dogs",
		"barking",
	}

	//template for displaying animal information

	tmpl, err := template.New("AnimalTemplate").Parse("{{.Name}} communicate through {{.Sound}} , body language and  facial expressions ")

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, dog)
	if err != nil {
		panic(err)
	}

}
