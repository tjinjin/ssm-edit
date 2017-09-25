package cli

import (
	"fmt"

	mySsm "github.com/tjinjin/ssm-edit/cli/ssm"
)

func List(profile string) {
	svc := mySsm.GetSession(profile)

	resp := mySsm.DescribeParameters(svc)

	for _, i := range resp.Parameters {
		fmt.Printf("Name: %s Type: %s\n", *i.Name, *i.Type)
	}
}