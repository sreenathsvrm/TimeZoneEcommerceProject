package repository

import (
	"context"
	"ecommerce/pkg/commonhelp/requests.go"
	"ecommerce/pkg/commonhelp/response"
	"ecommerce/pkg/domain"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUserSignup(t *testing.T) {

	tests := []struct {
		name           string
		input          requests.Usersign
		expectedOutput response.UserValue
		buildStub      func(mock sqlmock.Sqlmock)
		expectedErr    error
	}{
		{
			name: "successful create user",

			input: requests.Usersign{
				Name:     "akshay",
				Email:    "akshay@gmail.com",
				Mobile:   "+919562461825",
				Password: "akshay@123",
			},
			expectedOutput: response.UserValue{
				ID:    6,
				Name:  "akshay",
				Email: "akshay@gmail.com",
			},

			buildStub: func(mock sqlmock.Sqlmock) {

				rows := sqlmock.NewRows([]string{"id", "name", "email", "mobile"}).
					AddRow(6, "akshay", "akshay@gmail.com", "+919562461825")

				mock.ExpectQuery("^INSERT INTO users (.+)$").
					WithArgs("akshay", "akshay@gmail.com", "+919562461825", "akshay@123").
					WillReturnRows(rows)
			},
			expectedErr: nil,
		},
		{
			name: "duplicate user",
			input: requests.Usersign{
				Name:     "akshay",
				Email:    "akshay@gmail.com",
				Mobile:   "+919562461825",
				Password: "akshay@123",
			},
			expectedOutput: response.UserValue{},
			buildStub: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("^INSERT INTO users (.+)$").
					WithArgs("akshay", "akshay@gmail.com", "+919562461825", "akshay@123").
					WillReturnError(errors.New("email or phone number alredy used"))
			},
			expectedErr: errors.New("email or phone number alredy used"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//New() method from sqlmock package create sqlmock database connection and a mock to manage expectations.
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			//initialize the db instance with the mock db connection
			gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
			if err != nil {
				t.Fatalf("an error '%s' was not expected when initializing a mock db session", err)
			}

			//create NewUserRepository mock by passing a pointer to gorm.DB
			userRepository := NewUserRepository(gormDB)

			// before we actually execute our function, we need to expect required DB actions
			tt.buildStub(mock)

			//call the actual method
			actualOutput, actualErr := userRepository.UserSignup(context.TODO(), tt.input)
			// validate err is nil if we are not expecting to receive an error
			if tt.expectedErr == nil {
				assert.NoError(t, actualErr)
			} else { //validate whether expected and actual errors are same
				assert.Equal(t, tt.expectedErr, actualErr)
			}

			if !reflect.DeepEqual(tt.expectedOutput, actualOutput) {
				t.Errorf("got %v, but want %v", actualOutput, tt.expectedOutput)
			}

			// Check that all expectations were met
			err = mock.ExpectationsWereMet()
			if err != nil {
				t.Errorf("Unfulfilled expectations: %s", err)
			}
		})
	}

}

func TestUserLogin(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput domain.Users
		buildStub      func(mock sqlmock.Sqlmock)
		expectedErr    error
	}{
		{
			name:  "userlogin",
			input: "sreenathsvrm@gmail.com",
			expectedOutput: domain.Users{
				ID:     1,
				Name:   "sreenath",
				Email:  "sreenathsvrm@gmail.com",
				Mobile: "+917994475799",
				Password: "sree@123",
				IsBlocked: false,
				CreatedAt: time.Time{},
			},
			buildStub: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`SELECT \* FROM users WHERE email\=?`).
					WithArgs("sreenathsvrm@gmail.com").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "mobile", "password", "is_blocked", "created_at"}).
                 AddRow(1, "sreenath", "sreenathsvrm@gmail.com", "+917994475799", "sree@123", false, time.Time{}))

			},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
			if err != nil {
				t.Fatalf("an error '%s' was not expected when initializing a mock db session", err)
			}

			userRepository := NewUserRepository(gormDB)

			tt.buildStub(mock)

			actualOutput, actualErr := userRepository.UserLogin(context.TODO(), tt.input)

			if tt.expectedErr == nil {
				assert.NoError(t, actualErr)
			} else {
				assert.EqualError(t, tt.expectedErr, actualErr.Error())
			}

			if !reflect.DeepEqual(tt.expectedOutput, actualOutput) {
				t.Errorf("got %v, but want %v", actualOutput, tt.expectedOutput)
			}

			err = mock.ExpectationsWereMet()
			if err != nil {
				t.Errorf("Unfulfilled expectations: %s", err)
			}
		})
	}
}
