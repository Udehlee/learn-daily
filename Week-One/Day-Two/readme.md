## overview

Learned about functions in Go templates.

Functions in Go templates help you perform different operations on the data you're working with, making your templates more flexible and powerful

- Here you can create your own custom function 
- use the built-in function provided in the go template packages

## keep in mind

// fm is a FuncMap containing template functions.
// "psw" is a custom template function key that maps to the pickSecondWord function
var fm = template.FuncMap{
	"psw": pickSecondWord,
}

func init() {
	tmpl = template.Must(template.New("").Funcs(fm).ParseFiles("tmpl.gohtml"))

}

// ExecuteTemplate executes the template and print the result to the standard output.
	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", data)

	if err != nil {
		panic(err)
	}
