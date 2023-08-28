package id_test

import (
	"fmt"

	"github.com/xuender/kdb/id"
)

func ExampleNew() {
	fmt.Println(id.New() > 0)

	// Output:
	// true
}

func ExampleID_Decode() {
	fmt.Println(id.ID(11456000074).Decode())
	fmt.Println(id.ID(11456000090).Decode())

	// Output:
	// 1693193954434 4 10
	// 1693193954434 5 10
}

func ExampleID_String() {
	fmt.Println(id.ID(11456000074))

	// Output:
	// 11456000074
}

func ExampleParse() {
	fmt.Println(id.Parse("11456000074"))
	fmt.Println(id.Parse("a"))

	// Output:
	// 11456000074 <nil>
	// 0 strconv.ParseUint: parsing "a": invalid syntax
}
