package cli

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/ssm"
	mySsm "github.com/tjinjin/ssm-edit/cli/ssm"
)

func List(profile string, region string) {

	var results []*ssm.DescribeParametersOutput

	svc := mySsm.GetSession(profile, region)

	input := &ssm.DescribeParametersInput{}
	resp := mySsm.DescribeParameters(svc, input)

	results = append(results, resp)

	input.NextToken = resp.NextToken

	for input.NextToken != nil {
		resp := mySsm.DescribeParameters(svc, input)

		results = append(results, resp)
		input.NextToken = resp.NextToken
	}

	for _, i := range results {
		for _, j := range i.Parameters {
			fmt.Printf("Name: %s Type: %s\n", *j.Name, *j.Type)
		}
	}
}
