//main.go (01)

package main

import (
	"html/template"
	"os"
	"strings"
)

var tmpl *template.Template

// Person represent a person with a name and string
type Person struct {
	Name    string
	Sayings string
}

// fm is a FuncMap containing template functions.
// "psw" is a custom template function key that maps to the pickSecondWord function
var fm = template.FuncMap{
	"psw": pickSecondWord,
}

func init() {
	tmpl = template.Must(template.New("").Funcs(fm).ParseFiles("tmpl.gohtml"))

}

// pickSecondWord returns the second word of a string
func pickSecondWord(word string) string {

	words := strings.Fields(word)

	if len(words) >= 2 {
		return words[1]
	}
	return ""
}

// main.go (02)

func main() {

	Nma := Person{
		"Nma",
		"No more procrastination",
	}

	Ngbede := Person{
		"Ngbede",
		"One step at a time",
	}

	Sayers := []Person{Nma, Ngbede}

	data := struct {
		wisemen []Person
	}{
		Sayers,
	}

	// ExecuteTemplate executes the template and print the result to the standard output.
	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", data)

	if err != nil {
		panic(err)
	}

}
