package database

type ColumnType int

const (
	_ ColumnType = iota
	TypeInt32
	TypeInt64
)

type (
	Column struct {
		name    string
		colType ColumnType
	}
	SchemaBuilder interface {
		SetTableName(string) SchemaBuilder
		SetColumns(...Column) SchemaBuilder
		Build() error
	}
)

func NewColumn(name string, colType ColumnType) Column {
	return Column{name, colType}
}
