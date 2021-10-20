package main

import (
	"bootcamp/example"
	myfmt "fmt"

	"github.com/google/uuid"
)

func main() {

	// s := store{
	// 	name: "a",
	// 	lat:  2.88888,
	// 	long: 1.22222,
	// }

	snew := store{
		name: "add",
		lat:  2.88978,
		long: 1.32222,
	}

	var nn name = "SG"

	myfmt.Printf(snew.address())
	myfmt.Print(nn.hello())
	myfmt.Printf("\n%d\n", example.Sub(4, 9))

	id, _ := uuid.NewRandom()
	myfmt.Println(id.String())

}
