/*
Exercise: HTTP Handlers

Implement the following types and define ServeHTTP methods on them. 
Register them to handle specific paths in your web server.

type String string

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}
For example, you should be able to register handlers using:

http.Handle("/string", String("I'm a frayed knot."))
http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
*/

package main

import (
	"fmt"
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s String) ServeHTTP(w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, s)
}

func (s *Struct) ServeHTTP(w http.ResponseWriter,
	r *http.Request) {
	fmt.Fprint(w, s.Greeting, s.Punct, s.Who)
}
func main() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	http.ListenAndServe("localhost:4000", nil)
}
