package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/xuqingfeng/hood/hood"
)

func main() {

	var (
		t *string = flag.String(
			"t",
			"",
			"template type(ansible, c, go, html, ignore, java, js, php, python)",
		)
		name *string = flag.String(
			"n",
			"hood",
			"name",
		)
	)

	flag.Parse()

	if len(*t) == 0 {
		fmt.Println("E! type must be speficied")
		os.Exit(1)
	}

	ok := hood.GenerateTemplate(*t, *name)
	if !ok {
		fmt.Println("E! something is wrong")
	}

}
