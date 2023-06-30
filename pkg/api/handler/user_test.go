package handler

import (
	"bytes"
	"ecommerce/pkg/commonhelp/requests.go"
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/usecase/mockusecase"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserSignup(t *testing.T) {
	ctrl := gomock.NewController(t)
	userUseCase := mockusecase.NewMockUserUseCase(ctrl)
	userHandler := NewUserHandler(userUseCase)

	testData := []struct {
		name             string
		input            requests.Usersign
		buildStub        func(userUseCase mockusecase.MockUserUseCase)
		expectedCode     int
		expectedResponse response.Response
		expectedData     response.UserValue
		expectedError    error
	}{
		{
			name: "successful",
			input: requests.Usersign{
				Name:     "Akshay",
				Email:    "akshay@gmail.com",
				Mobile:   "+917994475799",
				Password: "akshay@123",
			},
			buildStub: func(userUseCase mockusecase.MockUserUseCase) {
				userUseCase.EXPECT().UserSignup(gomock.Any(), requests.Usersign{
					Name:     "Akshay",
					Email:    "akshay@gmail.com",
					Mobile:   "+917994475799",
					Password: "akshay@123",
				}).Times(1).
					Return(response.UserValue{
						ID:        1,
						Name:      "Akshay",
						Email:     "akshay@gmail.com",
						Password:  "akshay@123",
						CreatedAt: time.Now(),
					}, nil)
			},
			expectedCode: 201,
			expectedResponse: response.Response{
				StatusCode: 201,
				Message:    "User signup successful",
				Data: response.UserValue{
					ID:        1,
					Name:      "Akshay",
					Email:     "akshay@gmail.com",
					CreatedAt: time.Now(),
				},
				Errors: nil,
			},
			expectedData: response.UserValue{
				ID:        1,
				Name:      "Akshay",
				Email:     "akshay@gmail.com",
				CreatedAt: time.Now(),
			},
			expectedError: nil,
		},
		{
         name : "duplicate user",
		 input: requests.Usersign{
			Name: "Akshay",
			Email: "akshay@gmail.com",
			Mobile: "+917994475799",
			Password: "akshay@123",
		 },
		 buildStub: func(userUseCase mockusecase.MockUserUseCase) {
			userUseCase.EXPECT().UserSignup(gomock.Any(), requests.Usersign{
				Name:     "Akshay",
				Email:    "akshay@gmail.com",
				Mobile:   "+917994475799",
				Password: "akshay@123",
			}).Times(1).
				Return(response.UserValue{}, errors.New("user is already exist "))
		},
          expectedCode: 400,
		  expectedResponse: response.Response{
			StatusCode: 400,
				Message:    "unable create account",
				Data:       response.UserValue{},
				Errors:     "user already exits",
		  },
		  expectedData:  response.UserValue{},
		  expectedError: errors.New("user already exists"),
		},
	}

	for _, tc := range testData {

		t.Run(tc.name, func(t *testing.T) {

			tc.buildStub(*userUseCase)

			
			server := gin.New()
			server.POST("/signup", userHandler.UserSignup)
			
			jsonData, err := json.Marshal(&tc.input)
			assert.NoError(t, err)
			body := bytes.NewBuffer(jsonData)
			
			mockReq, err := http.NewRequest(http.MethodPost, "/signup", body)
			assert.NoError(t, err)

			responseRec := httptest.NewRecorder()

			server.ServeHTTP(responseRec, mockReq)

			//validate
			assert.Equal(t, tc.expectedCode, responseRec.Code)
		})
	}
}
