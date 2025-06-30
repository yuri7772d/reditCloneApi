package main

import (
	"log"

	"github.com/yuri7772d/reditCloneApi/config"
	"github.com/yuri7772d/reditCloneApi/databases"
	"github.com/yuri7772d/reditCloneApi/entities"
)

func main() {
	conf, _ := config.Get()
	db := databases.NewPosgresDatabase(conf.Database)
	log.Println("On migration ", conf.Database.Dbname)

	ProfileMigration(db)
	TopicMigration(db)
	LikeTopicMigration(db)
	RepliedMigration(db)
	RepliedLikeMigration(db)

}

func TopicMigration(db databases.Database) {
	db.Connet().Migrator().CreateTable(&entities.Topic{})
	db.Connet().Migrator().CreateTable(&entities.TopicImage{})
	log.Printf("ok - Topic succes.")
}
func LikeTopicMigration(db databases.Database) {
	db.Connet().Migrator().CreateTable(&entities.TopicLike{})
	db.Connet().Migrator().CreateTable(&entities.TopicLikeTracking{})
	log.Printf("ok - LikeTopic succes.")

}
func ProfileMigration(db databases.Database) {
	db.Connet().Migrator().CreateTable(&entities.Profile{})

	log.Printf("ok - Profile succes.")

}
func RepliedMigration(db databases.Database) {
	db.Connet().Migrator().CreateTable(&entities.Replied{})
	db.Connet().Migrator().CreateTable(&entities.RepliedImage{})
	log.Printf("ok - Replied succes.")
}
func RepliedLikeMigration(db databases.Database) {
	db.Connet().Migrator().CreateTable(&entities.RepliedLike{})
	db.Connet().Migrator().CreateTable(&entities.RepliedLikeTracking{})
	log.Printf("ok - RepliedLike succes.")
}
