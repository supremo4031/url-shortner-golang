package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"url-shortner/handler"
)


func main() {

	yamlFile := flag.String("yaml", "url-mapping.yaml", "a yaml file with short and long url mapping")
	flag.Parse()

	yaml, err := os.ReadFile(*yamlFile)
	if err != nil {
		panic(err)
	}

	mux := defaultMux()

	pathsToUrls := map[string]string {
		"/ggl": "https://www.google.com",
		"/lkdn": "https://www.linkedin.com",
		"/lpfl": "https://www.linkedin.com/in/arpan-layek-155003192/",
	}

	mapHanlder := handler.MapHandler(pathsToUrls, mux)

	yamlHandler, err := handler.YAMLHandler([]byte(yaml), mapHanlder)

	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}