package main

import (
	"restaurant/diner"
)

func main() {

	data := diner.ReadData()
	diner.FindDuplicate(data)
	diner.ReadTop3Menu(data)
}
