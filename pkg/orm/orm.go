package orm

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Config struct {
	DSN          string
	MaxIdleConns int64
	MaxOpenConns int64
	MaxLifetime  int64
}

type DB struct {
	*gorm.DB
}

type ormLog struct {
	LogLevel logger.LogLevel
}

func (l *ormLog) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l *ormLog) Info(ctx context.Context, format string, v ...interface{}) {
	if l.LogLevel < logger.Info {
		return
	}

	logx.WithContext(ctx).Infof(format, v...)
}

func (l *ormLog) Warn(ctx context.Context, format string, v ...interface{}) {
	if l.LogLevel < logger.Warn {
		return
	}

	logx.WithContext(ctx).Infof(format, v...)
}

func (l *ormLog) Error(ctx context.Context, format string, v ...interface{}) {
	if l.LogLevel < logger.Error {
		return
	}

	logx.WithContext(ctx).Errorf(format, v...)
}

func (l *ormLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	logx.WithContext(ctx).WithDuration(elapsed).Infof("[%.3fms] [rows:%v] %s", float64(elapsed.Nanoseconds())/1e6, rows, sql)
}

func NewMySQL(conf *Config) (*DB, error) {
	db, err := gorm.Open(mysql.Open(conf.DSN), &gorm.Config{
		Logger: &ormLog{},
	})

	if err != nil {
		return nil, err
	}

	sdb, err := db.DB()
	if err != nil {
		return nil, err
	}

	sdb.SetMaxIdleConns(int(conf.MaxIdleConns))
	sdb.SetMaxOpenConns(int(conf.MaxOpenConns))
	sdb.SetConnMaxLifetime(time.Second * time.Duration(conf.MaxLifetime))

	err = db.Use(NewCustomPlugin())
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

func MustNewMySQL(conf *Config) *DB {
	db, err := NewMySQL(conf)
	if err != nil {
		panic(err)
	}
	return db
}

// Paginate 分页
func Paginate(pageNum, pageSize int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNum == 0 {
			pageNum = 1
		}

		if pageSize == 0 {
			pageSize = 10
		}

		offset := (pageNum - 1) * pageSize

		return db.Offset(int(offset)).Limit(int(pageSize))
	}
}
