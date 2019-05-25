# stack.go

This is my implementation of a LIFO stack in Go. Who hasn't made one of these? 

Supports the following methods:

* `Push`
    
    Pushes a value unto the stack

* `Pop`

    Pops a value from the top of the stack

* `Peek`

    Gets the top element of the stack without popping it

* `Bottom`

    Gets the bottom element of the stack without popping it

* `Contains`

    Tests whether a value is in the stack

* `PopFirst`

    Pops the first value from the top of the stack that matches the given element

* `PopLast`

    Pops the first value from the bottom of the stack that matches the given element


# Design considerations

I've avoided doing the whole `return nil, err` thing for such a simple data structure. I've opted instead to just return `nil` on all illegal cases.

* Returns `nil` when popping from an empty stack
* Returns `nil` when peeking an empty stack
* `PopFirst` and `PopLast` returns `true` or `false` depending on whether the item was in the stack. 
    Since, you'd need a value that matches the item anyways to be able to search it. It makes no sense having them return the value you just searched for,

# LICENSE

This is licensed under the MIT license