package shopee

type MyCircularDeque struct {
	q        []int
	capacity int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		q:        []int{},
		capacity: k,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	this.q = append([]int{value}, this.q...)
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}

	this.q = append(this.q, value)
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}

	this.q = this.q[1:]
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}

	this.q = this.q[:len(this.q)-1]
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}

	return this.q[0]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.q[len(this.q)-1]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return len(this.q) == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return len(this.q) == this.capacity
}
