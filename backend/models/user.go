package models

// User 用户结构体
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// UserResponse 用户响应结构体
type UserResponse struct {
	User User `json:"user"`
}

// UsersResponse 用户列表响应结构体
type UsersResponse struct {
	Users []User `json:"users"`
	Count int    `json:"count"`
}

// CreateUserRequest 创建用户请求结构体
type CreateUserRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required,min=1,max=150"`
}

// UserService 用户服务接口
type UserService interface {
	GetAllUsers() []User
	GetUserByID(id int) (*User, error)
	CreateUser(name string, age int) *User
}

// InMemoryUserService 内存用户服务实现
type InMemoryUserService struct {
	users  []User
	nextID int
}

// NewInMemoryUserService 创建新的内存用户服务
func NewInMemoryUserService() *InMemoryUserService {
	return &InMemoryUserService{
		users: []User{
			{ID: 1, Name: "张三", Age: 25},
			{ID: 2, Name: "李四", Age: 30},
			{ID: 3, Name: "王五", Age: 28},
		},
		nextID: 4,
	}
}

// GetAllUsers 获取所有用户
func (s *InMemoryUserService) GetAllUsers() []User {
	return s.users
}

// GetUserByID 根据ID获取用户
func (s *InMemoryUserService) GetUserByID(id int) (*User, error) {
	for _, user := range s.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, ErrUserNotFound
}

// CreateUser 创建新用户
func (s *InMemoryUserService) CreateUser(name string, age int) *User {
	user := User{
		ID:   s.nextID,
		Name: name,
		Age:  age,
	}
	s.users = append(s.users, user)
	s.nextID++
	return &user
}
