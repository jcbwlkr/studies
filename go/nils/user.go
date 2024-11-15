package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// User is an example struct
type User struct {
	Name string
}

// Clear sets the user to nil
func (u *User) Clear() {
	v := reflect.ValueOf(u)
	// To set it back to a zero value (but not nil)
	//v.Elem().Set(reflect.Zero(v.Elem().Type()))

	//p := unsafe.Pointer(v.Elem().UnsafeAddr())

	v.Elem().Addr().SetPointer(unsafe.Pointer(nil))
	//reflect.ValueOf(nil)

	describe("clear", u)
}

func runUser() {
	var user *User

	describe("init", user)

	user = &User{Name: "Bob"}

	describe("set", user)

	user.Clear()

	describe("end", user)
}

func describe(what string, u *User) {
	fmt.Printf("%s\t%v, %p\n", what, u, u)
}
