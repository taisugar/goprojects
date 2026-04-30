package domain

type Tool struct {
	Name        string
	Description string
	Handler     func(map[string]interface{}) (any, error)
}

type Resource struct {
	URI         string
	Description string
	Handler     func() (any, error)
}

type Prompt struct {
	Name        string
	Description string
	Handler     func(map[string]interface{}) (any, error)
}
