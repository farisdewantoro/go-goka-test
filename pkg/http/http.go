package http

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/middleware/requestid"
)

type Configurator func(*Bootstrapper)

type Bootstrapper struct {
	*iris.Application
}

func NewApplication() *Bootstrapper {
	return &Bootstrapper{
		Application: iris.New(),
	}
}

// Configure accepts configurations and runs them inside the Bootstraper's context.
func (b *Bootstrapper) Configure(cs ...Configurator) *Bootstrapper {
	for _, c := range cs {
		c(b)
	}
	return b
}

func (b *Bootstrapper) SetDefaultErrorHandlers() *Bootstrapper {
	b.OnAnyErrorCode(func(c iris.Context) {

	})
	return b
}

func (b *Bootstrapper) SetDefaultMiddleware() *Bootstrapper {
	b.Use(recover.New())
	b.UseRouter(requestid.New())
	return b
}

// Listen starts the http server with the specified "addr".
func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	b.Run(iris.Addr(addr), cfgs...)
}
