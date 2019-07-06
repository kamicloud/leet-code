//Design a stack that supports push, pop, top, and retrieving the minimum element in constant time. 
//
// 
// push(x) -- Push element x onto stack. 
// pop() -- Removes the element on top of the stack. 
// top() -- Get the top element. 
// getMin() -- Retrieve the minimum element in the stack. 
// 
//
// 
//
// Example: 
//
// 
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> Returns -3.
//minStack.pop();
//minStack.top();      --> Returns 0.
//minStack.getMin();   --> Returns -2.
// 
//
// 
//

type ListNode struct {
	Val  int
	Next *ListNode
}

type MinStack struct {
	elements []*ListNode
	delist *ListNode
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		elements: []*ListNode{},
		delist: nil,
	}
}


func (this *MinStack) Push(x int)  {
	node := ListNode{
		Val: x,
	}
	this.elements = append(this.elements, &node)

	if this.delist == nil {
		this.delist = &node
		return
	}

	p := this.delist
	var pre *ListNode

	for {
		if x > p.Val {
			if p.Next == nil || p.Next.Val > x {
				node.Next = p.Next
				p.Next = &node
				return
			}
		} else {
			node.Next = p
			if pre != nil {
				pre.Next = &node
			} else {
				this.delist = &node
			}
			return
		}
		pre = p
		p = p.Next
	}
}


func (this *MinStack) Pop()  {
	if len(this.elements) == 0 {
		return
	}
	ele := this.elements[len(this.elements) - 1]
	this.elements = this.elements[:len(this.elements) - 1]



	if this.delist == ele {
		this.delist = ele.Next
	} else {
		p := this.delist
		for {
			if p == nil {
				break
			}
			if p.Next == ele {
				p.Next = p.Next.Next
				break
			}
			p = p.Next
		}
	}
}


func (this *MinStack) Top() int {
	return this.elements[len(this.elements) - 1].Val
}


func (this *MinStack) GetMin() int {
	return this.delist.Val
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */