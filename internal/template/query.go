package template

// DefaultQuery default query
const DefaultQuery = `
var (
	Q =new(Query)
	{{range $name,$d :=.Data -}}
	{{$d.ModelStructName}} *{{$d.QueryStructName}}
	{{end -}}
)

func SetDefault(db *gorm.DB, opts ...sqlgen.DOOption) {
	*Q = *Use(db,opts...)
	{{range $name,$d :=.Data -}}
	{{$d.ModelStructName}} = &Q.{{$d.ModelStructName}}
	{{end -}}
}

`

// QueryMethod query method template
const QueryMethod = `
func Use(db *gorm.DB, opts ...sqlgen.DOOption) *Query {
	return &Query{
		db: db,
		{{range $name,$d :=.Data -}}
		{{$d.InternalModelStructName}}: new{{$d.ModelStructName}}(db,opts...),
		{{end -}}
	}
}

type Query struct{
	db *gorm.DB

	{{range $name,$d :=.Data -}}
	{{$d.InternalModelStructName}} {{$d.QueryStructName}}
	{{end}}
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db: db,
		{{range $name,$d :=.Data -}}
		{{$d.InternalModelStructName}}: q.{{$d.InternalModelStructName}}.clone(db),
		{{end}}
	}
}

func (q *Query) ReadDB() *Query {
	return q.clone(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.clone(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db: db,
		{{range $name,$d :=.Data -}}
		{{$d.InternalModelStructName}}: q.{{$d.InternalModelStructName}}.replaceDB(db),
		{{end}}
	}
}

type queryInstance struct {
	{{range $name,$d :=.Data -}}
	{{$d.ModelStructName}} {{$d.QueryStructName}}
	{{end}}
}

func (q *Query) Instance(ctx context.Context) *queryInstance {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
    if ok {
		new := q.clone(tx)
		return &queryInstance{
			{{range $name,$d :=.Data -}}
			{{$d.ModelStructName}}: new.{{$d.InternalModelStructName}},
			{{end}}
		}
    }
	return &queryInstance{
		{{range $name,$d :=.Data -}}
		{{$d.ModelStructName}}: q.{{$d.InternalModelStructName}},
		{{end}}
	}
}

type queryDo struct{ 
	{{range $name,$d :=.Data -}}
	{{$d.ModelStructName}} {{$d.ReturnObject}}
	{{end}}
}

func (q *queryInstance) WithContext(ctx context.Context) *queryDo  {
	return &queryDo{
		{{range $name,$d :=.Data -}}
		{{$d.ModelStructName}}: q.{{$d.ModelStructName}}.WithContext(ctx),
		{{end}}
	}
}

// 用来承载事务的上下文
type contextTxKey struct{}

func (q *Query) ExecTx(ctx context.Context, f func(ctx context.Context) error) error {
	return q.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx := context.WithValue(ctx, contextTxKey{}, tx)
		return f(ctx) 
	})
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}

`

// QueryMethodTest query method test template
const QueryMethodTest = `

const dbName = "gen_test.db"

var db *gorm.DB
var once sync.Once

func init() {
	InitializeDB()
	db.AutoMigrate(&_another{})
}

func InitializeDB() {
	once.Do(func() {
		var err error
		db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
		if err != nil {
			panic(fmt.Errorf("open sqlite %q fail: %w", dbName, err))
		}
	})
}

func assert(t *testing.T, methodName string, res, exp interface{}) {
	if !reflect.DeepEqual(res, exp) {
		t.Errorf("%v() gotResult = %v, want %v", methodName, res, exp)
	}
}

type _another struct {
	ID uint64 ` + "`" + `gorm:"primaryKey"` + "`" + `
}

func (*_another) TableName() string { return "another_for_unit_test" }

func Test_Available(t *testing.T) {
	if !Use(db).Available() {
		t.Errorf("query.Available() == false")
	}
}

func Test_WithContext(t *testing.T) {
	query := Use(db)
	if !query.Available() {
		t.Errorf("query Use(db) fail: query.Available() == false")
	}

	type Content string
	var key, value Content = "gen_tag", "unit_test"
	qCtx := query.WithContext(context.WithValue(context.Background(), key, value))

	for _, ctx := range []context.Context{
		{{range $name,$d :=.Data -}}
		qCtx.{{$d.ModelStructName}}.UnderlyingDB().Statement.Context,
		{{end}}
	} {
		if v := ctx.Value(key); v != value {
			t.Errorf("get value from context fail, expect %q, got %q", value, v)
		}
	}
}

func Test_Transaction(t *testing.T) {
	query := Use(db)
	if !query.Available() {
		t.Errorf("query Use(db) fail: query.Available() == false")
	}

	err := query.Transaction(func(tx *Query) error { return nil })
	if err != nil {
		t.Errorf("query.Transaction execute fail: %s", err)
	}

	tx := query.Begin()

	err = tx.SavePoint("point")
	if err != nil {
		t.Errorf("query tx SavePoint fail: %s", err)
	}
	err = tx.RollbackTo("point")
	if err != nil {
		t.Errorf("query tx RollbackTo fail: %s", err)
	}
	err = tx.Commit()
	if err != nil {
		t.Errorf("query tx Commit fail: %s", err)
	}

	err = query.Begin().Rollback()
	if err != nil {
		t.Errorf("query tx Rollback fail: %s", err)
	}
}
`
