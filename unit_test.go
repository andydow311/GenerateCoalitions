package main

import (
	"testing"
)

func TestAgentsAreCorrectlyGenerated(t *testing.T) {
	agents := getAgents("agents_test.txt")

	if agents != "AB" {
		t.Errorf("Expected value of ACEF but got %v", agents)
	}
}

func TestComputeCoalitionValue(t *testing.T) {

	value := computeCoalitionValue("AB", "function_data_test.txt")

	if value != 1 {
		t.Errorf("Expected value of 1 but got %v", value)
	}

	valueTwo := computeCoalitionValue("AC", "function_data_test.txt")

	if valueTwo != -1 {
		t.Errorf("Expected value of -1 but got %v", value)
	}
}

func TestOrderCoalitions(t *testing.T) {

	value := orderCoalitions("FECA")

	if value != "ACEF" {
		t.Errorf("Expected value of ACEF but got %v", value)
	}

}

func TestCoalitionSpaceCorrectlyGenerated(t *testing.T) {

	coalitionSpace := make(map[int][]coalition)

	buildCoalitionSpace("AB", "function_data_test.txt", 2, coalitionSpace)
	buildCoalitionSpace("AB", "function_data_test.txt", 1, coalitionSpace)

	coaltionsOfSizeTwo := coalitionSpace[2]
	coaltionsOfSizeOne := coalitionSpace[1]

	if len(coaltionsOfSizeTwo) != 1 {
		t.Errorf("Expected value of 1 but got %v", len(coaltionsOfSizeTwo))
	}

	if len(coaltionsOfSizeOne) != 2 {
		t.Errorf("Expected value of 1 but got %v", len(coaltionsOfSizeOne))
	}
}
