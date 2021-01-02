package main

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
)

func main() {
	// https://forum.golangbridge.org/t/get-golang-variable-value-in-html-script-tag/5349
	tmpl, err := template.ParseFiles("./third_party/mail/html/invitation.html")
	if err != nil {
		fmt.Println("Error")
		return
	}

	data := struct {
		RepName string
		Inviter string
	}{
		RepName: "Chee Wai Xiong",
		Inviter: "Calum",
	}

	var htmlBuffer bytes.Buffer
	err = tmpl.Execute(&htmlBuffer, data)
	if err != nil {
		fmt.Println("Error")
		return
	}

	// open output file
	fo, err := os.Create("cmd/program/mail/generated.html")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()
	if _, err := fo.Write(htmlBuffer.Bytes()); err != nil {
		panic(err)
	}

	fmt.Println("Done")
}
