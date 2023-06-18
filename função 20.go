package main

import (
	"errors"
	"fmt"
)

func filterStringsByLength(strings []string) ([]string, error) {
	if len(strings) == 0 {
		return nil, errors.New("Erro: slice vazio")
	}

	filteredStrings := []string{}

	for _, str := range strings {
		if len(str) > 5 {
			filteredStrings = append(filteredStrings, str)
		}
	}

	return filteredStrings, nil
}

func main() {
	strings := []string{"Hello", "World", "Go", "Programming"}

	filtered, err := filterStringsByLength(strings)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Strings filtradas:", filtered)
	}

	emptySlice := []string{}

	filtered, err = filterStringsByLength(emptySlice)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Strings filtradas:", filtered)
	}
}
