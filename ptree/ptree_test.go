package ptree

import (
	"testing"
)

func TestInsert(t *testing.T) {
	root := New()

	root.Insert("cat")

	if !root.SubNodes['c'].SubNodes['a'].SubNodes['t'].Complete {
		t.Error("Expected cat to be inserted")
	}
}

func TestInsertMultiple(t *testing.T) {
	root := New()

	root.Insert("cat")
	root.Insert("cats")

	if !root.SubNodes['c'].SubNodes['a'].SubNodes['t'].Complete {
		t.Error("Expected cat to be inserted")
	}

	if !root.SubNodes['c'].SubNodes['a'].SubNodes['t'].SubNodes['s'].Complete {
		t.Error("Expected cats to be inserted")
	}
}

func TestMerge(t *testing.T) {
	root1 := New()
	root1.Insert("cat")

	root2 := New()
	root2.Insert("dog")

	root1.Merge(root2)

	if !root1.SubNodes['d'].SubNodes['o'].SubNodes['g'].Complete {
		t.Error("Expected dog to be merged")
	}
}

func TestClear(t *testing.T) {
	root := New()
	root.Insert("cat")

	root.Clear()

	if root.Len != 0 {
		t.Error("Expected tree to be cleared")
	}
}

// func TestWalk(t *testing.T) {

// 	root := New()
// 	root.Insert("cat")
// 	root.Insert("dog")

// 	var output []byte
// 	root.Walk(func(n *Node) {
// 		length := []byte(strconv.Itoa(n.Len))
// 		output = append(output, length...)
// 	})

// 	expected := []byte("1011")
// 	if !reflect.DeepEqual(output, expected) {
// 		t.Errorf("Walk output %v, expected %v", output, expected)
// 	}

// }

// func TestWalk(t *testing.T) {

// 	root := New()

// 	root.Insert("cat")
// 	root.Insert("can")
// 	//root.Insert("dog")
// 	//root.Insert("gov")

// 	var actual []string

// 	// Walk and print each node
// 	root.Walk(func(n *Node) {
// 		fmt.Println(n, fmt.Sprint(n.Len))

// 		actual = append(actual, fmt.Sprint(n.Len))
// 	})

// 	expected := []string{"1", "2", "3", "1", "0", "1"}

// 	if !reflect.DeepEqual(actual, expected) {
// 		t.Errorf("Walk output %v, expected %v", actual, expected)
// 	}

// }
