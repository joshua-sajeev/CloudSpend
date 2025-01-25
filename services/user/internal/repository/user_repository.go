package repository

type UserRepository interface {
	// Create a new user
	Create(user *models.User) error

	// Find user by specific fields
	FindByID(id uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	FindByUsername(username string) (*models.User, error)

	// Update user information
	Update(user *models.User) error

	// Delete user
	Delete(id uint) error

	// List and search users
	List(page, pageSize int) ([]models.User, error)
	Search(query string) ([]models.User, error)
}
