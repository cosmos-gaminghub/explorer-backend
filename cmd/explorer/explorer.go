package main

import (
	"fmt"

	"github.com/cosmos-gaminghub/explorer-backend/client"
)

func main() {
	result, err := client.GetContract("juno19f5yfd3trdt2n5pugln3aqzn83v5qkrjnqgf08f66sx6avgll2tssk4ugc")
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
