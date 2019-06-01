package pkg

import (
	"github.com/jinzhu/gorm"
)

type MysqlRepository struct {
	db *gorm.DB
}

func NewMysqlRepository(db *gorm.DB) *MysqlRepository {
	return &MysqlRepository{db}
}

func (repo *MysqlRepository) HasTable(value interface{}) bool {
	return repo.db.HasTable(value)
}

func (repo *MysqlRepository) AutoMigrate(value interface{}) (interface{}, error) {
	if dbc := repo.db.AutoMigrate(value); dbc.Error != nil {
		return nil, dbc.Error
	}
	return nil, nil
}

func (repo *MysqlRepository) FindOne(id string, out interface{}) (interface{}, error) {
	if dbc := repo.db.Where("id = ?", id).First(out); dbc.Error != nil {
		return nil, dbc.Error
	}
	return nil, nil
}

func (repo *MysqlRepository) Find(out interface{}, where ...interface{}) (interface{}, error) {
	if dbc := repo.db.Find(out).Where(where); dbc.Error != nil {
		return nil, dbc.Error
	}
	return nil, nil
}

func (repo *MysqlRepository) Save(value interface{}) {
	repo.db.Create(value)
}
