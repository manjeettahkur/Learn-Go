package queue


type Queue struct {
	Items []int
}


func (q *Queue) Push(val int) {
	q.Items = append(q.Items, val)
}

func (q *Queue) Pop()  {
	 q.Items = q.Items[1:]
}