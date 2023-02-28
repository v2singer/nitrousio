package dao

import (
	"gorm.io/gorm"
)

const tableName = "project"

// ProjectORM project
type ProjectORM struct {
}

type CodeI struct {
}

// TableName table name
func (ProjectORM) TableName() string {
	return tableName
}

type ProjectDao struct {
	db *gorm.DB
}

type UpdateOption func(*gorm.DB) *gorm.DB

func withIsSync(isSync bool) UpdateOption {
	val := 0
	if isSync {
		val = 1
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Set("is_sync", val)
	}
}

func withCodeI(codei CodeI) UpdateOption {
	return func(db *gorm.DB) *gorm.DB {
		return db
	}
}

func (g *ProjectDao) UpdateOptionByID(id string, options ...UpdateOption) error {
	data := g.db
	for _, option := range options {
		data = option(data)
	}
	return nil

}

func (p *ProjectDao) Create(project *ProjectORM) error {
	return nil
}

func (p *ProjectDao) UpdateByID(id string) error {
	return nil
}

func (p *ProjectDao) DeleteByID(id string) error {
	return nil
}

func (p *ProjectDao) GetByID(id string) *ProjectORM {
	return &ProjectORM{}
}

func (p *ProjectDao) List(filter map[string]interface{}) []*ProjectORM {
	return []*ProjectORM{}
}
