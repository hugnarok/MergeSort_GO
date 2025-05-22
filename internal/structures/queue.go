package structures

// ==========================
// Fila Linear (slice)
// ==========================
type QueueLinear struct {
    items []int
}

func NewQueueLinear() *QueueLinear {
    return &QueueLinear{items: []int{}}
}

func (q *QueueLinear) Enqueue(value int) {
    q.items = append(q.items, value)
}

func (q *QueueLinear) Dequeue() (int, bool) {
    if len(q.items) == 0 {
        return 0, false
    }
    val := q.items[0]
    q.items = q.items[1:]
    return val, true
}

func (q *QueueLinear) ToSlice() []int {
    return append([]int{}, q.items...) // cópia
}

// ==========================
// Fila Dinâmica (ponteiros)
// ==========================
type QueueNode struct {
    value int
    next  *QueueNode
}

type QueueDynamic struct {
    head *QueueNode
    tail *QueueNode
}

func NewQueueDynamic() *QueueDynamic {
    return &QueueDynamic{}
}

func (q *QueueDynamic) Enqueue(value int) {
    newNode := &QueueNode{value, nil}
    if q.tail != nil {
        q.tail.next = newNode
    }
    q.tail = newNode
    if q.head == nil {
        q.head = newNode
    }
}

func (q *QueueDynamic) Dequeue() (int, bool) {
    if q.head == nil {
        return 0, false
    }
    val := q.head.value
    q.head = q.head.next
    if q.head == nil {
        q.tail = nil
    }
    return val, true
}

func (q *QueueDynamic) ToSlice() []int {
    var result []int
    for node := q.head; node != nil; node = node.next {
        result = append(result, node.value)
    }
    return result
}
