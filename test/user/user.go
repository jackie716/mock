//go:generate mockgen -destination mock_user/user.go -source=$GOFILE
package user

import "net"

type User struct {
	Id   int
	Name string
	Age  int
}

type UserService interface {
	GetUserById(id int) User
	UpdateUser(id int, u User) error
	UpdateUsers(u User, ids ...int) error

	Find(a int, b map[string]int, c net.Conn, d []net.Conn, e [3]int)
}
