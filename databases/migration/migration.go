package main

import (
	"log"

	"github.com/yuri7772d/reditCloneApi/config"
	"github.com/yuri7772d/reditCloneApi/databases"
	"github.com/yuri7772d/reditCloneApi/entities"
)

func main() {
	conf, _ := config.GetingConfig()
	db := databases.NewPosgresDatabase(conf.Database)
	log.Println("On migration ", conf.Database.Dbname)
	TopicMigration(db)
	TopicLikeMigration(db)

}

func TopicMigration(db databases.Database) {
	db.ConnetionGeting().Migrator().CreateTable(&entities.TopicRecord{})
	db.ConnetionGeting().Migrator().CreateTable(&entities.TopicImageRecord{})
	log.Printf("ok - Migration Topic Succes.")
}
func TopicLikeMigration(db databases.Database) {
	db.ConnetionGeting().Migrator().CreateTable(&entities.TopicLikeCountRecord{})
	db.ConnetionGeting().Migrator().CreateTable(&entities.TopicLikeRecord{})
	log.Printf("ok - Migration TopicLike Succes.")

}
