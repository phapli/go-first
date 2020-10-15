package models

import (
	"github.com/jinzhu/gorm"
)

type News struct {
	gorm.Model
	Name      string
	Thumbnail string
	Content   string
	Tags      string
}

func (n *News) SaveNews(db *gorm.DB) (*News, error) {
	var err error
	err = db.Debug().Create(&n).Error
	if err != nil {
		return &News{}, err
	}
	return n, err
}

func (n *News) FindAllNews(db *gorm.DB) (*[]News, error) {
	var err error
	var news []News
	err = db.Debug().Model(News{}).Limit(100).Find(&news).Error
	if err != nil {
		return &[]News{}, err
	}
	return &news, err
}

func (n *News) FindByID(db *gorm.DB, id uint) (*News, error) {
	var err error
	var news News
	err = db.Debug().Model(News{}).Where("id = ?", id).Take(&news).Error
	if err != nil {
		return &News{}, err
	}
	return &news, err
}

func (n *News) UpdateANews(db *gorm.DB) (*News, error) {

	var err error
	db = db.Debug().Model(&News{}).Where("id = ?", n.ID).Take(&News{}).UpdateColumns(
		map[string]interface{}{
			"content":   n.Content,
			"name":      n.Name,
			"tags":      n.Tags,
			"thumbnail": n.Thumbnail,
		},
	)
	if db.Error != nil {
		return &News{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&News{}).Where("id = ?", n.ID).Take(&n).Error
	if err != nil {
		return &News{}, err
	}
	return n, nil
}

func (n *News) DeleteANews(db *gorm.DB, id uint) (int64, error) {

	db = db.Debug().Model(&News{}).Where("id = ?", id).Take(&News{}).Delete(&News{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
