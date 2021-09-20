package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	extractedList := ConnectorName()
	fmt.Println(extractedList)

	//connectorList := []string{"gbiConnector", "PersistentDBConnector", "randomConnector"}
}

func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func ConnectorName() (tokenList []string) {
	// Opening a file
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Opening file error", err)
	}
	// Reading a file
	scanner := bufio.NewScanner(file)
	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()
	for _, eachline := range txtlines {
		parse := strings.Fields(eachline)
		for _, v := range parse {
			tokens := strings.Split(v, "=")
			if len(tokens) > 1 {
				if tokens[0] == "connectorName" {
					tokenList = Unique(append(tokenList,tokens[1:]...))
				}
			}
		}
	}
	return tokenList
}

func Unique(list []string) []string {
	seen := make(map[string]struct{})
	for _, item := range list {
		// Mark `item` as seen by adding it to the map
		seen[item] = struct{}{}
	}
	var res []string
	// Gather all items that were seen into a slice to return it
	for item := range seen {
		res = append(res, item)
	}
	return res
}
