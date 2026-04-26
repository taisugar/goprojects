package storage

type Strategy[T any] interface {
	Save(data T) error
	Load(data *T) error
}
