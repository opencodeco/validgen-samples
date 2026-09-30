package main

import (
	"fmt"

	"github.com/opencodeco/validgen-samples/account"
)

func main() {
	profiles := []*account.Profile{
		{},
		{FirstName: "Ada", LastName: "Lovelace", Age: 36},
	}
	for _, profile := range profiles {
		errs := account.ProfileValidate(profile)
		if len(errs) == 0 {
			fmt.Printf("%+v is valid\n", profile)
			continue
		}
		fmt.Printf("%+v\n", profile)
		for _, err := range errs {
			fmt.Printf("  %s\n", err)
		}
	}
}
