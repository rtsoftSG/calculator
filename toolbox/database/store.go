package database

type (
	field struct {
		name  string
		value interface{}
	}
	Storage interface {
		Insert(tableName string, fields ...field) error
	}
)

func Field(name string, value interface{}) field {
	return field{name, value}
}
