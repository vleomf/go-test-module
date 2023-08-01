package pk1

type Impl1 struct {
	Value string
}

func (x Impl1) SayFoo() string {
	return x.Value
}
