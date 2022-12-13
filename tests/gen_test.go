package tests_test

import (
	"context"
	"sync"

	"github.com/go-leo/sqlgen/tests/.expect/dal_test/query"
)

var useOnce sync.Once
var ctx = context.Background()

func CRUDInit() {
	query.Use(DB)
	query.SetDefault(DB)
}
