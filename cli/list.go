package cli

import (
	"fmt"

	mySsm "github.com/tjinjin/ssm-edit/cli/ssm"
)

func List(profile string, region string) {

	svc := mySsm.GetSession(profile, region)

	results := mySsm.ListParameters(svc)
	for _, i := range results {
		for _, j := range i.Parameters {
			fmt.Printf("Name: %s Type: %s\n", *j.Name, *j.Type)
		}
	}
}
