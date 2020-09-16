package main

import (
	"net/http"
	"io"
	"log"
	"bytes"
	"fmt"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(fmt.Sprintf("===> New request from %s", r.RemoteAddr))
	pac := MustAsset("gfwlist.pac")

	w.Header().Set("Content-Type", "application/x-ns-proxy-autoconfig")

	cl := fmt.Sprintf("%v",len(pac))
	w.Header().Set("Content-Length", cl)

	pacReader := bytes.NewReader(pac)
	io.Copy(w, pacReader)
}

func main() {
	addr := ":1080"

	s := &http.Server{
                Addr:           addr,
                Handler:        http.HandlerFunc(myHandler),
                ReadTimeout:    10 * time.Second,
                WriteTimeout:   10 * time.Second,
                MaxHeaderBytes: 1 << 20,
        }

	log.Println(fmt.Sprintf("Server start at %s", addr))

	log.Fatal(s.ListenAndServe())
}
