package boot

import (
	"reflect"
)

type Context struct {
	Factories map[string]func(*Context) any
	Instances map[string]any
}

func NewContext() *Context {
	return &Context{
		Factories: make(map[string]func(*Context) any),
		Instances: make(map[string]any),
	}
}

func typeOf[T any]() string {
	return reflect.TypeOf((*T)(nil)).Elem().String()
}

func Register[T any](c *Context, factory func(*Context) *T) {
	key := typeOf[T]()
	c.Factories[key] = func(ctx *Context) any {
		return factory(ctx)
	}
}

func Get[T any](c *Context) *T {
	key := typeOf[T]()

	if instance, exists := c.Instances[key]; exists {
		return instance.(*T)
	}

	if factory, exists := c.Factories[key]; exists {
		instance := factory(c).(*T)
		c.Instances[key] = instance

		return instance
	}

	return nil
}
