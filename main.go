package main

import (
	// "fmt"
	"log"
	"flag"
	"net/http"
	"os/exec"
	"io"
	"bytes"
	"strings"
	
	"github.com/txgruppi/parseargs-go"
)

func executeCommand(c string, args []string, stdin io.Reader) (*bytes.Buffer, error){
	cmd := exec.Command(c, args...)
	if stdin != nil {
		cmd.Stdin = stdin
	} 

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	return &stdout, nil
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	// get command
	command := req.URL.Path
	command = strings.TrimLeft(command, "/")

	// get command args
	args := req.Header.Get("X-Args")
	parsed, err := parseargs.Parse(args)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	log.Printf("get command %s with args '%s' ", command, args)

	// get stdin
	stdout, err := executeCommand(command, parsed, req.Body)

	if err != nil {
		_, err = w.Write([]byte(err.Error()))
	} else {
		_, err = w.Write(stdout.Bytes())
	}


	if err != nil {
		log.Printf("write execute command error %s", err)
	}

}


func main() {
	var port string
	flag.StringVar(&port, "port", ":7890", "listen port")
	flag.Parse()

	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	
	http.HandleFunc("/", indexHandler)

	log.Printf("start listen on %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
