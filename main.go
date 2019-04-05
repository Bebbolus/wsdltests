package main

import (
        "log"
        "./gen"
        "github.com/hooklift/gowsdl/soap"
)

func main() {
        client := soap.NewClient("http://www.dneonline.com/calculator.asmx")
        service := gen.NewCalculatorSoap(client)
        reply, err := service.Add(&gen.Add{IntA:4,IntB:5})
        if err != nil {
                log.Fatalf("could't get caluculator: %v", err)
        }
        log.Println(reply)
}

