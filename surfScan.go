package main

import (
	"fmt"
	"github.com/Smolkar/go-masscan"

)

func main(){
	var lPort string
	var lRange string
	var lRate string
	var lAdapter string
	fmt.Println("SURFSCAN")
	fmt.Println("Powered by Masscan")
	fmt.Println("-------------------\n")

	fmt.Print("PORT: ")
	fmt.Scanln(&lPort)
	fmt.Print("RANGE: ")
	fmt.Scanln(&lRange)
	fmt.Print("RATE OF SCANNING: ")
	fmt.Scanln(&lRate)
	fmt.Print("ADAPTER PORT(if not specified leave empty): ")
	fmt.Scanln(&lAdapter)

	scanner := masscan.New()
	scanner.SetPorts(lPort)
	scanner.SetRanges(lRange)
	scanner.SetRate(lRate)
	scanner.SetAdapter_port(lAdapter)
	err := scanner.Run()
	if err != nil{
		fmt.Println("sccanner failed ", err)
		return
	}
	results, err := scanner.Parse()
	if err != nil{
		fmt.Println("Failed scanner")
		return
	}
	for _ , result := range results{
		fmt.Println(result)
	}

}
