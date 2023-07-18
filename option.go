package gooption

type Option[A any] struct {
	a *A
}

func Some[A any](a A) Option[A] {
	return Option[A]{a: &a}
}

func None[A any]() Option[A] {
	return Option[A]{}
}

// This will panic if option is empty
func (o Option[A]) Get() A {
	if o.a == nil {
		panic("no value in Option!")
	}

	return *o.a
}
