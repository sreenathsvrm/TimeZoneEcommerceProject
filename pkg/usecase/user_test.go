package usecase

import (
	"context"
	"ecommerce/pkg/commonhelp/requests.go"
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/domain"
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
			defer ctrl.Finish()
			userRepo := mockrepo.NewMockUserRepository(ctrl)
			userUseCase := NewUserUseCase(userRepo)
			tt.buildStub(userRepo, tt.input)

			user, err := userUseCase.UserSignup(context.TODO(), tt.input)

			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, user, tt.expectedOutput)
		})
	}

}

func TestLoginWithEmail(t *testing.T) {
	//NewController from gomock package returns a new controller for testing
	ctrl := gomock.NewController(t)

	// NewMockUserRepository creates a new mockRepo instance
	userRepo := mockrepo.NewMockUserRepository(ctrl)
	userUseCase := NewUserUseCase(userRepo)

	
	

	// testData is a slice of struct which holds multiple test cases
	testData := []struct {
		name           string
		input          requests.Login
		buildStub      func(userRepo *mockrepo.MockUserRepository)
		isExpectingOutput bool
		expectedError  error
	}{
		{
			name: "get details from database",
			input: requests.Login{
				Email:    "sreenathsvrm@gmail.com",
				Password: "sree@123",
			},
			buildStub: func(userRepo *mockrepo.MockUserRepository) {
				userRepo.EXPECT().UserLogin(gomock.Any(), "sreenathsvrm@gmail.com").Times(1).
					Return(domain.Users{}, errors.New("no user found"))
			},
			isExpectingOutput: false,
			expectedError:  errors.New("no user found"),
		},

		{
			name: "blocked user",
			input: requests.Login{
				Email:    "sreenathsvrm@gmail.com",
				Password: "sree@123",
			},
			buildStub: func(userRepo *mockrepo.MockUserRepository) {
				userRepo.EXPECT().UserLogin(gomock.Any(), "sreenathsvrm@gmail.com").Times(1).
					Return(domain.Users{
						ID: 1,
						Email: "sreenathsvrm@gmail.com",
						Password: "sree@123",
						IsBlocked: true,
					}, errors.New("user is blocked"))
			},
			isExpectingOutput: false,
			expectedError:  errors.New("user is blocked"),
		},
   
        {
			name: "blocked user",
			input: requests.Login{
				Email:    "sreenathsvrm@gmail.com",
				Password: "sree@123",
			},
			buildStub: func(userRepo *mockrepo.MockUserRepository) {
				userRepo.EXPECT().UserLogin(gomock.Any(), "sreenathsvrm@gmail.com").Times(1).
					Return(domain.Users{
						ID: 1,
						Email: "sreenathsvrm@gmail.com",
						Password: "sree@123",
						IsBlocked: true,
					}, errors.New("user is blocked"))
			},
			isExpectingOutput: false,
			expectedError:  errors.New("user is blocked"),
		},
      
		{
			name: "successfull login to watch shop",
			input: requests.Login{
				Email:    "sreenathsvrm@gmail.com",
				Password: "sree@123",
			},
			buildStub: func(userRepo *mockrepo.MockUserRepository) {
				hashedPassword,err:=bcrypt.GenerateFromPassword([]byte("sree@123"),10)
				if err!=nil{
					t.Fatalf("")
				}
				userRepo.EXPECT().UserLogin(gomock.Any(), "sreenathsvrm@gmail.com").Times(1).
					Return(domain.Users{
						ID: 1,
						Email: "sreenathsvrm@gmail.com",
						Password: string(hashedPassword),
						IsBlocked: false,
					}, nil)
			},
			isExpectingOutput: true,
			expectedError:  nil,
		},
	
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.buildStub(userRepo)
			tokenString, actualErr := userUseCase.UserLogin(context.TODO(), tt.input)

			if tt.expectedError == nil {
				assert.Nil(t, actualErr)
			} else {
				assert.Equal(t, tt.expectedError, actualErr)
			}

			if tt.isExpectingOutput{
				assert.NotEmpty(t,tokenString)
			}else{
				assert.Empty(t,tokenString)
			}

		})
	}
}
