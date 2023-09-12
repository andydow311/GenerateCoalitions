package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type coalition struct {
	agents string
	value  float64
}

func buildCoalition(agents string, value float64) coalition {
	return coalition{agents: agents, value: value}
}

func buildCoalitionSpace(agents string, filePath string, size int, coalitions map[int][]coalition) {
	coalitionSpace := []coalition{}

	if size == 1 {
		for _, character := range agents {
			coalitionSpace = append(coalitionSpace, buildCoalition(string(character), computeCoalitionValue(string(character), filePath)))
		}
	} else if size == len(agents) {
		coalitionSpace = append(coalitionSpace, buildCoalition(agents, computeCoalitionValue(agents, filePath)))
	} else {
		prev_cs := coalitions[size-1]
		for _, character := range agents {
			for _, coalition := range prev_cs {
				if !strings.Contains(coalition.agents, string(character)) {
					new_agents := coalition.agents + string(character)
					c := buildCoalition(new_agents, computeCoalitionValue(string(new_agents), filePath))
					if !coalitionAlreadyConsidered(coalitionSpace, c) {
						coalitionSpace = append(coalitionSpace, c)
					}

				}
			}
		}
	}
	coalitions[size] = coalitionSpace
}

func orderCoalitions(coaltion string) string {
	orderedCoaltion := strings.Split(coaltion, "")
	sort.Strings(orderedCoaltion)
	return strings.Join(orderedCoaltion, "")
}

func coalitionAlreadyConsidered(coalitionSpace []coalition, coalition coalition) bool {

	for _, existingCoalition := range coalitionSpace {
		if orderCoalitions(existingCoalition.agents) == orderCoalitions(coalition.agents) {
			return true
		}
	}

	return false

}

func computeCoalitionValue(agents string, filePath string) float64 {

	var value = 0.0

	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		firstAgent := strings.Split(fileScanner.Text(), "	")[0]
		secondAgent := strings.Split(fileScanner.Text(), "	")[1]
		synergy := strings.Split(fileScanner.Text(), "	")[2]
		externality := strings.Split(fileScanner.Text(), "	")[3]

		if strings.Contains(agents, firstAgent) && strings.Contains(agents, secondAgent) {
			synergy, err := strconv.ParseFloat(synergy, 64)
			if err != nil {
				fmt.Println(err)
			}
			value = value + synergy
		} else if strings.Contains(agents, firstAgent) || strings.Contains(agents, secondAgent) {
			externality, err := strconv.ParseFloat(externality, 64)
			if err != nil {
				fmt.Println(err)
			}
			value = value + externality
		}
	}

	readFile.Close()

	return value

}
