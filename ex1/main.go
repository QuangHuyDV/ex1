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

type User struct {
	name    string
	age     int64
	gender  bool
	address string
}

func (u *User) GetName() string {
	return u.name
}
func (u *User) GetAge() int64 {
	return u.age
}
func (u *User) GetGender() bool {
	return u.gender
}
func (u *User) GetAddress() string {
	return u.address
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
	fmt.Println("=========================")
	//ex6
	num := []int{11, 34, 77, 99, 109, 66, 20, 88, 34}
	fmt.Println("Thứ tự ban đầu: ", num)
	//bubbleSoft(a)
	//fmt.Println("Sau khi sắp xếp: ", a)
	var x []int = num[1:7]
	fmt.Println("Thứ tự từ 1 -> 7: ", x)
	//var y []int = x[1:15]
	//fmt.Println(y)
	fmt.Println("Copy slice trên với index từ 1 đến 15 -> bị lỗi ")
	fmt.Println("=========================")
	//ex7
	user1 := User{name: "Trung", age: 12, gender: true, address: "Hue"}
	user2 := User{name: "Hoa", age: 15, gender: true, address: "HN"}
	user3 := User{name: "Ha", age: 13, gender: false, address: "HN"}
	user4 := User{name: "Trang", age: 12, gender: true, address: "HN"}
	user5 := User{name: "Minh", age: 13, gender: false, address: "HN"}
	user6 := User{name: "Lan", age: 11, gender: true, address: "HN"}

	a := map[User]string{user1: user1.name, user2: user2.name, user3: user3.name, user4: user4.name, user5: user5.name, user6: user6.name}

	for str, i := range a {
		fmt.Println(str, i)
	}
}
