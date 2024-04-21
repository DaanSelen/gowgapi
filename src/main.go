package main

import (
	"gowgapi/ifmana"
	"log"
)

func main() {
	err := ifmana.StartService("nigger")
	if err != nil {
		log.Println(err)
	}
}
