package rect

import "fmt"

type Rect struct {
	Length, Width, Area float64
}

func AreaOfRect(r Rect) float64 {
	return r.Length * r.Width
}

// Go has a notation call receivers
// methods have receivers but functions do not
// a pointer receiver can maintain the state of the object
// if not a pointer receive , generally receivers are call by value and so it creates a new receiver object
// receivers are not only for structs

func (r Rect) AreaOfRect() float64 {
	r.Area = r.Length * r.Width
	fmt.Println("Area Stored for R", r.Area)
	fmt.Printf("Address of r %p\n", &r)
	return r.Area
}

func (r *Rect) AreaOfRectPtr() float64 {
	r.Area = r.Length * r.Width
	fmt.Println("Area Stored for R", r.Area)
	fmt.Printf("Address of r %p\n", r)
	return r.Area
}

// create a type from a defined type

type UserMap map[string]string

func (um UserMap) Display() {
	for key, val := range um {
		fmt.Println("Key:", key, " Value:", val)
	}
}

// create a type called int and convert it to string

type MyInt int

func (mi MyInt) ToString() string {
	return fmt.Sprint(mi)
}
