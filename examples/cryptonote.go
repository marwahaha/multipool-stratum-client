package main

import (
	"fmt"

	mf "github.com/lerenn/mining-feeder"
)

func main() {
	// port := 3336
	// addr := "ltc.pool.minergate.com"
	port := 45560
	addr := "xmr.pool.minergate.com"
	username := "louis.fradin@gmail.com"
	password := "x"

	pool := mf.CryptonoteStratumPool{}
	if err := pool.Connect(addr, port); err != nil {
		panic(err)
	}

	res, err := pool.GetWork(username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
