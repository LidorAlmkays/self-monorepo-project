package auth

import (
	"testing"
	"time"

	"github.com/LidorAlmkays/self-monorepo-project/apps/user/dtos/incoming"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/entities"
)

const testToken = "ac13a8c54f6b9cb299fb89961d159e78c36fe38bbf8b61b8a8f800f519673de5"

func TestGenerateSalt(t *testing.T) {
	auth := Auth{
		saltLetters:   "abcde",
		pepperLength:  2,
		pepperLetters: "abcde",
		saltLength:    5,
	}
	salt, err := auth.generateRandomSalt()
	if err != nil {
		t.Errorf("failed to create a random salt")
	}
	if len(salt) != auth.saltLength {
		t.Errorf("didn't receive expected salt length")
	}
}

func TestGeneratePepper(t *testing.T) {
	auth := Auth{
		saltLetters:   "abcde",
		pepperLength:  2,
		pepperLetters: "abcde",
		saltLength:    5,
	}
	pepper, err := auth.generateRandomPepper()
	if err != nil {
		t.Errorf("failed to create a random pepper")
	}
	if len(pepper) != auth.pepperLength {
		t.Errorf("didn't receive expected pepper length")
	}
}

func TestMakePepperFromStringFromIndexValues(t *testing.T) {
	auth := Auth{
		saltLetters:   "abcde",
		pepperLength:  3,
		pepperLetters: "abcde",
		saltLength:    5,
	}
	val, err := auth.makePepperFromStringFromIndexValues([]int{1, 2, 4})
	if err != nil {
		t.Errorf(err.Error())
	}
	if "bce" != val {
		t.Errorf("didn't receive expected pepper")
	}
	val, err = auth.makePepperFromStringFromIndexValues([]int{4, 4, 4})
	if err != nil {
		t.Errorf(err.Error())
	}
	if "eee" != val {
		t.Errorf("didn't receive expected pepper")
	}
	val, err = auth.makePepperFromStringFromIndexValues([]int{4, 5, 4})
	if err == nil {
		t.Errorf("the function should crash because the index of the 2 element is out of bound")
	}
}

func TestCreateHashPassword(t *testing.T) {
	auth := Auth{
		saltLetters:   "abcde",
		pepperLength:  3,
		pepperLetters: "abcde",
		saltLength:    5,
	}
	salt := "This is not a random salt"
	pepper := "aat"
	email := "test@gmail.com"
	password := "password"
	result := auth.createHashPassword(salt, email, password, pepper)
	expectedInCode := "f9f0711fa714162accbb5704e77fa82c59e7ce250661fadb2cb872fa27f7243e"
	if result != expectedInCode {
		t.Errorf("failed to create hash received unexpected result.\nreceived: %s", result)
	}
}

func TestGenerateSecretPassword(t *testing.T) {
	u := &entities.User{
		Name:     "Test",
		UserName: "TestUserName",
		Password: "This is a weird password!",
		BirthDay: time.Now(),
		Email:    "Test@gmail.com",
	}
	auth := Auth{
		saltLetters:   "abcde",
		pepperLength:  3,
		pepperLetters: "abcde",
		saltLength:    5,
	}
	SecretPassword, err := auth.GenerateSecretPassword(u)
	if err != nil {
		t.Errorf("failed to generate a secret password.\nreason: %s", err.Error())
	}
	t.Log("generated the password:", SecretPassword)
	testingUser := incoming.AuthenticateUserDTO{
		Email:    u.Email,
		Password: u.Password,
	}
	u.Password = SecretPassword
	if !auth.AuthenticateUser(testingUser, u) {
		t.Errorf("failed to authenticate created user")
	}

}

func TestAuthenticateUser(t *testing.T) {
	u := &entities.User{
		Name:     "Test",
		UserName: "TestUserName",
		Password: "This is a weird password!",
		BirthDay: time.Now(),
		Email:    "Test@gmail.com",
	}
	auth := Auth{
		saltLetters:   "abcde",
		pepperLength:  3,
		pepperLetters: "abcde",
		saltLength:    5,
	}
	SecretPassword, err := auth.GenerateSecretPassword(u)
	if err != nil {
		t.Errorf("failed to generate a secret password.\nreason: %s", err.Error())
	}
	t.Log("generated the password:", SecretPassword)
	testingUser := incoming.AuthenticateUserDTO{
		Email:    u.Email,
		Password: u.Password,
	}
	u.Password = SecretPassword
	if !auth.AuthenticateUser(testingUser, u) {
		t.Errorf("failed to authenticate created user")
	}

	failUserPassword := incoming.AuthenticateUserDTO{
		Email:    u.Email,
		Password: "Failme",
	}
	if auth.AuthenticateUser(failUserPassword, u) {
		t.Errorf("user auth was successful when it should have failed on password")
	}

	failUserEmail := incoming.AuthenticateUserDTO{
		Email:    "Failme@gmail.com",
		Password: u.Password,
	}
	if auth.AuthenticateUser(failUserEmail, u) {
		t.Errorf("user auth was successful when it should have failed on email")
	}

}
