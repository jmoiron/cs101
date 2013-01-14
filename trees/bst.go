package main

import (
	"fmt"
	"strings"
)

/* square i */
func sq(i int) int { return i * i }

/* fast int power, can be used to pprint properly and pre-allocate slices */
func pow(x, n int) int {
	r := 1
	for {
		if n&0x00000001 == 1 {
			r *= x
		}
		n /= 2
		if n == 0 {
			break
		}
		x *= x
	}
	return r
}

/* return the max of the ints passed;  returns 0 when nothing is passed */
func max(integers ...int) int {
	if len(integers) == 0 {
		return 0
	}
	var m int = integers[0]
	for i := 1; i < len(integers); i++ {
		if integers[i] > m {
			m = integers[i]
		}
	}
	return m
}

/* a simple binary tree node type */
type Node struct {
	val   int
	left  *Node
	right *Node
}

func (n *Node) isEmpty() bool {
	return n.left == nil && n.right == nil
}

/* add a value as a new node to the right place in the bst */
func (n *Node) Add(val int) {
	var sub **Node
	if val > n.val {
		sub = &n.right
	} else {
		sub = &n.left
	}
	if *sub != nil {
		(*sub).Add(val)
	} else {
		*sub = &Node{val, nil, nil}
	}
}

/* Return an in-order (dfs) string of the node values */
func (n *Node) InOrderString() string {
	var left, right string
	if n.left != nil {
		left = n.left.InOrderString()
	}
	if n.right != nil {
		right = n.right.InOrderString()
	}
	if len(left) > 0 || len(right) > 0 {
		return fmt.Sprintf("%s %d %s", left, n.val, right)
	}
	return fmt.Sprintf("%d", n.val)
}

/* the height of the tree (twitter..) */
func (n *Node) Height() int {
	subs := []int{}
	if n.left != nil {
		subs = append(subs, n.left.Height())
	}
	if n.right != nil {
		subs = append(subs, n.right.Height())
	}
	return max(subs...) + 1
}

/* search the binary tree for a value */
func (n *Node) Has(val int) bool {
	if n.val == val {
		return true
	}
	if val > n.val {
		if n.right == nil {
			return false
		}
		return n.right.Has(val)
	}
	if n.left == nil {
		return false
	}
	return n.left.Has(val)
}

/* print a tree out in somewhat pretty fashion */
func (n *Node) PrettyString() string {
	if n.isEmpty() {
		return ""
	}
	h := n.Height()
	ss := make([][]string, h)
	levels := make([][]*Node, h)
	levels[0] = []*Node{n}
	ss[0] = []string{fmt.Sprint(n.val)}
	for i := 1; i < n.Height(); i++ {
		/* allocate slices of 2**i */
		levels[i] = make([]*Node, pow(2, i))
		ss[i] = make([]string, pow(2, i))
		for j := 0; j < len(levels[i-1]); j++ {
			parent := levels[i-1][j]
			for k, ptr := range []*Node{parent.left, parent.right} {
				pos := j*2 + k
				if ptr != nil {
					levels[i][pos] = ptr
					ss[i][pos] = fmt.Sprint(ptr.val)
				} else {
					levels[i][pos] = &Node{0, nil, nil}
					ss[i][pos] = "-"
				}
			}
		}
	}
	s := ""
	for i := 0; i < h; i++ {
		spacing := max(0, (sq(h-i) / 2))
		s += strings.Repeat(" ", spacing) + fmt.Sprintf("%v\n", ss[i])
	}
	return s
}

func main() {
	tree := &Node{10, nil, nil}
	for _, val := range []int{5, 12, 3, 13, 4, 19, 21, 11, 15} {
		tree.Add(val)
	}
	fmt.Println(tree.PrettyString())
	for _, val := range []int{5, 29, 6, 11} {
		fmt.Printf("Has %d? %v\n", val, tree.Has(val))
	}

	for _, pair := range [][]int{[]int{2, 3}, []int{2, 5}, []int{3, 4}, []int{2, 27}} {
		fmt.Printf("%d ** %d = %d\n", pair[0], pair[1], pow(pair[0], pair[1]))
	}
}
