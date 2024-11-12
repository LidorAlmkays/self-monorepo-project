package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"math"
	"math/rand"

	"github.com/LidorAlmkays/self-monorepo-project/apps/user/dtos/incoming"
	"github.com/LidorAlmkays/self-monorepo-project/apps/user/internal/entities"
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/logger"
)

type Auth struct {
	saltLetters   string
	pepperLength  int
	pepperLetters string
	saltLength    int
	l             logger.CustomLogger
}

func NewPepperSaltAuthenticator(saltLetters, pepperLetters string, saltLength, pepperLength int, l logger.CustomLogger) AuthPort {
	return &Auth{
		saltLetters,
		pepperLength,
		pepperLetters,
		saltLength,
		l,
	}
}

func (a *Auth) generateRandomSalt() (string, error) {
	if len(a.saltLetters) <= 0 {
		return "", errors.New("salt letters string is empty cant create a random salt")
	}
	b := make([]byte, a.saltLength)
	for i := range b {
		b[i] = a.saltLetters[rand.Intn(len(a.saltLetters))]
	}
	return string(b), nil
}

func (a *Auth) generateRandomPepper() (string, error) {
	if len(a.pepperLetters) == 0 {
		return "", errors.New("the pepper string is empty") // return zero value if the string is empty
	}
	if a.pepperLength < 0 {
		return "", errors.New("invalid amount of pepper letters")
	}
	pepper := make([]byte, a.pepperLength)
	for i := 0; i < a.pepperLength; i++ {
		pepper[i] = a.pepperLetters[rand.Intn(len(a.pepperLetters))]
	}
	return string(pepper), nil
}

func (a *Auth) createHashPassword(salt, email, password string, pepper string) string {
	combined := salt + password + pepper + email
	// Create a SHA-256 hash of the combined string
	hash := sha256.Sum256([]byte(combined))
	return hex.EncodeToString(hash[:])
}

func (a *Auth) GenerateSecretPassword(u *entities.User) (string, error) {
	// Combine username, password, and salt
	pepper, err := a.generateRandomPepper()
	if err != nil {
		return "", err
	}
	salt, err := a.generateRandomSalt()
	if err != nil {
		return "", err
	}

	u.Salt = salt
	return a.createHashPassword(salt, u.Email, u.Password, pepper), nil
}

// authenticate incoming user data with user from database
func (a *Auth) AuthenticateUser(loginInfo incoming.AuthenticateUserDTO, userFromDb *entities.User) bool {
	pepperIndexes := make([]int, a.pepperLength)
	for range int(math.Pow(float64(len(a.pepperLetters)), float64(a.pepperLength))) {
		if pepperIndexes[0] == len(a.pepperLetters) {
			for j := 0; j < a.pepperLength-1 && pepperIndexes[j] == len(a.pepperLetters); j++ {
				pepperIndexes[j] = 0
				pepperIndexes[j+1]++
			}
		}
		currentPepper, err := a.makePepperFromStringFromIndexValues(pepperIndexes)
		if err != nil {
			a.l.Error(err)
			return false
		}
		if userFromDb.Password == a.createHashPassword(userFromDb.Salt, loginInfo.Email, loginInfo.Password, currentPepper) {
			return true
		}
		pepperIndexes[0]++
	}
	return false
}

func (a *Auth) GenerateRandomToken() (string, error) {
	n := 512
	b := make([]byte, n)
	_, err := rand.Read(b)

	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return "", err
	}

	hasher := sha256.New()
	hasher.Write(b)
	sha := hex.EncodeToString(hasher.Sum(nil))

	return sha, nil

}

func (a *Auth) makePepperFromStringFromIndexValues(indexes []int) (string, error) {
	pepper := ""
	for _, val := range indexes {
		if val >= len(a.pepperLetters) {
			return "", errors.New("the index received in the pepper creation is out of bound")
		}
		pepper += string(a.pepperLetters[val])
	}
	return pepper, nil
}
