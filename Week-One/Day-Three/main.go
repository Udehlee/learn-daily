package main

import (
	"html/template"
	"math"
	"os"
)

var tmpl *template.Template

func init() {

	tmpl = template.Must(template.New("").Funcs(fm).ParseFiles("tmpl.gohtml"))

}

// fmn and fsqr are custom template function that maps out to the mean and square functions respectively
var fm = template.FuncMap{
	"fmn":  mean,
	"fsqr": square,
}

// mean returns the sum of nums divide by len of nums
func mean(nums []float64) float64 {

	sum := 0.0

	for _, num := range nums {
		sum += num
	}

	return sum / float64(len(nums))

}

// square retuns square of nums
func square(nums float64) float64 {

	return math.Pow(float64(nums), 2)

}

func main() {

	nums := []float64{1, 2, 3, 4, 5, 6}

	//ExecuteTemplate executes the template and print to the standard output
	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.gohtml", nums)
	if err != nil {
		panic(err)
	}

}
