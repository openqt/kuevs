package kuevs

import (
	"github.com/golang/glog"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Database struct {
	Engine *gorm.DB
}

func InitDatabase(name string) *Database {
	db, err := gorm.Open(sqlite.Open(name), &gorm.Config{})
	Check(err)

	err = db.AutoMigrate(&Event{})
	Check(err)

	return &Database{
		Engine: db,
	}
}

func (d *Database) Save(evt *Event) {
	data := Event{UID: evt.UID, UpdatedAt: evt.UpdatedAt}
	result := d.Engine.Find(&data)
	Check(result.Error)
	if result.RowsAffected > 0 {
		glog.V(5).Infof("Find: %s (%s => %s)\n", evt.UID, data.UpdatedAt, evt.UpdatedAt)
	}

	Check(d.Engine.Clauses(clause.OnConflict{DoNothing: true}).Create(evt).Error)
}
