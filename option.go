package gooption

import "encoding/json"

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

func (o Option[A]) GetOrElse(a A) A {
	if o.a == nil {
		return a
	}
	return *o.a
}

func (o Option[A]) IsDefined() bool {
	return o.a != nil
}

func (o Option[A]) IsEmpty() bool {
	return !o.IsDefined()
}

func (o Option[A]) MarshalJSON() ([]byte, error) {
	if o.IsDefined() {
		return json.Marshal(o.a)
	} else {
		return []byte("null"), nil
	}
}

func (o *Option[A]) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.a = nil
		return nil
	}

	var value A

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	o.a = &value
	return nil
}

func Map[A any, B any](o Option[A], f func(a A) B) Option[B] {
	if o.a != nil {
		return Some(f(*o.a))
	}
	return None[B]()
}

func FlatMap[A any, B any](o Option[A], f func(a A) Option[B]) Option[B] {
	if o.a != nil {
		return f(*o.a)
	}
	return None[B]()
}
