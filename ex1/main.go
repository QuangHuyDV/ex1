package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
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

func ex3(serves []*Serve) {
	key := "admin"
	for i := range serves {
		if serves[i].Name == key {
			fmt.Println(serves[i])
		}
	}

}

func ex4(serves []*Serve) []*Serve {

	add := Serve{Name: "fileCustome", Class: "org.cofax.cds.FileServlet.Custome"}
	new := append(serves, &add)
	return new
}

func bubbleSoft(num sort.IntSlice) {
	for j := 0; j < len(num)-1; j++ {
		for i := len(num) - 2; i >= j; i-- {
			if num[i] > num[i+1] {
				tg := num[i]
				num[i] = num[i+1]
				num[i+1] = tg
			}
		}
	}
}

func ex6() {
	num := []int{11, 34, 77, 99, 109, 66, 20, 88, 34}
	fmt.Println("Thứ tự ban đầu: ", num)
	var x []int = num[1:7]
	fmt.Println("Thứ tự từ 1 -> 7: ", x)
	bubbleSoft(num)
	fmt.Println("Sau khi sắp xếp: ", num)
}

func ex7() {
	user1 := User{name: "Trung", age: 12, gender: true, address: "Hue"}
	user2 := User{name: "Hoa", age: 15, gender: true, address: "HN"}
	user3 := User{name: "Ha", age: 13, gender: false, address: "HN"}
	user4 := User{name: "Trang", age: 12, gender: true, address: "HN"}
	user5 := User{name: "Minh", age: 13, gender: false, address: "HN"}
	user6 := User{name: "Lan", age: 11, gender: true, address: "HN"}

	users := map[User]string{user1: user1.name, user2: user2.name, user3: user3.name, user4: user4.name, user5: user5.name, user6: user6.name}
	for str, i := range users {
		fmt.Println(str, i)
	}
}

func main() {

	jsonFile, err := os.Open("serve.json")

	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	defer jsonFile.Close()
	serves := []*Serve{}
	json.Unmarshal(byteValue, &serves)
	log.Println(serves)
	for i := range serves {
		log.Printf("Name: %s, Class: %s\n", serves[i].Name, serves[i].Class)
	}
	fmt.Println("=================")

	ex3(serves)
	fmt.Println("=================")
	for i := range ex4(serves) {
		fmt.Printf("Name: %s, Class: %s\n", ex4(serves)[i].Name, ex4(serves)[i].Class)
	}
	fmt.Println("=================")

	//ex5
	for i := range ex4(serves) {
		fmt.Println("Class:", ex4(serves)[i].Class)
	}
	fmt.Println("=================")

	ex6()
	fmt.Println("=================")

	ex7()

}
