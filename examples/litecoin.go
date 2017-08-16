package main

import (
	"fmt"

	msc "github.com/lerenn/multipool-stratum-client"
)

func main() {
	port := 3336
	addr := "ltc.pool.minergate.com"
	username := "louis.fradin@gmail.com"
	password := "x"

	pool := msc.BitcoinStratumPool{}
	if err := pool.Connect(addr, port); err != nil {
		panic(err)
	}

	res, err := pool.GetWork(username, password)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
