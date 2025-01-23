package heap

import "fmt"

type MinHeap struct {
	elements []int
}

func (m *MinHeap) Insert(val int) {
	m.elements = append(m.elements, val)
	m.MinHeapfiyUp(len(m.elements) - 1)
}

func (m *MinHeap) MinHeapfiyUp(index int) {
	for m.elements[parent(index)] < m.elements[index] {
		m.Swap(parent(index), index)
		index = parent(index)
	}
}

func (m *MinHeap) Extract() int{
	extracted := m.elements[0]
	fmt.Println(extracted)
	m.elements = append(m.elements[:0], m.elements[1:]... )
	return extracted
}

func parent(i int) int {
	return (i - 1) / 2
}

func Left(i int) int {
	return 2*i + 1
}

func Right(i int) int {
	return 2*i + 2
}

func (m *MinHeap) Swap(i1, i2 int) {
	m.elements[i1], m.elements[i2] = m.elements[i2], m.elements[i1]
}
