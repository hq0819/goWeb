package main

import (
	"fmt"
	"reflect"
)

type People interface {
	Chifan()
}

type User struct {
	Name  string
	Age   int
	Email string
}

func (u User) Chifan() {
	fmt.Println("吃饭")
}

func (u User) GetUsername(e *User) string {
	return e.Name
}

func main() {

	user := User{Name: "heqin", Age: 19, Email: "1234@QQ.COM"}
	of := reflect.TypeOf(user)

	valueOf := reflect.ValueOf(user)
	u := valueOf.Interface().(User)
	fmt.Println(u)

	name, b := of.MethodByName("Chifan")
	byName := valueOf.MethodByName("Chifan")
	fmt.Println(byName)

	fmt.Println(name, b)

	fmt.Println(of, valueOf)

}
