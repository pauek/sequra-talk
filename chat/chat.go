package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"text/template"

	"golang.org/x/net/websocket"
)

const listenAddr = "demo.pauek.info:80"

func cp(w io.Writer, r io.Reader, errc chan<- error) {
	_, err := io.Copy(w, r)
	errc <- err
}

type socket struct {
	io.ReadWriter
	done chan bool
}

func (s socket) Close() error {
	s.done <- true
	return nil
}

func chat(a, b io.ReadWriteCloser) {
	fmt.Fprintln(a, "Encontrado! Dile algo...")
	fmt.Fprintln(b, "Encontrado! Dile algo...")
	errc := make(chan error, 1)
	go cp(a, b, errc)
	go cp(b, a, errc)
	if err := <-errc; err != nil {
		log.Println(err)
	}
	a.Close()
	b.Close()
}

var partner = make(chan io.ReadWriteCloser)

func match(c io.ReadWriteCloser) {
	fmt.Fprint(c, "Esperando un compañero...")
	select {
	case partner <- c:
		// Si el otro `match` nos coge el canal
		// ya se encarga él del chat
	case p := <-partner:
		chat(p, c)
	}
}

func socketHandler(ws *websocket.Conn) {
	s := socket{ws, make(chan bool)} // HL
	go match(s)
	<-s.done
}

var T = template.Must(template.New("index").ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	T.ExecuteTemplate(w, "index.html", listenAddr)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/socket", websocket.Handler(socketHandler))
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
