package main

import (
	"fmt"
	"os"
)

type agents []string

func main() {
	agentsfilePath := os.Args[1]
	functionFilePath := os.Args[2]

	agents := getAgents(agentsfilePath)

	fmt.Println("Agents: ", agents)
     
	coalitionSpace := make(map[int][]coalition)

	for size := 1; size <= len(agents); size++ {
		buildCoalitionSpace(agents, functionFilePath, size, coalitionSpace)
	}

	fmt.Println("coalition space: ", coalitionSpace)

}
