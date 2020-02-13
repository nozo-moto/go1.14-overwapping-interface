package a

type A interface {
	Get() string
}

type Aimpl struct {
}

func (a *Aimpl) Get() string  {
    return "a a impl"
}
