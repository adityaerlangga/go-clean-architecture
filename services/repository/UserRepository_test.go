package repository

import (
	"database/sql"
	"go-clean-architecture/domains"
	"go-clean-architecture/utils"
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
	DB             *gorm.DB
	mock           sqlmock.Sqlmock
	userRepository *UserRepository
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
	s.DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	assert.NoError(s.T(), err, "Failed to open gorm DB")
	s.userRepository = InitUserRepository(s.DB)
}

func (s *Suite) TestUserRepository_FindAllSuccess() {
	rows := sqlmock.NewRows([]string{"id", "email", "password", "firstname", "lastname"}).AddRow("1", "email1", "password1", "firstname1", "lastname1").AddRow("2", "email2", "password2", "firstname2", "lastname2")
	s.mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users`")).WillReturnRows(rows)
	users, err := s.userRepository.FindAll()
	s.NotEmpty(users)
	s.NoError(err)
}

func (s *Suite) TestUserRepository_FindByIDSuccess() {
	query := "SELECT * FROM `users` WHERE `users`.`id` = ? ORDER BY `users`.`id` LIMIT 1"
	var id uint = 1
	rows := sqlmock.NewRows([]string{"id", "email", "password", "firstname", "lastname"}).AddRow("1", "email1", "password1", "firstname1", "lastname1")
	s.mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(id).WillReturnRows(rows)
	users, err := s.userRepository.FindByID(id)

	s.NotEmpty(users)
	s.NoError(err)
}

func (s *Suite) TestUserRepository_FindByEmailSuccess() {
	query := "SELECT * FROM `users` WHERE email = ? ORDER BY `users`.`id` LIMIT 1"
	email := "aditya@gmail.com"
	rows := sqlmock.NewRows([]string{"id", "email", "password", "firstname", "lastname"}).AddRow("1", "aditya@gmail.com", "aditya123", "Aditya", "Erlangga")
	s.mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(email).WillReturnRows(rows)
	user, err := s.userRepository.FindByEmail(email)
	s.NotEmpty(user)
	s.NoError(err)
}

func (s *Suite) TestUserRepository_Update() {
	user := &domains.ChangePassword{
		ID:          1,
		OldPassword: "aditya123",
		NewPassword: utils.PasswordHash("aditya1234"),
	}
	const sqlUpdate = "UPDATE `users` SET `password`=? WHERE id = ?"

	s.mock.ExpectBegin() // start transaction
	s.mock.ExpectExec(regexp.QuoteMeta(sqlUpdate)).WithArgs(user.NewPassword, user.ID).WillReturnResult(sqlmock.NewResult(0, 1))
	s.mock.ExpectCommit() // commit transaction
	err := s.userRepository.UpdatePassword(user)
	s.NoError(err)
}

func (s *Suite) TestUserRepository_Create() {
	user := &domains.Register{
		Email:     "dummy@gmail.com",
		Password:  "password_dummy",
		FirstName: "firstname_dummy",
		LastName:  "lastname_dummy",
	}

	// const sqlInsert = `INSERT INTO "users" ("email","password","firstname","lastname") VALUES ($1,$2,$3,$4) RETURNING "id"`
	// const sqlInsert = "INSERT INTO `users` (`email`,`password`,`firstname`,`lastname`) VALUES (?,?,?,?)"
	const sqlInsert = "INSERT INTO `users` (`first_name`,`last_name`,`email`,`password`) VALUES (?,?,?,?)"
	// var newId uint = 3
	s.mock.ExpectBegin() // start transaction
	s.mock.ExpectExec(regexp.QuoteMeta(sqlInsert)).WithArgs(user.FirstName, user.LastName, user.Email, user.Password).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit() // commit transaction

	err := s.userRepository.Create(user)
	s.NoError(err)
}

func (s *Suite) TestUserRepository_Delete() {
	// Find By ID
	queryFindId := "SELECT * FROM `users` WHERE `users`.`id` = ? ORDER BY `users`.`id` LIMIT 1"
	var id uint = 1
	rows := sqlmock.NewRows([]string{"id", "email", "password", "firstname", "lastname"}).AddRow("1", "email1", "password1", "firstname1", "lastname1")
	// Delete
	queryDelete := "DELETE FROM `users` WHERE `users`.`id` = ?"

	s.mock.ExpectQuery(regexp.QuoteMeta(queryFindId)).WithArgs(id).WillReturnRows(rows)
	s.mock.ExpectBegin()
	s.mock.ExpectExec(regexp.QuoteMeta(queryDelete)).WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))
	s.mock.ExpectCommit()
	err := s.userRepository.Delete(id)
	s.NoError(err)
}

func TestSuiteRepository(t *testing.T) {
	suite.Run(t, new(Suite))
}
