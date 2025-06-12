package databases

import (
	"fmt"
	"sync"

	"github.com/yuri7772d/reditCloneApi/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	*gorm.DB
}

var (
	postgresDatabaseInstance *postgresDatabase
	once                     sync.Once
)

func NewPosgresDatabase(conf *config.Database) Database {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s search_path=%s",
			conf.Host,
			conf.Port,
			conf.User,
			conf.Password,
			conf.Dbname,
			conf.Sslmode,
			conf.Schema,
		)
		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		postgresDatabaseInstance = &postgresDatabase{conn}
	})
	return postgresDatabaseInstance
}

func (db *postgresDatabase) ConnetionGeting() *gorm.DB {
	return db.DB
}
