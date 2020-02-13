package b

type B interface {
	Get() string
}

type Bimpl struct {
}

func (b *Bimpl) Get() string  {
    return "b bimpl get"
}
