package avltree

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

func PrintTree[T Comparable[T]](tree UniqAvlTree[T]) {
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

//function printBinaryTree(root) {
//let treeLevel = [], temp = [];
//treeLevel.push(root);
//let counter = 0;
//let height = heightOfTree(root) - 1;
//let numberOfElements = Math.pow(2, (height + 1)) - 1;
//while (counter <= height) {
//let removed = treeLevel.shift();
//if (temp.length == 0) {
//printSpace(numberOfElements / Math.pow(2, counter + 1), removed);
//} else {
//printSpace(numberOfElements / Math.pow(2, counter), removed);
//}
//if (removed == null) {
//temp.push(null);
//temp.push(null);
//} else {
//temp.push(removed.left);
//temp.push(removed.right);
//}
//if (treeLevel.length == 0) {
//console.log("\n");
//treeLevel = temp;
//temp = [];
//counter++;
//}
//}
//}
//
