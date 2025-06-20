// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"kratos-demo/app/user/internal/data/gen/model"
)

func newTestCompany(db *gorm.DB, opts ...gen.DOOption) testCompany {
	_testCompany := testCompany{}

	_testCompany.testCompanyDo.UseDB(db, opts...)
	_testCompany.testCompanyDo.UseModel(&model.TestCompany{})

	tableName := _testCompany.testCompanyDo.TableName()
	_testCompany.ALL = field.NewAsterisk(tableName)
	_testCompany.ID = field.NewInt32(tableName, "id")
	_testCompany.Name = field.NewString(tableName, "name")
	_testCompany.Phone = field.NewString(tableName, "phone")
	_testCompany.Password = field.NewString(tableName, "password")
	_testCompany.AppID = field.NewString(tableName, "app_id")
	_testCompany.AppSecret = field.NewString(tableName, "app_secret")
	_testCompany.CreateTime = field.NewTime(tableName, "create_time")
	_testCompany.UpdateTime = field.NewTime(tableName, "update_time")

	_testCompany.fillFieldMap()

	return _testCompany
}

// testCompany 公司
type testCompany struct {
	testCompanyDo

	ALL        field.Asterisk
	ID         field.Int32
	Name       field.String
	Phone      field.String
	Password   field.String
	AppID      field.String
	AppSecret  field.String
	CreateTime field.Time // 创建时间
	UpdateTime field.Time // 更新时间

	fieldMap map[string]field.Expr
}

func (t testCompany) Table(newTableName string) *testCompany {
	t.testCompanyDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t testCompany) As(alias string) *testCompany {
	t.testCompanyDo.DO = *(t.testCompanyDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *testCompany) updateTableName(table string) *testCompany {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt32(table, "id")
	t.Name = field.NewString(table, "name")
	t.Phone = field.NewString(table, "phone")
	t.Password = field.NewString(table, "password")
	t.AppID = field.NewString(table, "app_id")
	t.AppSecret = field.NewString(table, "app_secret")
	t.CreateTime = field.NewTime(table, "create_time")
	t.UpdateTime = field.NewTime(table, "update_time")

	t.fillFieldMap()

	return t
}

func (t *testCompany) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *testCompany) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 8)
	t.fieldMap["id"] = t.ID
	t.fieldMap["name"] = t.Name
	t.fieldMap["phone"] = t.Phone
	t.fieldMap["password"] = t.Password
	t.fieldMap["app_id"] = t.AppID
	t.fieldMap["app_secret"] = t.AppSecret
	t.fieldMap["create_time"] = t.CreateTime
	t.fieldMap["update_time"] = t.UpdateTime
}

func (t testCompany) clone(db *gorm.DB) testCompany {
	t.testCompanyDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t testCompany) replaceDB(db *gorm.DB) testCompany {
	t.testCompanyDo.ReplaceDB(db)
	return t
}

type testCompanyDo struct{ gen.DO }

type ITestCompanyDo interface {
	gen.SubQuery
	Debug() ITestCompanyDo
	WithContext(ctx context.Context) ITestCompanyDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITestCompanyDo
	WriteDB() ITestCompanyDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITestCompanyDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITestCompanyDo
	Not(conds ...gen.Condition) ITestCompanyDo
	Or(conds ...gen.Condition) ITestCompanyDo
	Select(conds ...field.Expr) ITestCompanyDo
	Where(conds ...gen.Condition) ITestCompanyDo
	Order(conds ...field.Expr) ITestCompanyDo
	Distinct(cols ...field.Expr) ITestCompanyDo
	Omit(cols ...field.Expr) ITestCompanyDo
	Join(table schema.Tabler, on ...field.Expr) ITestCompanyDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITestCompanyDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITestCompanyDo
	Group(cols ...field.Expr) ITestCompanyDo
	Having(conds ...gen.Condition) ITestCompanyDo
	Limit(limit int) ITestCompanyDo
	Offset(offset int) ITestCompanyDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITestCompanyDo
	Unscoped() ITestCompanyDo
	Create(values ...*model.TestCompany) error
	CreateInBatches(values []*model.TestCompany, batchSize int) error
	Save(values ...*model.TestCompany) error
	First() (*model.TestCompany, error)
	Take() (*model.TestCompany, error)
	Last() (*model.TestCompany, error)
	Find() ([]*model.TestCompany, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TestCompany, err error)
	FindInBatches(result *[]*model.TestCompany, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TestCompany) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITestCompanyDo
	Assign(attrs ...field.AssignExpr) ITestCompanyDo
	Joins(fields ...field.RelationField) ITestCompanyDo
	Preload(fields ...field.RelationField) ITestCompanyDo
	FirstOrInit() (*model.TestCompany, error)
	FirstOrCreate() (*model.TestCompany, error)
	FindByPage(offset int, limit int) (result []*model.TestCompany, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Rows() (*sql.Rows, error)
	Row() *sql.Row
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITestCompanyDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t testCompanyDo) Debug() ITestCompanyDo {
	return t.withDO(t.DO.Debug())
}

func (t testCompanyDo) WithContext(ctx context.Context) ITestCompanyDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t testCompanyDo) ReadDB() ITestCompanyDo {
	return t.Clauses(dbresolver.Read)
}

