package dao

import "time"

// CreateCard 创建卡密
func (d *dao) CreateCard(key string, date time.Time, userID uint) error {
	return d.db.Create(&Card{Key: key, EndDate: date, UserID: userID}).Error
}

// DeleteCard 删除卡密
func (d *dao) DeleteCard(key string) error {
	return d.db.Delete(&Card{}, "key = ?", key).Error
}

// CheckCard 校验卡密
func (d *dao) CheckCard(key string) (bool, error) {
	var card Card
	err := d.db.Where("key = ?", key).First(&card).Error
	if err != nil {
		return false, err
	}
	if card.EndDate.Before(time.Now()) {
		return false, nil
	}
	return true, nil
}

// UpdateEndDataByName 通过名字修改有效期
func (d *dao) UpdateEndDataByName(key string, date time.Time) error {
	return d.db.Model(&Card{}).Where("key = ?", key).Update("end_date", date).Error
}

// UpdateEndDataByID 通过ID修改有效期
func (d *dao) UpdateEndDataByID(id uint, date time.Time) error {
	return d.db.Model(&Card{}).Where("id = ?", id).Update("end_date", date).Error
}
