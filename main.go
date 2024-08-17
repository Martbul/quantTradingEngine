package main
//! gledam strima i pravq proekta za da se nauà
//! kogato ima neshto koeto ne razbirash vednaga googlevahs/gpt za da go nauchis i stanesh high value developer
//! ne iskap gotov produkt a znaniq

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

func main() {
	var asks map[float64]float64

	f,err := os.Open("./data/asks.gob")
	if err != nil{
		log.Fatal(err)
	}
	
	if err := gob.NewDecoder(f).Decode(&asks); err != nil{
		log.Fatal(err)
	}

	for price, size := range asks {
		fmt.Printf("%.2f - %.2f\n", price, size)
	}
}