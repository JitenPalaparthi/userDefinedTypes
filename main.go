// keywords
// packages
// built in functions
// This is a inline comment
/*
This
	is
		multiline
			comment */
package main

import (
	"demo/rect"
	"fmt"
	"reflect"
)

func main() {

	r := rect.Rect{Length: 12.34, Width: 13.56}
	area := r.AreaOfRect()

	fmt.Println("Area of Rect:", area, "Stored:", r.Area)
	fmt.Printf("Address of r local to main %p\n", &r)

	fmt.Println("-----------------------------------")
	area1 := (&r).AreaOfRectPtr()
	fmt.Println("Area of Rect:", area1, "Stored:", r.Area)
	fmt.Printf("Address of r local to main %p\n", &r)
	fmt.Println("-----------------------------------")
	r1 := &rect.Rect{Length: 12.34, Width: 13.56}
	area2 := r1.AreaOfRectPtr()
	fmt.Println("Area of Rect:", area2, "Stored:", r1.Area)
	fmt.Printf("Address of r local to main %p\n", r)

	var um rect.UserMap

	um = make(rect.UserMap)
	fmt.Println(reflect.TypeOf(um))

	um["Name"] = "Jiten"
	um["Address"] = "Blr"
	um["Email"] = "JitenP@Outlook.Com"
	um["Mobile"] = "9618558500"

	um.Display()

	// The usual string convertion from int
	var i int = 97
	fmt.Println(string(i))

	// Using user defined type of Int and called the method of that
	var mi rect.MyInt = 97
	fmt.Println(mi.ToString())

}

// go run hello.go
// go build hello.go
// go build -o app hello.go
// go install hello.go

// builtin functions
// println , complex , make , len, cap , append , copy  ,delete, new

// keywords
// break , default , func ,   case ,  map ,
// else ,  package , switch , const , fallthrough , if ,
// range , continue , for , import , return , var

// keywords Yet to use
// interface ,defer , go ,struct ,chan ,goto ,select ,type

// to create a package
// 1- the name of the package must be the name of the directory
// 2- there are not special access specifiers there in Go
// 	anything inside a package starts with Uppercase is considered as public
//	anything inside a pacakge starts with lowerCase is considered as private

// ways to consume packages
// there are 3 types of pacakges
// 	1- go standard packages. eg fmt, time, math/rand
//	2- packages created inside the program/project. eg area
// 	3- thirdparty packages those usually downloaded from the web. eg gorm, gin

// using go module
// go to the project root directory
// go mod init demo // name of the module can be any valid identifier
// call the pacakge with modulename eg. demo/area
