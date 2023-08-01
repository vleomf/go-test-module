package pk1

type Impl2 struct {
	OtherValue string
}

func (x *Impl2) SayFoo() string {
	return x.OtherValue
}
