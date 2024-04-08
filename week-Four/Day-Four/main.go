package main

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	p := "lee"

	bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	if err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(bs)
}
