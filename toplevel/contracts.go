package toplevel

import (
	"github.com/rtsoftSG/plugin/toolbox/cache"
	"github.com/rtsoftSG/plugin/toolbox/database"
	"time"
)

//CalculationResult result of single calculation.
type CalculationResult interface {
	//Type return measurement type.
	Type() string
	//Time return time.Time associated with calculation result.
	Time() time.Time
	//Sources return list of sources(actor names).
	Sources() []string
	//CalculatorId return calculator instance unique id (it will be used as ActorName in measurement).
	CalculatorId() string
	//Fields return calculation data.
	Fields() interface{}
}

// Common interface for algorithms.
type Calculator interface {
	Calculate() ([]CalculationResult, error)
}

// CalculatorFactoryFunc creates calculator instance. This function must be realized by every plugin that provides calculator functionality.
//
// NewCalculator function with such signature has to be implemented in order to lookup calculator.
type CalculatorFactoryFunc func(id string, cfg string, cache cache.Storage, qff database.QueryFactoryFunc) (Calculator, error)

type ResultDbMapper interface {
	PrepareSchema(builder database.SchemaBuilder) error
	SubscribedTo() string

	DataIntoInsertCommands(data map[string]interface{}) []*database.InsertCommand
}

// MappersFactoryFunc must return mappers that will be registered for handle incoming data.
//
// Function name must be GetMappers.
type MappersFactoryFunc func() []ResultDbMapper
