package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go/service/ssm"
	prompt "github.com/tjinjin/ssm-edit/cli/prompt"
	mySsm "github.com/tjinjin/ssm-edit/cli/ssm"
)

func Edit(profile string, region string, name string, size int) {
	svc := mySsm.GetSession(profile, region)
	if name == "" {
		name = prompt.Prompt(profile, region, size)
	}

	withDecryption := true

	input := &ssm.GetParameterInput{
		Name:           &name,
		WithDecryption: &withDecryption,
	}

	resp := mySsm.GetParameter(svc, input)

	body, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// bodyをValueだけにする？
	tempfilePath := createTempfile(name, body)
	defer os.Remove(tempfilePath)

	editBody := editFile(tempfilePath)

	t := &ssm.GetParameterOutput{}
	b := []byte(editBody)
	json.Unmarshal(b, &t)

	// compare []uint8 to []uint8
	if reflect.DeepEqual(body, b) {
		fmt.Println("No Changed")
		return
	}

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
	tempfilePath = "/tmp/" + fileName + ".json"

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
