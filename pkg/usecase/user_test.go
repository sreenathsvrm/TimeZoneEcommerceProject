package usecase

import (
	"context"
	"ecommerce/pkg/commonhelp/requests.go"
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/repository/mockrepo"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"golang.org/x/crypto/bcrypt"
)

type eqCreateUserParamsMatcher struct {
	arg      requests.Usersign
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(requests.Usersign)
	if !ok {
		return false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(arg.Password), []byte(e.password)); err != nil {
		return false
	}
	e.arg.Password = arg.Password
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}
func EqCreateUserParams(arg requests.Usersign, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}

func TestUserSignup(t *testing.T) {

	constTime := time.Now()

	testData := []struct {
		name           string
		input          requests.Usersign
		buildStub      func(userRepo *mockrepo.MockUserRepository, user requests.Usersign)
		expectedOutput response.UserValue
		expectedError  error
	}{
		{
			name: "FailedToSaveUserOnDatabase",
			input: requests.Usersign{
				Name:     "sreenath",
				Email:    "sreenath@gmail.com",
				Mobile:   "7994475799",
				Password: "sreenath@123",
			},
			buildStub: func(userRepo *mockrepo.MockUserRepository, user requests.Usersign) {

				userRepo.EXPECT().UserSignup(gomock.Any(), EqCreateUserParams(user, user.Password)).Times(1).
					Return(response.UserValue{}, errors.New("error on database"))
			},
			expectedOutput: response.UserValue{},
			expectedError:  errors.New("error on database"),
		},
		{
			name: "SuccessSignup",
			input: requests.Usersign{
				Name:     "sreenath",
				Email:    "sreenath@gmail.com",
				Mobile:   "7994475799",
				Password: "sreenath@123",
			},
			buildStub: func(userRepo *mockrepo.MockUserRepository, user requests.Usersign) {
				userRepo.EXPECT().UserSignup(gomock.Any(), EqCreateUserParams(user, user.Password)).Times(1).
					Return(response.UserValue{
						ID:        1,
						Name:      "sreenath",
						Email:     "sree@gmail.com",
						Password:  "hashed password",
						CreatedAt: constTime,
					}, nil)
			},
			expectedOutput: response.UserValue{
				ID:        1,
				Name:      "sreenath",
				Email:     "sree@gmail.com",
				Password:  "hashed password",
				CreatedAt: constTime,
			},
			expectedError: nil,
		},
	}
	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			ddefer ctrl.Finish()
			userRepo := mockrepo.NewMockUserRepository(ctrl)
			userUseCase := NewUserUseCase(userRepo)
			tt.buildStub(userRepo, tt.input)

			user, err := userUseCase.UserSignup(context.TODO(), tt.input)

			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, user, tt.expectedOutput)
		})
	}

}
