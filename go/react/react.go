package react

type callback func(int)

type cell struct {
	store     int
	callbacks map[*func(int)]callback
}

type canceller struct {
	f func()
}

type reactor struct {
}

func (c canceller) Cancel() {
	c.f()
}

func (c *cell) SetValue(v int) {
	oldValue := c.Value()

	c.store = v

	if oldValue != c.Value() {
		for _, clb := range c.callbacks {
			clb(c.Value())
		}
	}
}

func (c *cell) Value() int {
	return c.store
}

func (c *cell) AddCallback(clb func(int)) Canceler {
	c.callbacks[&clb] = clb

	return canceller{f: func() { delete(c.callbacks, &clb) }}
}

func (r *reactor) CreateInput(val int) InputCell {
	return &cell{
		store:     val,
		callbacks: make(map[*func(int)]callback),
	}
}

func (r *reactor) CreateCompute1(c Cell, f func(int) int) ComputeCell {
	input := c.(*cell)

	compute := &cell{
		store:     f(c.Value()),
		callbacks: make(map[*func(int)]callback),
	}

	input.AddCallback(func(newVal int) {
		newComputed := f(newVal)
		if compute.store != newComputed {
			compute.store = newComputed
			for _, clb := range compute.callbacks {
				clb(compute.Value())
			}
		}
	})

	return compute
}

func (r *reactor) CreateCompute2(c1, c2 Cell, f func(int, int) int) ComputeCell {
	input1 := c1.(*cell)
	input2 := c2.(*cell)

	computed := f(c1.Value(), c2.Value())

	computeCell := &cell{
		store:     computed,
		callbacks: make(map[*func(int)]callback),
	}

	callback := func(newVal int) {
		newComputed := f(c1.Value(), c2.Value())
		if computed != newComputed {
			computeCell.store = newComputed
			for _, clb := range computeCell.callbacks {
				clb(newComputed)
			}
		}
	}

	input1.AddCallback(callback)
	input2.AddCallback(callback)

	return computeCell
}

func New() *reactor {
	return &reactor{}
}
