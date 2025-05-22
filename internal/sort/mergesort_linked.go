package sort

import "MergSortGoLanguage/internal/structures"

// MergeSortLinked aplica merge sort em uma lista encadeada
func MergeSortLinked(head *structures.Node) *structures.Node {
    if head == nil || head.Next == nil {
        return head
    }

    // Divide a lista ao meio
    middle := getMiddle(head)
    nextOfMiddle := middle.Next
    middle.Next = nil

    // Recursão para cada metade
    left := MergeSortLinked(head)
    right := MergeSortLinked(nextOfMiddle)

    // Junta as listas ordenadas
    sorted := mergeLinked(left, right)
    return sorted
}

func getMiddle(head *structures.Node) *structures.Node {
    if head == nil {
        return head
    }

    slow := head
    fast := head

    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    return slow
}

func mergeLinked(a, b *structures.Node) *structures.Node {
    if a == nil {
        return b
    }
    if b == nil {
        return a
    }

    var result *structures.Node

    if a.Value <= b.Value {
        result = a
        result.Next = mergeLinked(a.Next, b)
    } else {
        result = b
        result.Next = mergeLinked(a, b.Next)
    }

    return result
}
