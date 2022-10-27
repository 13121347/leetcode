package DFSIsland

import (
	"fmt"
	"testing"
)

func Test_islandPerimeter(t *testing.T) {
	testArr := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'}}

	fmt.Println(numIslands(testArr))
}
