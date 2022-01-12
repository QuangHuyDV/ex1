package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Serve struct {
	Name  string
	Class string
}

func main() {
	//ex1,ex2
	jsonFile, err := os.Open("serve.json")

	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	defer jsonFile.Close()

	b := []Serve{}
	json.Unmarshal(byteValue, &b)
	log.Println(b)
	fmt.Println("=========================")
	//ex3
	for i := range b {
		if b[i].Name == "cofaxAdmin" {
			fmt.Printf("Name: %s, Class: %s\n", b[i].Name, b[i].Class)
		}
	}
	fmt.Println("=========================")
	//ex4
	c := Serve{Name: "fileCustome", Class: "org.cofax.cds.FileServlet.Custome"}
	d := append(b, c)
	fmt.Println(d)
	fmt.Println("=========================")
	//ex5
	for i := range d {
		fmt.Printf("Name: %s, Class: %s\n", d[i].Name, d[i].Class)
	}
}
