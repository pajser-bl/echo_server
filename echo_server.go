package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const PORT = "10000"
const address = "localhost" + ":" + PORT

func main() {
	fmt.Printf("Starting server at %s\n", PORT)
	http.HandleFunc("/", echo)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatal(err)
	}
}

func echo(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	_, _ = fmt.Fprintf(os.Stdout, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		_, _ = fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		_, _ = fmt.Fprintf(os.Stdout, "Header[%q] = %q\n", k, v)
	}
	_, _ = fmt.Fprintf(w, "Host = %q\n", r.Host)
	_, _ = fmt.Fprintf(os.Stdout, "Host = %q\n", r.Host)
	_, _ = fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	_, _ = fmt.Fprintf(os.Stdout, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		_, _ = fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
		_, _ = fmt.Fprintf(os.Stdout, "Form[%q] = %q\n", k, v)
	}
	body, _ := ioutil.ReadAll(r.Body)
	_, _ = fmt.Fprintf(w, "Body = %s\n", string(body))
	_, _ = fmt.Fprintf(os.Stdout, "Body = %s\n", string(body))
}
