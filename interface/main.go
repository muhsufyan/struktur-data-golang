package main

import "sync"

/*
source https://www.mohitkhare.com/blog/stack-in-golang/
We don’t want data inconsistency which might result in failure of struct
so we use mutex
*/
// Item interface to store any data type in stack
type Item interface{}

// Stack struct which contains a list of Items
type Stack struct {
	items []Item
	// lock (for concurrency purposes)
	mutex sync.Mutex
}

// NewEmptyStack() returns a new instance of Stack with zero elements
func NewEmptyStack() *Stack {
	return &Stack{
		items: nil,
	}
}

// NewStack() returns a new instance of Stack with list of specified elements
func NewStack(items []Item) *Stack {
	return &Stack{
		items: items,
	}
}

// Push() adds new item to top of existing/empty stack
// If stack is full, this returns overflow error. Here we don’t handle overflow, you can add custom check and handle the error.
func (stack *Stack) Push(item Item) {
	// lock the mutex so that no other changes are made and consistency is maintained (in concurrent environment)
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	stack.items = append(stack.items, item)
}

// Pop() removes most recent item(top) from stack
func (stack *Stack) Pop() Item {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	if len(stack.items) == 0 {
		return nil
	}

	lastItem := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]

	return lastItem
}

// Top() returns the last inserted item in stack without removing it.
func (stack *Stack) Top() Item {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	if len(stack.items) == 0 {
		return nil
	}

	return stack.items[len(stack.items)-1]
}

// IsEmpty() returns whether the stack is empty or not (boolean result)
func (stack *Stack) IsEmpty() bool {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	return len(stack.items) == 0
}

// Size() returns the number of items currently in the stack
func (stack *Stack) Size() int {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	return len(stack.items)
}

// Clear() removes all items from the stack
// removes all elements from stack
func (stack *Stack) Clear() {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	stack.items = nil
}
