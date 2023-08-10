package mysql

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"mall.com/config/global"
	"mall.com/store"
	"sync"
	"time"
)

var (
	mysqlFactory store.Factory
	once         sync.Once
)

type Options struct {
	Url      string
	Username string
	Password string
}

// New create a new gorm db instance with the given options.
func New(opts *Options) (*gorm.DB, error) {
	dsn := fmt.Sprintf(`%s:%s@%s`,
		opts.Username,
		opts.Password,
		opts.Url,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(10)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(100)

	return db, nil
}

// GetMySQLFactoryOr create mysql factory with the given config.
func GetMySQLFactoryOr() (store.Factory, error) {
	//if mysqlFactory == nil {
	//	return nil, fmt.Errorf("failed to get mysql store fatory")
	//}

	//var err error
	//var dbIns *gorm.DB
	once.Do(func() {
		//	options := &Options{
		//		Url:                  global.Config.Mysql.Url,
		//		Username:             global.Config.Mysql.Username,
		//		Password:              global.Config.Mysql.Password,
		//	}
		//	dbIns, err = New(options)

		//mysqlFactory = &datastore{dbIns}
		mysqlFactory = &datastore{global.Db}

	},
	)

	if mysqlFactory == nil {
		return nil, fmt.Errorf("failed to get mysql store fatory, mysqlFactory: %+v", mysqlFactory)
	}

	return mysqlFactory, nil
}

type datastore struct {
	db *gorm.DB

	// can include two database instance if needed
	// docker *grom.DB
	// db *gorm.DB
}

func (ds *datastore) Categorys() store.CategoryStore {
	return newCategorys(ds)
}

func (ds *datastore) Products() store.ProductStore {
	return newProducts(ds)
}

func (ds *datastore) Orders() store.OrderStore {
	return newOrders(ds)
}

func (ds *datastore) Close() error {
	db, err := ds.db.DB()
	if err != nil {
		return errors.Wrap(err, "get gorm db instance failed")
	}

	return db.Close()
}