func (t testCompanyDo) WriteDB() ITestCompanyDo {
	return t.Clauses(dbresolver.Write)
}

func (t testCompanyDo) Session(config *gorm.Session) ITestCompanyDo {
	return t.withDO(t.DO.Session(config))
}

func (t testCompanyDo) Clauses(conds ...clause.Expression) ITestCompanyDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t testCompanyDo) Returning(value interface{}, columns ...string) ITestCompanyDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t testCompanyDo) Not(conds ...gen.Condition) ITestCompanyDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t testCompanyDo) Or(conds ...gen.Condition) ITestCompanyDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t testCompanyDo) Select(conds ...field.Expr) ITestCompanyDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t testCompanyDo) Where(conds ...gen.Condition) ITestCompanyDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t testCompanyDo) Order(conds ...field.Expr) ITestCompanyDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t testCompanyDo) Distinct(cols ...field.Expr) ITestCompanyDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t testCompanyDo) Omit(cols ...field.Expr) ITestCompanyDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t testCompanyDo) Join(table schema.Tabler, on ...field.Expr) ITestCompanyDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t testCompanyDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITestCompanyDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t testCompanyDo) RightJoin(table schema.Tabler, on ...field.Expr) ITestCompanyDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t testCompanyDo) Group(cols ...field.Expr) ITestCompanyDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t testCompanyDo) Having(conds ...gen.Condition) ITestCompanyDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t testCompanyDo) Limit(limit int) ITestCompanyDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t testCompanyDo) Offset(offset int) ITestCompanyDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t testCompanyDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITestCompanyDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t testCompanyDo) Unscoped() ITestCompanyDo {
	return t.withDO(t.DO.Unscoped())
}

func (t testCompanyDo) Create(values ...*model.TestCompany) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t testCompanyDo) CreateInBatches(values []*model.TestCompany, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t testCompanyDo) Save(values ...*model.TestCompany) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t testCompanyDo) First() (*model.TestCompany, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TestCompany), nil
	}
}

func (t testCompanyDo) Take() (*model.TestCompany, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TestCompany), nil
	}
}

func (t testCompanyDo) Last() (*model.TestCompany, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TestCompany), nil
	}
}

func (t testCompanyDo) Find() ([]*model.TestCompany, error) {
	result, err := t.DO.Find()
	return result.([]*model.TestCompany), err
}

func (t testCompanyDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TestCompany, err error) {
	buf := make([]*model.TestCompany, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t testCompanyDo) FindInBatches(result *[]*model.TestCompany, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t testCompanyDo) Attrs(attrs ...field.AssignExpr) ITestCompanyDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t testCompanyDo) Assign(attrs ...field.AssignExpr) ITestCompanyDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t testCompanyDo) Joins(fields ...field.RelationField) ITestCompanyDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t testCompanyDo) Preload(fields ...field.RelationField) ITestCompanyDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t testCompanyDo) FirstOrInit() (*model.TestCompany, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TestCompany), nil
	}
}

func (t testCompanyDo) FirstOrCreate() (*model.TestCompany, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TestCompany), nil
	}
}

func (t testCompanyDo) FindByPage(offset int, limit int) (result []*model.TestCompany, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t testCompanyDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t testCompanyDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t testCompanyDo) Delete(models ...*model.TestCompany) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *testCompanyDo) withDO(do gen.Dao) *testCompanyDo {
	t.DO = *do.(*gen.DO)
	return t
}
