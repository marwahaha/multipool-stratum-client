package main

import (
	"fmt"

	mf "github.com/lerenn/mining-feeder/pkg"
)

func main() {
	port := 3336
	addr := "ltc.pool.minergate.com"
	username := "louis.fradin@gmail.com"
	password := "x"

	pool, err := mf.NewStratumPool(addr, port)
	if err != nil {
		panic(err)
	}
	pool.AddCredentials(username, password)

	res, err := pool.GetWork()
	if err != nil {
		panic(err)
	}

	fmt.Println("Result: " + res)
}
