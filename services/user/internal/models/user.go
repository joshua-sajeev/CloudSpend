package models
type User struct {
    ID        uint
    Username  string
    Email     string
    Password  string
    Role      UserRole
    CreatedAt time.Time
    UpdatedAt time.Time
}

type UserRole string

const (
    RoleAdmin    UserRole = "admin"
    RoleUser     UserRole = "user"
    RoleManager  UserRole = "manager"
)

// Validation methods
func (u *User) Validate() error
func (u *User) BeforeCreate() error
func (u *User) BeforeUpdate() errortype User struct {
    ID        uint
    Username  string
    Email     string
    Password  string
    Role      UserRole
    CreatedAt time.Time
    UpdatedAt time.Time
}

type UserRole string

const (
    RoleAdmin    UserRole = "admin"
    RoleUser     UserRole = "user"
    RoleManager  UserRole = "manager"
)

// Validation methods
func (u *User) Validate() error
func (u *User) BeforeCreate() error
func (u *User) BeforeUpdate() error
