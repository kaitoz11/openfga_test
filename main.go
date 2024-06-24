package main

import (
	"context"
	"fmt"

	fgaSdk "github.com/openfga/go-sdk/client"
)

func main() {

	fgaClient, err := fgaSdk.NewSdkClient(&fgaSdk.ClientConfiguration{
		ApiUrl:               "http://localhost:8080",      // required, e.g. https://api.fga.example
		StoreId:              "01J15APP88MY5NTG0VT61DZT86", // optional, not needed for \`CreateStore\` and \`ListStores\`, required before calling for all other methods
		AuthorizationModelId: "01J15TYXFEXM7M0GXYTDCJNAMP", // Optional, can be overridden per request
	})

	if err != nil {
		// .. Handle error
		panic(err)
	}

	bodyCanCheckin := fgaSdk.ClientCheckRequest{
		User:     "user:chien",
		Relation: "can_checkin",
		Object:   "ticket:abcd",
	}

	bodyIsCheckedIn := fgaSdk.ClientCheckRequest{
		User:     "ticket:abcd",
		Relation: "checked_in",
		Object:   "show:sad",
	}

	isCheckedIn, err := fgaClient.Check(context.Background()).Body(bodyIsCheckedIn).Execute()
	if err != nil {
		panic(err)
	}

	if *isCheckedIn.Allowed {
		fmt.Println("is checked in")
	} else {
		fmt.Println("is not checked in")
	}

	canCheckin, err := fgaClient.Check(context.Background()).Body(bodyCanCheckin).Execute()
	if err != nil {
		panic(err)
	}

	if *canCheckin.Allowed {
		fmt.Println("can checkin")
	} else {
		fmt.Println("cannot checkin")
	}

	// Perform checkin
	checkinBody := fgaSdk.ClientWriteRequest{
		Writes: []fgaSdk.ClientTupleKey{
			{
				User:     "ticket:abcd",
				Relation: "checked_in",
				Object:   "show:sad",
			},
		},
	}

	data, err := fgaClient.Write(context.Background()).Body(checkinBody).Execute()
	if err != nil {
		panic(err)
	}
	_ = data

	isCheckedInAgain, err := fgaClient.Check(context.Background()).Body(bodyIsCheckedIn).Execute()
	if err != nil {
		panic(err)
	}

	if *isCheckedInAgain.Allowed {
		fmt.Println("is checked in")
	} else {
		fmt.Println("is not checked in")
	}
}
