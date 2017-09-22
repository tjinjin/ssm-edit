package cli

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/ssm"
	mySsm "github.com/tjinjin/ssm-edit/cli/ssm"
)

func List(profile string) {
	svc := mySsm.GetSession(profile)

	input := &ssm.DescribeParametersInput{}
	resp, err := svc.DescribeParameters(input)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, i := range resp.Parameters {
		fmt.Printf("Name: %s Type: %s\n", *i.Name, *i.Type)
	}
}
