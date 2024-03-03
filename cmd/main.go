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
	//todo
	rgA := types.NewRange(0, 5)
	rgB := types.NewRange(-3, 7)
	rgC := types.NewRange(5, 9)

	rgList := types.NewRangeList()
	rgList.Add(rgA)
	rgList.Add(rgB)
	rgList.Add(rgC)
	rgList.Remove(rgB)
	fmt.Printf(rgList.ToString())
}
