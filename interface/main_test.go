package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	stack := NewEmptyStack()
	t.Run("when item is pushed", func(t *testing.T) {
		expectedStack := NewStack([]Item{
			1,
			2,
			3,
		})
		t.Run("stack should be updated", func(t *testing.T) {
			stack.Push(1)
			stack.Push(2)
			stack.Push(3)
			assert.Equal(t, stack, expectedStack)
		})
	})
}

func TestPop(t *testing.T) {
	stack := NewEmptyStack()
	t.Run("when stack is empty", func(t *testing.T) {
		expectedStack := NewEmptyStack()
		t.Run("it should return nil response", func(t *testing.T) {
			// Pop 1
			poppedItem := stack.Pop()
			assert.Equal(t, poppedItem, nil)
			assert.Equal(t, stack, expectedStack)
		})
	})

	t.Run("when stack is not empty", func(t *testing.T) {
		expectedStack := NewStack([]Item{
			1,
		})
		t.Run("it should return last inserted item and update the stack", func(t *testing.T) {
			// Push two items
			stack.Push(1)
			stack.Push(2)

			// Pop 1
			poppedItem := stack.Pop()
			assert.Equal(t, poppedItem, 2)
			assert.Equal(t, stack, expectedStack)
		})
	})
}

func TestIsEmpty(t *testing.T) {
	stack := NewEmptyStack()
	t.Run("when stack is not empty", func(t *testing.T) {
		t.Run("it should return false", func(t *testing.T) {
			// Push one item
			stack.Push(1)
			assert.Equal(t, stack.IsEmpty(), false)
			// Pop one item
			stack.Pop()
			assert.Equal(t, stack.IsEmpty(), true)
		})
	})

	t.Run("when stack is empty", func(t *testing.T) {
		t.Run("it should return true", func(t *testing.T) {
			// Pop one item
			stack.Pop()
			assert.Equal(t, stack.IsEmpty(), true)
		})
	})
}

func TestTop(t *testing.T) {
	t.Run("when stack is not empty", func(t *testing.T) {
		stack := NewEmptyStack()
		expectedStack := NewStack([]Item{
			1,
			2,
		})
		t.Run("it should return last inserted item and no change in stack", func(t *testing.T) {
			// Push two items
			stack.Push(1)
			stack.Push(2)

			item := stack.Top()
			assert.Equal(t, item, 2)
			assert.Equal(t, stack, expectedStack)
		})
	})

	t.Run("when stack is empty", func(t *testing.T) {
		stack := NewEmptyStack()
		t.Run("it should return nil", func(t *testing.T) {
			item := stack.Top()
			assert.Equal(t, item, nil)
		})
	})
}

func TestSize(t *testing.T) {
	stack := NewEmptyStack()
	t.Run("when stack is provided", func(t *testing.T) {
		t.Run("it should return number of items in stack", func(t *testing.T) {
			// Push one item
			stack.Push(1)
			assert.Equal(t, stack.Size(), 1)
			// Pop one item
			stack.Pop()
			assert.Equal(t, stack.Size(), 0)
		})
	})
}

func TestClear(t *testing.T) {
	stack := NewEmptyStack()
	t.Run("when stack is provided", func(t *testing.T) {
		expectedStack := NewEmptyStack()
		t.Run("it should clear all items in stack", func(t *testing.T) {
			// Push one item
			stack.Push(1)
			stack.Push(2)
			stack.Clear()
			assert.Equal(t, stack, expectedStack)
		})
	})
}
