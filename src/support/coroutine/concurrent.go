package coroutine

// concurrent is a helper for concurrent programming.
type concurrent struct {
	ch chan struct{}
}

// NewConcurrent returns a new concurrent helper.
func NewConcurrent(limit int) *concurrent {
	return &concurrent{
		ch: make(chan struct{}, limit),
	}
}

// Create Wait blocks until the limit is reached.
func (c *concurrent) Create(fn func()) {
	c.ch <- struct{}{}

	go func() {
		fn()

		<-c.ch
	}()
}

func (c *concurrent) IsEmpty() bool {
	return len(c.ch) == 0
}

func (c *concurrent) IsFull() bool {
	return len(c.ch) == cap(c.ch)
}

func (c *concurrent) Length() int {
	return len(c.ch)
}

func (c *concurrent) Capacity() int {
	return cap(c.ch)
}
