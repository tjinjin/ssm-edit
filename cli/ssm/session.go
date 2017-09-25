package cli

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func GetSession(profile string) (svc *ssm.SSM) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String("ap-northeast-1")},
		Profile: profile,
	}))
	test := ssm.New(sess)

	return test
}

func DescribeParameters(svc *ssm.SSM) (resp *ssm.DescribeParametersOutput) {

	input := &ssm.DescribeParametersInput{}
	resp, err := svc.DescribeParameters(input)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return resp
}

func GetParameter(svc *ssm.SSM, input *ssm.GetParameterInput) (resp *ssm.GetParameterOutput) {
	resp, err := svc.GetParameter(input)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return resp
}

func PutParameter(svc *ssm.SSM, input *ssm.PutParameterInput) {
	_, err := svc.PutParameter(input)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
