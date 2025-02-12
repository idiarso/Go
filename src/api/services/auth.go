package services

import (
    "errors"
    "github.com/dgrijalva/jwt-go/v4"
    "golang.org/x/crypto/bcrypt"
    "github.com/idiarso/belajar-git/src/api/models"
)

var users []models.User

func Register(username, password, role string) (models.User, error) {
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
    user := models.User{Username: username, Password: string(hashedPassword), Role: role}
    users = append(users, user)
    return user, nil
}

func Login(username, password string) (string, error) {
    for _, user := range users {
        if user.Username == username {
            err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
            if err != nil {
                return "", errors.New("invalid password")
            }
            token := jwt.New(jwt.SigningMethodHS256)
            return token.SignedString([]byte("your_secret_key"))
        }
    }
    return "", errors.New("user not found")
}
