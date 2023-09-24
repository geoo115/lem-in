// package lem_test

// import (
// 	lem "lem/methods"
// 	"testing"
// )

// func TestAddEdge(t *testing.T) {
// 	graph := &lem.Graph{}
// 	_ = lem.AddRoom(graph, "Room1")
// 	_ = lem.AddRoom(graph, "Room2")

// 	err := lem.AddEdge(graph, "Room1", "Room2")
// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
// 	}

// 	err = lem.AddEdge(graph, "Room1", "Room1")
// 	if err == nil {
// 		t.Error("Expected error, but got nil")
// 	}

// 	err = lem.AddEdge(graph, "Room1", "Room2")
// 	if err == nil {
// 		t.Error("Expected error, but got nil")
// 	}
// }

package lem_test

import (
	"testing"

	lem "lem/methods" // Import the package under test
)

func TestAddEdge(t *testing.T) {
	graph := &lem.Graph{}
	_ = lem.AddRoom(graph, "Room1")
	_ = lem.AddRoom(graph, "Room2")

	t.Run("Valid Edge", func(t *testing.T) {
		err := lem.AddEdge(graph, "Room1", "Room2")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
	})

	t.Run("Invalid Self-Edge", func(t *testing.T) {
		err := lem.AddEdge(graph, "Room1", "Room1")
		if err == nil {
			t.Error("Expected error for self-edge, but got nil")
		} else {
			expectedErrMsg := "❌ ERROR: invalid edge(Room1 --> Room1)"
			if err.Error() != expectedErrMsg {
				t.Errorf("Expected error message '%s', but got '%s'", expectedErrMsg, err)
			}
		}
	})

	t.Run("Existing Edge", func(t *testing.T) {
		err := lem.AddEdge(graph, "Room1", "Room2")
		if err == nil {
			t.Error("Expected error for existing edge, but got nil")
		} else {
			expectedErrMsg := "❌ ERROR: existing edge(Room1 --> Room2)"
			if err.Error() != expectedErrMsg {
				t.Errorf("Expected error message '%s', but got '%s'", expectedErrMsg, err)
			}
		}
	})
}
