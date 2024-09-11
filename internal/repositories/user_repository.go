package repositories

import (
	"fiber-app-paafff/internal/domain/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// Fungsi NewUserRepository menerima parameter db yang merupakan pointer ke gorm.DB
// Buat objek UserRepository baru dengan db sebagai salah satu propertinya
// Kembalikan pointer ke objek UserRepository yang baru dibuat
func NewUserRepository(db *gorm.DB) *UserRepository {
	// Membuat dan mengembalikan objek UserRepository baru
	return &UserRepository{db: db}
}

// CreateUser creates a new user in the database
func (r *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// GetUser retrieves a user by ID from the database
func (r *UserRepository) GetUser(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates an existing user in the database
func (r *UserRepository) UpdateUser(user *models.User) (*models.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteUser deletes a user by ID from the database
func (r *UserRepository) DeleteUser(id uint) error {
	if err := r.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

// ListUsers retrieves all users from the database
func (r *UserRepository) ListUsers() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
