package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"go-clean-architecture/app/config"
	"go-clean-architecture/domains"
	"regexp"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
	// user           User
	userRepository UserRepository
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	assert.NoError(s.T(), err, "Failed to open gorm DB")
	assert.NotNil(s.T(), db, "Mock DB is null")
	assert.NotNil(s.T(), s.mock, "SQLMock is null")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)
	s.DB, err = gorm.Open(mysql.Open(dsn))
	assert.NoError(s.T(), err, "Failed to open gorm DB")
	assert.NotNil(s.T(), s.DB, "Mock DB is null")

	s.userRepository = UserRepository{DB: s.DB}
	defer db.Close()
}

func (s *Suite) TestUserRepository_FindAllSuccess() {
	rows := sqlmock.NewRows([]string{"id", "email", "password", "firstname", "lastname"}).AddRow("id1", "email1", "password1", "firstname1", "lastname1").AddRow("id2", "email2", "password2", "firstname2", "lastname2")
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).WillReturnRows(rows)
	users, err := s.userRepository.FindAll()

	s.NotEmpty(users)
	s.NotNil(users)
	s.Nil(err)
}

// Error
// func (s *Suite) TestUserRepository_FindAllFailed() {
// 	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).WillReturnError(errors.New("error"))
// 	users, err := s.userRepository.FindAll()
// 	s.NotEmpty(users)
// 	s.NoError(err)
// }

func (s *Suite) TestUserRepository_FindByIDSuccess() {
	query := "SELECT * FROM `users` WHERE id = ?"
	var id uint = 1
	rows := sqlmock.NewRows([]string{"id", "email", "password", "firstname", "lastname"}).AddRow("id1", "email1", "password1", "firstname1", "lastname1")
	s.mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(id).WillReturnRows(rows)
	users, err := s.userRepository.FindByID(id)

	s.NotEmpty(users)
	s.NotNil(users)
	s.Nil(err)
}
func (s *Suite) TestUserRepository_FindByIDFailed() {
	query := "SELECT * FROM `users` WHERE id = ?"
	var id uint = 0
	s.mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(id).WillReturnError(errors.New("error"))
	users, err := s.userRepository.FindByID(id)
	s.Empty(users)
	s.Error(err)
}

func (s *Suite) TestUserRepository_CreateSuccess() {
	user := &domains.Register{
		Email:     "email1",
		Password:  "password1",
		FirstName: "firstname1",
		LastName:  "lastname1",
	}

	const sqlInsert = `INSERT INTO "users" ("email","password","firstname","lastname") VALUES ($1,$2,$3,$4) RETURNING "id"`
	var newId uint = 3
	s.mock.ExpectBegin() // start transaction
	s.mock.ExpectQuery(regexp.QuoteMeta(sqlInsert)).WithArgs(user.Email, user.Password, user.FirstName, user.LastName).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(newId))
	s.mock.ExpectCommit() // commit transaction

	err := s.userRepository.Create(user)
	s.NoError(err)

}

func TestSuiteRepository(t *testing.T) {
	suite.Run(t, new(Suite))
}
