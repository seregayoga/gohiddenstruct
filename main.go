package gohiddenstruct

type Exported interface {
	GetField() string
}

type hidden struct {
	field string
}

func NewExported(field string) Exported {
	return hidden{field: field}
}

func (h hidden) GetField() string {
	return h.field
}
