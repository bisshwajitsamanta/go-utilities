package main

import (
	"fmt"
	"github.com/jedib0t/go-pretty/table"
	"log"
	"os"
	"os/exec"
)

type CountDeployment struct {
	Tables []struct {
		Content string `json:"Content"`
		Rows    []struct {
			Name string `json:"name"`
		} `json:"Rows"`
	} `json:"Tables"`
}

func boshWatch(cmd string, argsArr []string) ([]byte, error) {
	args := argsArr
	out := exec.Command(cmd, args...)

	out.Stdout = os.Stdout
	out.Stdin = os.Stdin
	out.Stderr = os.Stderr

	err := out.Run()
	if err != nil {
		log.Fatal(err)

	}
	return nil, nil
}

func countDeployments() {

}

func main() {

	deploymentList, err := boshWatch("bosh", []string{"ds", "--json"})
	if err != nil {
		log.Println("Unable to Connect to Bosh Instances!")
	}
	fmt.Println(deploymentList)

	// Dummy Data

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Environment", "Health Status", "Number of Deployments", "Remarks"})
	t.AppendRows([]table.Row{
		{1, "AWS-Dev1", "PASSING", 3, "All VMS are Passing"},
		{2, "AWS-Dev2", "PASSING", 4, "All VMS are Passing"},
		{3, "AWS-Dev3", "PASSING", 5, "All VMS are Passing"},
		{4, "AWS-Dev4", "PASSING", 6, "All VMS are Passing"},
	})
	t.AppendSeparator()
	t.AppendRows([]table.Row{
		{1, "Schwab-UAT", "PASSING", "All VMS are Passing"},
		{2, "Schwab-PROD", "PASSING", "All VMS are Passing"},
	})
	t.Render()
}
