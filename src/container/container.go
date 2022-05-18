package container

type Container struct {
	bindings map[string]interface{}
}

func NewContainer() *Container {
	return &Container{}
}
