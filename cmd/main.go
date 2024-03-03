package main

import (
	"github.com/GeorgeGuo2018/rangeList/pkg/types"

	"fmt"
	"net/http"
	"net/http/pprof"
)

func main() {
	// start pprof server
	http.Handle("/debug", pprof.Handler("rangelist"))

	//the following is just a example
	rgA, _ := types.NewRange(0, 5)
	rgB, _ := types.NewRange(-3, 7)
	rgC, _ := types.NewRange(5, 9)

	rgList := types.NewRangeList()
	rgList.Add(rgA)
	rgList.Add(rgB)
	rgList.Add(rgC)
	rgList.Remove(rgB)
	fmt.Printf(rgList.ToString())
}
