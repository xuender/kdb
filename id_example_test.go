package kdb_test

import (
	"fmt"

	"github.com/xuender/kdb"
)

func ExampleDecode() {
	fmt.Println(kdb.Decode(11456000074))
	fmt.Println(kdb.Decode(11456000090))

	// Output:
	// 1693193954434 4 10
	// 1693193954434 5 10
}
