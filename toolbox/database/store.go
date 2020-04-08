package database

type (
	Field struct {
		name  string
		value interface{}
	}
	Storage interface {
		Insert(tableName string, fields ...Field) error
	}
)

func NewField(name string, value interface{}) Field {
	return Field{name, value}
}
