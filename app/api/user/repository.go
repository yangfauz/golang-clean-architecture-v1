package user

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]User, error)
	FindByID(ID int) (User, error)
	Save(user User) (User, error)
	Update(user User) (User, error)
	Delete(user User) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]User, error) {
	var users []User

	err := r.db.Find(&users).Error

	if err != nil {
		return users, err
	}

	return users, nil
}

func (r *repository) FindByID(ID int) (User, error) {
	var user User

	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Save(user User) (User, error) {

	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) Delete(user User) (bool, error) {
	err := r.db.Delete(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
