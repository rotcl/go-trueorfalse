package main

import (
	"fmt"
	"net/http"
)

func False(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, false)
}
