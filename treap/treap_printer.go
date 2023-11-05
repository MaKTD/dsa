package treap

import (
	"fmt"
	"math"
)

func printSpaceAndElement[T Comparable[T]](n int, removed *node[T]) {
	for i := 0; i < n; i++ {
		fmt.Print(" ")
	}
	if removed == nil {
		fmt.Print(" ")
	} else {
		fmt.Printf("%v", removed.val)
	}
}

func getHeightOfTree[T Comparable[T]](root *node[T]) int {
	if root == nil {
		return 0
	}

	leftH := getHeightOfTree(root.left)
	rightH := getHeightOfTree(root.right)

	if leftH > rightH {
		return 1 + leftH
	} else {
		return 1 + rightH
	}
}

func PrintTree[T Comparable[T]](tree UniqTreap[T]) {
	treeLevel := []*node[T]{tree.root}
	temp := make([]*node[T], 0)
	counter := 0
	height := getHeightOfTree(tree.root)
	numberOfElemets := int(math.Pow(2, float64(height+1)) - 1)

	for counter <= height {
		removed := treeLevel[0]
		treeLevel = treeLevel[1:]
		if len(temp) == 0 {
			printSpaceAndElement(numberOfElemets/int(math.Pow(2, float64(counter+1))), removed)
		} else {
			printSpaceAndElement(numberOfElemets/int(math.Pow(2, float64(counter))), removed)
		}

		if removed == nil {
			temp = append(temp, nil)
			temp = append(temp, nil)
		} else {
			temp = append(temp, removed.left)
			temp = append(temp, removed.right)
		}
		if len(treeLevel) == 0 {
			fmt.Print("\n")
			treeLevel = temp
			temp = make([]*node[T], 0)
			counter += 1
		}

	}

}
