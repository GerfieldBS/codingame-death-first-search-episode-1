package main

import (
	"fmt"
	// "slices"; // cannot import slices as it is not in GOROOT and cannot modify via codingame GUI
)

func contains(slice []string, element string) bool {
	// cannot import slices as it is not in GOROOT and cannot modify via codingame GUI
	for _, el := range slice {
		if el == element {
			return true
		}
	}
	return false
}

func main() {
	// N: the total number of nodes in the level, including the gateways
	// L: the number of links
	// E: the number of exit gateways
	var N, L, E int
	_, errNLE := fmt.Scan(&N, &L, &E) // bind
	if errNLE != nil {
		fmt.Println(errNLE.Error()) // printing error will end in incorrect result and process end
	}

	var linkCombinationSlice [][]string
	for i := 0; i < L; i++ {
		// N1: N1 and N2 defines a link between these nodes
		var N1, N2 string
		_, errN1N2 := fmt.Scan(&N1, &N2) // bind
		if errN1N2 != nil {
			fmt.Println(errN1N2.Error()) // printing error will end in incorrect result and process end
		}
		linkCombinationSlice = append(linkCombinationSlice, []string{N1, N2})
	}

	var exitSlice []string
	for i := 0; i < E; i++ {
		// EI: the index of a gateway node
		var EI string
		_, errEI := fmt.Scan(&EI) // bind
		if errEI != nil {
			fmt.Println(errEI.Error()) // printing error will end in incorrect result and process end
		}
		exitSlice = append(exitSlice, EI)
	}

	for {
		// SI: The index of the node on which the Bobnet agent is positioned this turn
		var SI string
		_, errSI := fmt.Scan(&SI) // bind
		if errSI != nil {
			fmt.Println(errSI.Error()) // printing error will end in incorrect result and process end
		}

		foundExitPath := false
		for _, linkCombination := range linkCombinationSlice {
			// check for all possible link moves of agent from his current node
			if contains(linkCombination, SI) {
				for _, exitIndex := range exitSlice {
					// check if there is exit node near agent's position
					if contains(linkCombination, exitIndex) {
						foundExitPath = true
						// block the link to the exit node
						fmt.Println(linkCombination[0] + " " + linkCombination[1])
					}
				}
			}
		}
		if !foundExitPath {
			// default link combination if agent is not near the exit mode
			fmt.Println(linkCombinationSlice[0][0] + " " + linkCombinationSlice[0][1])
			// valid solution is also this:
			// fmt.Println("0 1")
			// but as string is hardcoded, game won't give you 100% score but 50%
			// if you use index of slice, it will give you 100%
		}
	}
}
