package database

type ColumnType int

const (
	_ ColumnType = iota
	TypeInt32
	TypeInt64
)

type (
	column struct {
		name    string
		colType ColumnType
	}
	SchemaBuilder interface {
		SetTableName(string) SchemaBuilder
		SetColumns(...column) SchemaBuilder
		Build() error
	}
)

func Column(name string, colType ColumnType) column {
	return column{name, colType}
}
