package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Baumanar/test_it/service"
)

func main() {
	fileName := flag.String("f", "", "file name in which mowers program are written")
	flag.Parse()

	mowers := service.NewMowersService()
	err := mowers.Init(*fileName)
	if err != nil {
		fmt.Printf("Failed initialization: %s\n", err)
		os.Exit(1)
	}
	err = mowers.Run()
	if err != nil {
		fmt.Printf("Failed while mowers were running: %s\n", err)
		os.Exit(1)
	}
	fmt.Print(mowers.ResToString())
}
