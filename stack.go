package stack

// Stack is a basic LIFO stack that can hold any value
type Stack struct {
	Values []interface{}
}

// NewStack creates a new stack and returns a pointer to it
func NewStack(capacity int) *Stack {
	return &Stack{
		make([]interface{}, 0, capacity),
	}
}

// Push adds a value to the top of the stack.
func (s *Stack) Push(val interface{}) {
	s.Values = append(s.Values, val)
}

// Pop pops a value off the top of the stack. Returns `nil` if the stack is empty
func (s *Stack) Pop() interface{} {
	if len(s.Values) == 0 {
		return nil
	}
	n := len(s.Values) - 1
	result := s.Values[n]
	s.Values[n] = nil
	s.Values = s.Values[:n]
	return result
}

// Peek returns the top value of the stack, but does not pop it
func (s Stack) Peek() interface{} {
	if len(s.Values) == 0 {
		return nil
	}
	return s.Values[len(s.Values)-1]
}

// Bottom returns the first pushed value of the stack, but does not pop it
func (s Stack) Bottom() interface{} {
	if len(s.Values) == 0 {
		return nil
	}
	return s.Values[0]
}

// Contains tests if a given value is present in the stack using the `==` operator
func (s Stack) Contains(val interface{}) bool {
	for _, v := range s.Values {
		if v == val {
			return true
		}
	}
	return false
}

// Len returns the number of items in the stack
func (s Stack) Len() int {
	return len(s.Values)
}

// Cap returns the capacity of the underlying slice used by the stack
func (s Stack) Cap() int {
	return cap(s.Values)
}

// PopFirst will search through the stack from top to bottom and remove the first element it can find
// that matches the given `find` parameter, shifting all elements above it down.
// The matching is done with the `==` operator.
// Returns `true` if the value was found, and `false` if it was not found.
func (s *Stack) PopFirst(find interface{}) bool {
	for i := len(s.Values) - 1; i >= 0; i-- {
		if s.Values[i] == find {
			cutout(&s.Values, i)
			return true
		}
	}
	return false
}

// PopLast is similar to `PopFirst`, except it starts searching from the bottom up
func (s *Stack) PopLast(find interface{}) bool {
	for i, existing := range s.Values {
		if existing == find {
			cutout(&s.Values, i)
			return true
		}
	}
	return false
}

func cutout(slice *[]interface{}, index int) {
	maxIndex := len(*slice) - 1
	// copy all the values from index i + 1 and paste them one index left
	// this will 'erase' the current pointed value and replace it.
	// NOTE: copy(destination <- source)
	copy((*slice)[index:], (*slice)[index+1:])
	// the last value of the slice should be nil'ed out
	(*slice)[maxIndex] = nil
	// resize s.Values to be one less
	(*slice) = (*slice)[:maxIndex]
}
