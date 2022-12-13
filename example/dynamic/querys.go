package dynamic

import gen "github.com/go-leo/sqlgen"

type Querier interface {
	// FindByNameAndAge query data by name and age and return it as map
	//
	// where("name=@name AND age=@age")
	FindByNameAndAge(name string, age int) (gen.M, error)
}
