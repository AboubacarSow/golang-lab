package main

import (
	"net/http"
)



func healthHandler(w http.ResponseWriter, r *http.Request){
	jsonHelper(w,200,struct{}{});
}

func errorHandler(w http.ResponseWriter, r * http.Request){
	errorHelper(w, 400,"Something went wrong")
}




