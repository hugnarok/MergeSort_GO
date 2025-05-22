package structures

type Node struct {
    Value int
    Next  *Node
}

// FromSlice cria uma lista encadeada a partir de um slice
func FromSlice(data []int) *Node {
    if len(data) == 0 {
        return nil
    }
    head := &Node{Value: data[0]}
    curr := head
    for _, v := range data[1:] {
        curr.Next = &Node{Value: v}
        curr = curr.Next
    }
    return head
}

// ToSlice transforma uma lista encadeada em slice
func (n *Node) ToSlice() []int {
    var result []int
    for curr := n; curr != nil; curr = curr.Next {
        result = append(result, curr.Value)
    }
    return result
}
