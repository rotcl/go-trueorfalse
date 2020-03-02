package main

import (
	"fmt"
	"net/http"
)

func True(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, true)
}
