package tree

import "testing"
import "strconv"
import "fmt"

func AddAndCountTest(t *testing.T) {
	tree := Binary{}
	tree.Add(&Element{Value: 1})
	if tree.Count() != 1 {
		fmt.Println("Count was " + strconv.Itoa(tree.Count()) + ", should have been 1.")
		t.Fail()
	}
}
