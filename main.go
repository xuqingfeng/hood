package main

import (
	"flag"
	"fmt"
	"github.com/xuqingfeng/hood/hood"
	"os"
)

func main() {

	var (
		t *string = flag.String(
			"t",
			"",
			"file template type",
		)
		name *string = flag.String(
			"n",
			"hood",
			"file name",
		)
	)

	flag.Parse()

	if len(*t) == 0 {
		fmt.Println("-t Is Required.")
		os.Exit(1)
	}

	ok := hood.GenerateTemplate(*t, *name)
	if !ok {
		fmt.Println("Something Is Wrong :(")
	}

}
