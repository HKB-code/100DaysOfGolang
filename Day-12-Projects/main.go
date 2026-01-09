package main

import (
	"fmt"

	filemanager "example.com/main/fileManager"
	"example.com/main/prices"
)



func main() {
	
	taxRates := []float64{0, 0.07, 0.1, 0.15}


	for _, taxRate := range taxRates {
		fm:= filemanager.NewFm("prices/prices.txt",fmt.Sprintf("result_%.0f.json",taxRate*100))

		// cmd:= cmdmanager.New()
		priceJob:=prices.New(*fm,taxRate)
		err:=priceJob.Process()
		if  err!= nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
		}

	}
	
}