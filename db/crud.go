package db

type Model interface {
	TableName() string
}

type PageContext struct {
	Page     int
	PageSize int
}

func PageFind[T Model](ctx PageContext) ([]T, error) {
	var retval []T
	offset := (ctx.Page - 1) * ctx.PageSize
	limit := ctx.PageSize
	result := Db.Find(&retval).Offset(offset).Limit(limit)
	if result.Error != nil {
		return []T{}, result.Error
	}
	return retval, nil
}

func FindByLimit[T Model](maxLimit int) []T {
	var retval []T
	limit := maxLimit
	if limit <= 0 {
		limit = 1000
	}
	result := Db.Find(&retval).Limit(limit)
	if result.Error != nil {
		return []T{}
	}
	return retval
}

func Create[T Model](model T) error {
	result := Db.Create(&model)
	return result.Error
}

func FindById[T Model, ID any](id ID) (T, error) {
	var model T
	result := Db.First(&model, id)
	return model, result.Error
}

// 跟新model，前提是这个model必须带有id
func UpdateById[T Model](model T) error {
	result := Db.Save(&model)
	return result.Error
}

func DeleteById[T Model, ID any](id ID) error {
	var model T
	result := Db.Delete(&model, id)
	return result.Error
}
