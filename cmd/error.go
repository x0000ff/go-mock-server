package main

import (
	"fmt"
	"log"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func checkFatalErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
