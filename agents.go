package main

import (
	"bufio"
	"fmt"
	"os"
)


func getAgents(filePath string) string {

	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	agents := ""

	for fileScanner.Scan() {
		agent := fileScanner.Text()
		agents = agents + agent
	}

	readFile.Close()
	return agents
}
