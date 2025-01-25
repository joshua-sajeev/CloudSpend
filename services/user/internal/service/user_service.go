package service

type UserService interface {
	// Authentication Methods
	Register(username, email, password string) (*models.User, error)
	Login(email, password string) (string, error) // Returns JWT token
	Logout(token string) error

	// Profile Management
	GetProfile(userID uint) (*models.User, error)
	UpdateProfile(userID uint, updates map[string]interface{}) error

	// Authorization Methods
	ValidateToken(token string) (*models.User, error)
	ChangePassword(userID uint, oldPassword, newPassword string) error

	// Admin Functions
	ListUsers(page, pageSize int) ([]models.User, error)
	UpdateUserRole(adminID, targetUserID uint, newRole models.UserRole) error
}
