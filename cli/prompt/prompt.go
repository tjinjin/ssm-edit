package cli

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/manifoldco/promptui"
	mySsm "github.com/tjinjin/ssm-edit/cli/ssm"
)

func Prompt(profile, region string, size int) string {
	svc := mySsm.GetSession(profile, region)

	results := mySsm.ListParameters(svc)
	var items []string
	for _, i := range results {
		for _, j := range i.Parameters {
			items = append(items, *j.Name)
		}
	}

	sort.Strings(items)
	return listSsm(items, size)
}

func listSsm(items []string, size int) string {
	searcher := func(input string, index int) bool {
		i := strings.ToLower(items[index])
		return strings.Contains(i, input)
	}
	prompt := promptui.Select{
		Label:    "Select ssm parameters",
		Items:    items,
		Size:     size,
		Searcher: searcher,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return result
}
