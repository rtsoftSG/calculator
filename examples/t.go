package examples

import "calculator/toolbox/database"

func createSchema(builder database.SchemaBuilder) {
	err := builder.SetTableName("q").SetColumns(
		database.Column("a", database.TypeInt32),
		database.Column("b", database.TypeInt64),
	).Build()

	if err != nil {
		panic(err)
	}
}

func insertData(storage database.Storage) {
	err := storage.Insert(
		"q",
		database.Field("a", 1),
		database.Field("b", 2),
	)

	if err != nil {
		panic(err)
	}
}
