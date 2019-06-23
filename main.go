package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

var args []string

func handler(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	cmd := exec.CommandContext(r.Context(), args[0], args[1:]...)
	cmd.Stdin = r.Body
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		io.WriteString(w, err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = cmd.Wait()
	if err != nil {
		log.Printf(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
	io.Copy(w, buf)
}

func main() {
	args = os.Args[1:]

	log.Printf("start server")
	log.Printf("args: %+v", args)

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
