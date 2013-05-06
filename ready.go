// Ready.Go is a web development framework for golang.

package ready

type App struct {
}

type Context struct {
	app      *App
	request  *Request
	response *Response
}

type Handler interface {
	Handle(ctx *Context)
}

func (this *App) Get(route string, handler *Handler) {
}
