package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pkg/errors"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("postgres", "host=localhost user=heaven dbname=breads sslmode=disable")
	if err != nil {
		fmt.Println("Error:", err)
		panic("failed to connect database")

	}

	err = DB.DB().Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		panic("failed to ping database")
	}

	// Migrate the schema
	DB.AutoMigrate(&Bread{})
}

type Bread struct {
	ID        string `gorm:"column:breadID"`
	Name      string
	CreatedAt string `gorm:"column:customCreatedAt"`
}

func SaveBread(bread *Bread) {
	result := DB.Create(bread)
	if result.Error != nil {
		fmt.Println("Error inserting bread:", result.Error)
	} else if result.RowsAffected == 0 {
		fmt.Println("No rows inserted")
	}
}
func GetAllBreads() ([]*Bread, error) {
	var breads []*Bread
	err := DB.Find(&breads).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to get all breads")
	}
	return breads, nil
}

func GetBreadByID(id string) (*Bread, error) {
	var bread Bread
	err := DB.Where("\"breadID\"= ?", id).First(&bread).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to get bread by ID")
	}
	return &bread, nil
}
