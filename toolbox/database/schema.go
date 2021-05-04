package database

type DataType int

const (
	_ DataType = iota
	TypeBool

	TypeInt8
	TypeInt16
	TypeInt32
	TypeInt64

	TypeUint8
	TypeUint16
	TypeUint32
	TypeUint64

	TypeFloat32
	TypeFloat64

	TypeString

	TypeArrayUint8
	TypeArrayUint16
	TypeArrayUint32
	TypeArrayUint64

	TypeArrayInt8
	TypeArrayInt16
	TypeArrayInt32
	TypeArrayInt64

	TypeArrayFloat32
	TypeArrayFloat64

	TypeNullBool

	TypeNullInt8
	TypeNullInt16
	TypeNullInt32
	TypeNullInt64
	
	TypeNullUint8
	TypeNullUint16
	TypeNullUint32
	TypeNullUint64
	
	TypeNullFloat32
	TypeNullFloat64

	TypeNullString
)

type (
	Column struct {
		name     string
		dataType DataType
	}
	SchemaBuilder interface {
		SetTableName(string) SchemaBuilder
		SetColumns(...Column) SchemaBuilder
		Build() error
	}
)

func NewColumn(name string, colType DataType) Column {
	return Column{name, colType}
}

func (c *Column) DataType() DataType {
	return c.dataType
}

func (c *Column) Name() string {
	return c.name
}
