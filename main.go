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
			"template file type",
		)
		name *string = flag.String(
			"n",
			"hood",
			"file name",
		)
	)

	flag.Parse()

	if len(*t) == 0 {
		fmt.Println("-t is required.")
		os.Exit(1)
	}

	ok := hood.GenerateTemplate(*t, *name)
	if !ok {
		fmt.Println("something is wrong :(")
	}

}