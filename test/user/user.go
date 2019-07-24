//go:generate mockgen -destination mock_user/user.go -source=$GOFILE
package user

type User struct {
	Id   int
	Name string
	Age  int
}

type UserService interface {
	GetUserById(id int) (User, error)
	UpdateUser(id int, u User) error
	UpdateUsers(u User, ids ...int) error
}
