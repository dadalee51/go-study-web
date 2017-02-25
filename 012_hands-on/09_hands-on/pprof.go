package main

import (
	_ "net/http/pprof"
	"net/http"
)

func main (){
	http.ListenAndServe("localhost:6060", nil)
}