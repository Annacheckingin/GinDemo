package mysql

import (
	"errors"
)

type Model interface {
	TableName() string
	IdValue() any
}

type PageContext struct {
	Page     int
	PageSize int
}

func PageFind[T Model](ctx PageContext) ([]T, error) {
	var retval []T
	offset := (ctx.Page - 1) * ctx.PageSize
	limit := ctx.PageSize
	result := Db.Offset(offset).Limit(limit).Find(&retval)
	if result.Error != nil {
		return []T{}, result.Error
	}
	return retval, nil
}

func FindByLimit[T Model](maxLimit int) ([]T, error) {
	var retval []T
	limit := maxLimit
	if limit <= 0 {
		limit = 1000
	}
	result := Db.Limit(limit).Find(&retval)
	if result.Error != nil {
		return []T{}, result.Error
	}
	return retval, nil
}

func Create[T Model](model *T) error {
	result := Db.Create(&model)
	return result.Error
}

func FindById[T Model, ID any](model T, id ID) (T, error) {
	var retval = model
	result := Db.First(&retval, id)
	return retval, result.Error
}

// UpdateById :跟新model，前提是这个model必须带有id
func UpdateById[T Model](model T) error {
	if model.IdValue() == nil {
		return errors.New("id can't be nil")
	}
	result := Db.Save(&model)
	return result.Error
}

func DeleteById[T Model, ID any](model T, id ID) error {
	intId, ok := any(id).(int)
	if !ok {
		return errors.New("id can't be empty")
	}
	if intId == 0 {
		return errors.New("id can't be empty")
	}
	result := Db.Delete(&model, intId)
	return result.Error
}

func Total[T Model](model T) (int64, error) {
	var count int64
	if err := Db.Model(&model).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
