package xorm

import (
	"database/sql"
	"io"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
)

// Session is interface for xorm.Session
type Session interface {
	Init()
	Close()
	Sql(string, ...interface{}) *xorm.Session
	Where(string, ...interface{}) *xorm.Session
	And(string, ...interface{}) *xorm.Session
	Or(string, ...interface{}) *xorm.Session

	Id(interface{}) *xorm.Session
	Table(interface{}) *xorm.Session
	In(string, ...interface{}) *xorm.Session
	Select(string) *xorm.Session
	Cols(...string) *xorm.Session
	MustCols(...string) *xorm.Session
	AllCols() *xorm.Session
	Distinct(...string) *xorm.Session
	Omit(...string) *xorm.Session
	Nullable(...string) *xorm.Session
	NoAutoTime() *xorm.Session

	Limit(int, ...int) *xorm.Session
	OrderBy(string) *xorm.Session
	Desc(...string) *xorm.Session
	Asc(...string) *xorm.Session
	Join(string, interface{}, string) *xorm.Session
	GroupBy(string) *xorm.Session
	Having(string) *xorm.Session

	Begin() error
	Rollback() error
	Commit() error
	Exec(string, ...interface{}) (sql.Result, error)

	CreateTable(interface{}) error
	CreateIndexes(interface{}) error
	CreateUniques(interface{}) error
	DropIndexes(interface{}) error
	DropTable(interface{}) error

	Rows(interface{}) (*xorm.Rows, error)
	Iterate(interface{}, xorm.IterFunc) error
	Get(interface{}) (bool, error)
	Count(interface{}) (int64, error)
	Find(interface{}, ...interface{}) error

	IsTableExist(interface{}) (bool, error)
	IsTableEmpty(interface{}) (bool, error)

	Query(string, ...interface{}) ([]map[string][]byte, error)
	Insert(...interface{}) (int64, error)
	InsertMulti(interface{}) (int64, error)
	InsertOne(interface{}) (int64, error)
	Update(interface{}, ...interface{}) (int64, error)
	Delete(interface{}) (int64, error)

	Sync2(...interface{}) error
}

// Engine

type Engine interface {
	SetDisableGlobalCache(bool)
	DriverName() string
	DataSourceName() string
	SetMapper(core.IMapper)
	SetTableMapper(core.IMapper)
	SetColumnMapper(core.IMapper)
	SupportInsertMany() bool
	QuoteStr() string
	Quote(string) string
	AutoIncrStr() string
	SetMaxOpenConns(int)
	SetMaxConns(int)
	SetMaxIdleConns(int)
	NoCache() *xorm.Session
	NoCascade() *xorm.Session

	NewSession() *xorm.Session
	Close() error
	Ping() error

	Sql(string, ...interface{}) *xorm.Session
	NoAutoTime() *xorm.Session
	DumpAllToFile(string) error
	DumpAll(io.Writer) error

	Cascade(...bool) *xorm.Session
	Where(string, ...interface{}) *xorm.Session
	Id(interface{}) *xorm.Session

	Before(func(interface{})) *xorm.Session
	After(func(interface{})) *xorm.Session
	Charset(string) *xorm.Session
	StoreEngine(string) *xorm.Session

	Distinct(...string) *xorm.Session
	Select(string) *xorm.Session
	Cols(...string) *xorm.Session
	AllCols() *xorm.Session
	MustCols(...string) *xorm.Session
	UseBool(...string) *xorm.Session
	Omit(...string) *xorm.Session
	Nullable(...string) *xorm.Session
	In(string, ...interface{}) *xorm.Session
	Incr(string, ...interface{}) *xorm.Session
	Decr(string, ...interface{}) *xorm.Session

	Table(interface{}) *xorm.Session
	Limit(int, ...int) *xorm.Session
	Desc(...string) *xorm.Session
	Asc(...string) *xorm.Session
	OrderBy(string) *xorm.Session
	Join(string, interface{}, string) *xorm.Session
	GroupBy(string) *xorm.Session
	Having(string) *xorm.Session

	GobRegister(interface{}) *xorm.Engine

	IsTableEmpty(interface{}) (bool, error)
	IsTableExist(interface{}) (bool, error)
	CreateIndexes(interface{}) error
	CreateUniques(interface{}) error
	ClearCacheBean(interface{}, string) error
	ClearCache(...interface{}) error
	Sync(...interface{}) error
	Sync2(...interface{}) error
	CreateTables(...interface{}) error
	DropTables(...interface{}) error

	Exec(string, ...interface{}) (sql.Result, error)
	Query(string, ...interface{}) ([]map[string][]byte, error)
	Insert(...interface{}) (int64, error)
	Update(interface{}, ...interface{}) (int64, error)
	Delete(interface{}) (int64, error)
	Get(interface{}) (bool, error)
	Find(interface{}, ...interface{}) error
	Iterate(interface{}, xorm.IterFunc) error
	Rows(interface{}) (*xorm.Rows, error)
	Count(interface{}) (int64, error)

	ImportFile(string) ([]sql.Result, error)
	Import(io.Reader) ([]sql.Result, error)

	TZTime(time.Time) time.Time
	NowTime(string) interface{}
	NowTime2(string) (interface{}, time.Time)
	FormatTime(string, time.Time) interface{}
	Unscoped() *xorm.Session
}
