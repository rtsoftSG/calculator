package toplevel

import "calculator/toolbox/database"

type ResultDbMapper interface {
	PrepareSchema(builder database.SchemaBuilder)
	Save(data map[string]interface{})
}

type MapperFactoryFunc func(storage database.Storage) (ResultDbMapper, error)
