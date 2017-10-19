package main

import (
	"fmt"
	"github.com/kris-nova/go-azure/arm"
	"github.com/kris-nova/go-azure/arm/aci"
	"log"
)

func main() {

	clientID := ""
	clientSecret := ""
	subscriptionId := ""
	tenantID := ""

	auth := arm.NewAuthentication(clientID, clientSecret, subscriptionId, tenantID)

	client := aci.NewClient(auth)

	input := &aci.CreateContainerInput{
	// TODO define container information here
	}

	output, err := client.Create(input)

	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf("%+v\n", output)
}
