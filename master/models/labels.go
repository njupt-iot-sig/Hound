package models

import (
	"github.com/jinzhu/gorm"
)

type Label struct {
	gorm.Model
	Name string
	Node string
}

func CreateLabelWithNode(name string, node string) error {
	err, ok := CheckNode(node)
	if err != nil {
		return err
	}
	if ok {
		L := Label{
			Name: name,
			Node: node,
		}
		if err := db.Create(&L).Error; err != nil {
			return err
		}
	}
	return nil
}

func ListAllLabels() ([]*Label, error) {
	var labelList []*Label
	if err := db.Find(&labelList).Error; err != nil {
		return nil, err
	} else {
		return labelList, nil
	}
}

//func CheckLabel(name string) (error, bool) {
//	if err := db.Where("name = ?", name).First(&Label{}).Error; err != nil {
//		return err, false
//	}
//	return nil, true
//}

func OnlyListLabels() ([]*Label, error) {
	var list []*Label
	if err := db.Table("label").Select("distinct label").Scan(&list).Error; err != nil {
		return list, err
	} else {
		return list, nil
	}
}

func GetLabelNodes(name string) (err error, node string) {
	label := Label{}
	if err := db.Where("name = ?", name).First(&label).Error; err != nil {
		return err, label.Node
	}
	return nil, label.Node
}
