package scaffold

import "gorm.io/gorm"

// Msg .
type Msg struct {
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
	Err  *string `json:"err"`
}

// IRemove .
func IRemove[T any](db *gorm.DB, ob *T, ids ...interface{}) (*Msg, error) {
	m := &Msg{}
	var err error
	if len(ids) == 0 {
		err = db.Delete(ob).Error

	} else {
		err = db.Delete(ob, ids).Error
	}
	if err != nil {
		errString := err.Error()
		m.Code = 404
		m.Err = &errString
		m.Msg = errString
		return m, err
	}
	m.Code = 400
	m.Msg = "Deleted successfully!"
	return m, nil
}

// IList .
func IList[T any](db *gorm.DB, obs []*T) ([]*T, error) {
	err := db.Find(&obs).Error
	return obs, err
}

// IListByOrder .
func IListByOrder[T any](db *gorm.DB, obs []*T) ([]*T, error) {
	err := db.Order("\"order\"").Find(&obs).Error
	return obs, err
}

// IListWithChildren .
func IListWithChildren[T any](db *gorm.DB, obs []*T, level int) ([]*T, error) {
	if level < 1 {
		level = 1
	}
	preload := "Children"
	for l := 1; l < level; l++ {
		preload += ".Children"
	}
	err := db.Order("\"order\"").Preload(preload).Where("id =1").Find(&obs).Error
	return obs, err
}

// IFindByID .
func IFindByID[T any](db *gorm.DB, ob *T, id interface{}) (*T, error) {
	err := db.Where("id = ?", id).First(&ob).Error
	return ob, err
}

// IFindByField: fn field name, fv field value
func IFindByField[T any](db *gorm.DB, ob []*T, fn string, fv interface{}) ([]*T, error) {
	err := db.Where(fn+" = ?", fv).First(&ob).Error
	return ob, err
}

// ICrerate .
func ICrerate[T any](db *gorm.DB, ob *T) (*T, error) {
	err := db.Omit("id").Create(&ob).Error
	return ob, err
}

// IUpdate .
func IUpdate[T any](db *gorm.DB, ob *T) (*T, error) {
	err := db.Save(&ob).Error
	return ob, err
}
