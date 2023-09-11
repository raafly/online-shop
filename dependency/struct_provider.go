/*
kita juga bisa membuat sebuah struct sebagai provider
*/

package dependency

type Foo struct {
}

func NewFoo() *Foo {
	return &Foo{}
}

type Bar struct {
}

func Newbar() *Bar {
	return &Bar{}
}

type FooBar struct {
	*Foo
	*Bar
}