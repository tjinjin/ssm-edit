package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/aws/aws-sdk-go/service/ssm"
	mySsm "github.com/tjinjin/ssm-edit/cli/ssm"
)

func Edit(profile string) {
	svc := mySsm.GetSession(profile)

	name := "test"

	input := &ssm.GetParameterInput{
		Name: &name,
	}

	resp := mySsm.GetParameter(svc, input)

	body, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// bodyをValueだけにする？
	tempfilePath := createTempfile("/tmp/"+"test.json", body)
	defer os.Remove(tempfilePath)

	editBody := editFile(tempfilePath)

	// GetParameterOutputの形式をparseしてPutParameterInputにしないといけない

	t := &ssm.GetParameterOutput{}
	b := []byte(editBody)
	json.Unmarshal(b, &t)

	// 変更がない場合更新しない

	s := &ssm.PutParameterInput{}

	s.SetName(*t.Parameter.Name)
	s.SetType(*t.Parameter.Type)
	s.SetValue(*t.Parameter.Value)
	s.SetOverwrite(true)

	mySsm.PutParameter(svc, s)
}

func createTempfile(path string, body []byte) (tempfilePath string) {
	keys := strings.Split(path, "/")
	fileName := keys[len(keys)-1]
	tempfilePath = "/tmp/" + fileName

	if err := ioutil.WriteFile(tempfilePath, body, os.ModePerm); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return
}

func editFile(path string) string {
	command := getDefaultEditor() + " " + path
	cmd := exec.Command("sh", "-c", command)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	changeFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(changeFile[:])
}

func getDefaultEditor() string {
	return os.Getenv("EDITOR")
}
