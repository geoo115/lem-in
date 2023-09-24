package main

import (
	"fmt"
	lem "lem/methods"
	"log"
	"os"
	"strings"
)

func main() {
	// fmt.Println("Hello Gophers!")
	// server.StartServer()

	if len(os.Args) != 2 {
		fmt.Println("❌ ERROR: Usage EX: go run . example00.txt")
		return
	}

	// Read the file
	content, err := os.ReadFile("examples/" + os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
	dataFile := strings.Split(string(content), "\n")
	if len(dataFile) == 0 {
		fmt.Println("❌ invalid file: the file is empty")
	}

	// Create an Input struct and populate it with the file data
	data := &lem.Input{}
	err = lem.GetData(data, dataFile)
	if err != nil {
		log.Fatal(err)
	}

	// Create a graph
	graph := &lem.Graph{}

	// Adding rooms
	for _, v := range data.Rooms {
		if err := lem.AddRoom(graph, v); err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	// Adding tunnels (links) between rooms
	for _, v := range data.Links {
		if err := lem.AddEdge(graph, v[0], v[1]); err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	// BFS to find all paths
	allPaths := &lem.Allpaths{}
	allPaths.Paths = lem.BFS(graph, data.StartR, data.EndR)
	if len(allPaths.Paths) == 0 {
		fmt.Println("❌ There is no path between Start and End rooms")
		return
	}

	// Generate path combinations
	lem.GenerateCombinations(allPaths, graph)

	// Calculate Time for each combination
	time := lem.CalculateTravelTimesForEachCombination(allPaths, data.Ants)

	// Identify the optimal path
	lem.FindOptimalPath(allPaths, time, graph)
	lem.PrintAntMovement(allPaths, data.Ants)

	fmt.Println("-----------All Paths-------------")
	fmt.Println()
	for i, path := range allPaths.Paths {
		fmt.Printf("Path %d:\n", i)
		for j, room := range path.P {
			fmt.Printf("%s", room.Key)
			if j < len(path.P)-1 { // Check if it's not the last room in the path
				fmt.Printf(" -> ")
			}
		}
		fmt.Println()
	}

	// fmt.Println("-----------All Paths-------------")
	// fmt.Println()
	// for i, path := range allPaths.Paths {
	// 	fmt.Printf("Path %d:\n", i)
	// 	for _, room := range path.P {
	// 		fmt.Printf("%s -> ", room.Key)
	// 		fmt.Println()
	// 	}

	// }

	fmt.Println("---------All Combinations--------")
	for i, combo := range allPaths.Combo {
		fmt.Println()
		fmt.Printf("Combination #️⃣ %d:\n", i)
		for j, path := range combo {
			fmt.Printf("🐜 Path %d:\n", j)
			for k, room := range path.P {
				fmt.Printf("%s", room.Key)
				if k < len(path.P)-1 {
					fmt.Print(" -> ")
				}

			}
			fmt.Println()
		}
	}
	// for i, combo := range allPaths.Combo {
	// 	fmt.Printf("Combination %d:\n", i)
	// 	for j, path := range combo {
	// 		fmt.Printf("Path %d:\n", j)
	// 		fmt.Println()
	// 		for _, room := range path.P {
	// 			fmt.Printf("%s -> ", room.Key)
	// 			fmt.Println()
	// 		}
	// 		// fmt.Println()
	// 	}
	// }

	// fmt.Println("----------Optimal Path-----------")
	// fmt.Println()
	// for _, room := range allPaths.OptPath[0].P {
	// 	fmt.Printf("%s -> ", room.Key)
	// 	// fmt.Println()
	// }
	// fmt.Println()

	fmt.Println("----------Optimal Path-----------")
	fmt.Println()
	optimalPath := allPaths.OptPath[0].P
	for i, room := range optimalPath {
		fmt.Printf("%s", room.Key)
		if room.Key != "end" && i < len(optimalPath)-1 {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()

	fmt.Println("-----------Travel Time-----------")
	fmt.Println()
	for i, time := range time {
		fmt.Printf("Combination %d: %d milliseconds \n", i, time)
		fmt.Println()
	}

}
