package cli

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

func Tui(profile, region string) {
	fmt.Println("test")
	key := listTui()
	fmt.Println(key)
}

func listTui() string {
	prompt := promptui.Select{
		Label: "Select profile",
		Items: []string{"aaa", "bbb", "ccc"},
		Size:  10,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return result
}
