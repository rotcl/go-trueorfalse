package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func Random(w http.ResponseWriter, r *http.Request) {
	bul := rand.Intn(10)
	if bul >= 5 {
		fmt.Fprint(w, true)
	} else {
		fmt.Fprint(w, false)
	}
}