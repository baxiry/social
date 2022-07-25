package main

import "fmt"

func CheckErr(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		return
	}
}
