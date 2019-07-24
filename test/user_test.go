package test

import (
	"runtime"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/golang/mock/test/user"
	"github.com/golang/mock/test/user/mock_user"
)

type MockHelper struct {
	*testing.T
}

func (h *MockHelper) Helper() {
	pc, file, line, _ := runtime.Caller(1)
	h.Logf("---> %v %v %v \n", runtime.FuncForPC(pc).Name(), file, line)
}

func MockUserSevice(t *testing.T) (*mock_user.MockUserService, *gomock.Controller) {
	helper := &MockHelper{t}
	ctrl := gomock.NewController(helper)
	return mock_user.NewMockUserService(ctrl), ctrl
}

func TestGetUser(t *testing.T) {
	service, ctrl := MockUserSevice(t)
	defer ctrl.Finish()

	call := service.EXPECT().GetUserById(1).Return(user.User{Name: "a"}, nil).Times(1)
	t.Logf("%v\n", call.String())

	u, _ := service.GetUserById(1)
	t.Logf("%v\n", u)
}
