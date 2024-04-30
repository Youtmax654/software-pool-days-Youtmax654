package jwt

import (
	"SoftwareGoDay3/server"

	"github.com/golang-jwt/jwt"
)

// Function to check if a given email is already registered
func IsRegistered(email string) bool {
    _, ok := users[email]
    return ok
}

// Function to register a new user
func Register(user User) AuthResponse {
    tokenString, err := GenerateToken(user.Email, user.Password)
    if err != nil {
        return AuthResponse{
            Message: "Failed to generate token",
        }
    }

    users[user.Email] = user

    return AuthResponse{
        AccessToken: tokenString,
        User: user,
        Message: "User successfully created",
    }
}

// Function to generate a token for a given email
func GenerateToken(email, password string) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)
    token.Claims = jwt.MapClaims{
        "email": email,
        "password": password,
    }

    return token.SignedString([]byte(server.Config.SecretKey))
}

// Function to check if a given password is correct
func IsPasswordCorrect(email, password string) bool {
    user, ok := users[email]
    if !ok {
        return false
    }

    return user.Password == password
}

// Function to get the user associated with a given token
func GetUserFromToken(tokenString string) (User, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        _, ok := token.Method.(*jwt.SigningMethodHMAC)
        if !ok {
            return nil, jwt.ErrSignatureInvalid
        }

        return []byte(server.Config.SecretKey), nil
    })

    if err != nil {
        return User{}, err
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return User{}, jwt.ErrInvalidKeyType
    }

    return User{
        Email: claims["email"].(string),
        Password: claims["password"].(string),
    }, nil
}