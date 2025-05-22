package structures

// ==========================
// Pilha Linear (slice)
// ==========================
type StackLinear struct {
    items []int
}

func NewStackLinear() *StackLinear {
    return &StackLinear{items: []int{}}
}

func (s *StackLinear) Push(value int) {
    s.items = append(s.items, value)
}

func (s *StackLinear) Pop() (int, bool) {
    if len(s.items) == 0 {
        return 0, false
    }
    val := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return val, true
}

func (s *StackLinear) ToSlice() []int {
    return append([]int{}, s.items...) // cópia
}

// ==========================
// Pilha Dinâmica (ponteiros)
// ==========================
type StackNode struct {
    value int
    next  *StackNode
}

type StackDynamic struct {
    top *StackNode
}

func NewStackDynamic() *StackDynamic {
    return &StackDynamic{nil}
}

func (s *StackDynamic) Push(value int) {
    s.top = &StackNode{value, s.top}
}

func (s *StackDynamic) Pop() (int, bool) {
    if s.top == nil {
        return 0, false
    }
    val := s.top.value
    s.top = s.top.next
    return val, true
}

func (s *StackDynamic) ToSlice() []int {
    var result []int
    for node := s.top; node != nil; node = node.next {
        result = append(result, node.value)
    }
    // Como a pilha é LIFO, o slice está em ordem reversa
    // Vamos inverter para manter consistência
    for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
        result[i], result[j] = result[j], result[i]
    }
    return result
}
