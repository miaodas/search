package models

import (
	"search/hiendy"
)

func GetResources(offset int, limit int, list *[]hiendy.Resource) error {
	limitNum := 20
	if limit > 0 {
		limitNum = limit
	}
	conn, err := hiendy.GetConnection()
	if err != nil {
		return err
	}
	err = conn.Offset(offset).Limit(limitNum).Order("id Asc").Find(list).Error
	if err != nil {
		return err
	}
	return err
}
