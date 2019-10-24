package initialize

import (
	"fmt"
	stdlog "log"
	"os"
	"runtime"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/liov/hoper/go/v2/utils/log"
	"github.com/liov/hoper/go/v2/utils/reflect3"
)

const (
	MYSQL = "mysql"
	POSTGRES = "postgres"
	SQLite = "sqlite3"
)

func (i *Init) P2DB() *gorm.DB {
	conf :=DatabaseConfig{}
	if exist := reflect3.GetFieldValue(i.conf,&conf);!exist{
		return nil
	}
	var url string
	if conf.Type == MYSQL {
		url = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			conf.User, conf.Password, conf.Host,
			conf.Port, conf.Database, conf.Charset)
	} else if conf.Type == POSTGRES {
		url = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
			conf.Host, conf.User, conf.Database, conf.Password)
	} else if conf.Type == SQLite {
		url = "/data/db/sqlite/"+ conf.Database+".db"
		if runtime.GOOS == "windows" {
			url=".."+url
		}
	}
	db, err := gorm.Open(conf.Type, url)
	log.Error(conf.Password)
	if err != nil {
		log.Error(err)
		os.Exit(10)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return conf.TablePrefix + defaultTableName
	}

	if i.Env != PRODUCT {
		//b不set输出空白
		db.SetLogger(stdlog.New(os.Stdout, "", 3))
		db.LogMode(true)
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(conf.MaxIdleConns)
	db.DB().SetMaxOpenConns(conf.MaxOpenConns)
	return db
}

