package stack

import "testing"

func expect(t *testing.T, got interface{}, expected interface{}) {
	t.Helper()
	if got != expected {
		t.Errorf("expected %s, got %s", expected, got)
	}
}

func TestPush(t *testing.T) {
	stack := NewStack(10)

	data := []string{"alpha", "beta", "gamma"}
	for _, v := range data {
		stack.Push(v)
	}

	for i, expected := range data {
		got := stack.Values[i]
		expect(t, got, expected)
	}
}

func TestPop(t *testing.T) {
	stack := NewStack(10)

	data := []string{"alpha", "beta", "gamma", "delta"}
	expected := []string{"delta", "gamma", "beta", "alpha"}

	for _, v := range data {
		stack.Push(v)
	}

	for _, ex := range expected {
		expect(t, stack.Pop(), ex)
	}

	// popping an empty stack should give nil
	expect(t, stack.Pop(), nil)
}

func TestPeek(t *testing.T) {
	stack := NewStack(10)

	// peeking an empty stack should give nil
	expect(t, stack.Peek(), nil)

	data := []string{"alpha", "beta", "gamma", "delta"}

	for _, v := range data {
		stack.Push(v)
		expect(t, stack.Peek(), v)
	}

	// make sure all the data is still there
	for i, v := range stack.Values {
		expect(t, v, data[i])
	}
}

func TestBottom(t *testing.T) {
	stack := NewStack(10)

	// should be nil since it is empty
	expect(t, stack.Bottom(), nil)

	data := []string{"alpha", "beta", "gamma", "delta"}

	for _, v := range data {
		stack.Push(v)
	}

	expect(t, stack.Bottom(), "alpha")

	// make sure all the data is still there
	for i, v := range stack.Values {
		expect(t, v, data[i])
	}
}

func TestPopFirst(t *testing.T) {
	stack := NewStack(10)

	// should not find anything because it is empty
	expect(t, stack.PopFirst(5), false)

	for i := 0; i < 10; i++ {
		stack.Push(i)
	}

	expect(t, stack.PopFirst(5), true)

	expected := []int{0, 1, 2, 3, 4, 6, 7, 8, 9}

	for i, got := range stack.Values {
		expect(t, got, expected[i])
	}

	// should return false because it is not in the stack
	expect(t, stack.PopFirst(5), false)

	expect(t, stack.PopFirst(9), true)
	expect(t, stack.PopFirst(0), true)
	expect(t, stack.PopFirst(3), true)

	expected = []int{1, 2, 4, 6, 7, 8}

	for i, got := range stack.Values {
		expect(t, got, expected[i])
	}
}

func TestPopFirstWithDuplicates(t *testing.T) {
	stack := NewStack(10)

	data := []string{"alpha", "alpha", "beta", "gamma", "alpha", "delta", "delta"}
	for _, v := range data {
		stack.Push(v)
	}

	var expected []string
	matchExpected := func() {
		for i, got := range stack.Values {
			expect(t, got, expected[i])
		}
	}

	stack.PopFirst("delta")
	expected = []string{"alpha", "alpha", "beta", "gamma", "alpha", "delta"}
	matchExpected()

	stack.PopFirst("alpha")
	expected = []string{"alpha", "alpha", "beta", "gamma", "delta"}
	matchExpected()

	stack.PopFirst("alpha")
	expected = []string{"alpha", "beta", "gamma", "delta"}
	matchExpected()

}

func TestPopLast(t *testing.T) {
	stack := NewStack(10)

	expect(t, stack.PopLast("something"), false)

	data := []string{"alpha", "alpha", "beta", "gamma", "alpha", "delta", "delta", "beta"}
	for _, v := range data {
		stack.Push(v)
	}

	var expected []string
	matchExpected := func() {
		for i, got := range stack.Values {
			expect(t, got, expected[i])
		}
	}

	expect(t, stack.PopLast("mishu"), false)

	expect(t, stack.PopLast("alpha"), true)
	expected = []string{"alpha", "beta", "gamma", "alpha", "delta", "delta", "beta"}
	matchExpected()

	expect(t, stack.PopLast("beta"), true)
	expected = []string{"alpha", "gamma", "alpha", "delta", "delta", "beta"}
	matchExpected()

	expect(t, stack.PopLast("alpha"), true)
	expected = []string{"gamma", "alpha", "delta", "delta", "beta"}
	matchExpected()
}

func TestContains(t *testing.T) {
	stack := NewStack(10)

	expect(t, stack.Contains("something"), false)

	data := []string{"alpha", "beta", "gamma", "delta"}

	for _, v := range data {
		stack.Push(v)
	}

	expect(t, stack.Contains("alpha"), true)
	expect(t, stack.Contains("omega"), false)
	expect(t, stack.Contains("beta"), true)
	expect(t, stack.Contains("theta"), false)
	expect(t, stack.Contains("gamma"), true)
	expect(t, stack.Contains("omicron"), false)
	expect(t, stack.Contains("delta"), true)
	expect(t, stack.Contains("sigma"), false)

}

func TestLenCap(t *testing.T) {
	stack := NewStack(10)

	expect(t, stack.Len(), 0)
	expect(t, stack.Cap(), 10)

	for i := 0; i < 314; i++ {
		stack.Push(i)
	}

	expect(t, stack.Len(), 314)
	expect(t, stack.Cap() >= 314, true)
}
