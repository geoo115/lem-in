package lem_test

import (
	"fmt"
	"io/ioutil"
	lem "lem/methods"
	"os"
	"strings"
	"testing"
)

func TestPrintAntMovement(t *testing.T) {
	// Create sample data for testing
	graph := &lem.Graph{}
	_ = lem.AddRoom(graph, "0")
	_ = lem.AddRoom(graph, "2")
	_ = lem.AddRoom(graph, "3")
	_ = lem.AddRoom(graph, "1")
	_ = lem.AddEdge(graph, "0", "2")
	_ = lem.AddEdge(graph, "2", "3")
	_ = lem.AddEdge(graph, "3", "1")

	p := &lem.Allpaths{
		OptPath: []*lem.Path{
			{
				P: []*lem.Room{graph.Start, graph.Rooms[1], graph.Rooms[2], graph.Rooms[3]},
			},
		},
		MinTravelTime: 7,
	}

	ants := 4
	// Redirect standard output for testing
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function to be tested
	lem.PrintAntMovement(p, ants)

	// Close the temporary writer and capture the output
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = oldStdout

	// Define expected output for comparison
	expectedOutput := "----------- Ants Movements ------------- " + "L1-2 \n" + "L1-3 L2-2 \n" + "L1-1 L2-3 L3-2 \n" +
		"L2-1 L3-3 L4-2 \n" + "L3-1 L4-3 \n" + "L4-1 \n" + "Number of turns: 6"

	// Trim spaces and newlines from both expected and actual outputs
	expectedOutput = strings.TrimSpace(expectedOutput)
	actualOutput := strings.TrimSpace(string(out))

	// Normalize whitespace (replace multiple spaces and newlines with a single space)
	expectedOutput = strings.Join(strings.Fields(expectedOutput), " ")
	actualOutput = strings.Join(strings.Fields(actualOutput), " ")

	// Print hex dumps for comparison and debugging
	fmt.Printf("Expected hex dump:\n%x\n", []byte(expectedOutput))
	fmt.Printf("Actual hex dump:\n%x\n", []byte(actualOutput))

	// Compare actual output with expected output
	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, actualOutput)
	}
}
