package middle

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/valentinrb1/go-api-rest.git/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/ssh"
)

var key = []byte("lab6key")

func VerifyToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}

func GetUsers() ([]models.User, error) {
	cmd := exec.Command("sh", "-c", "grep '^[^:]*:[^:]*:[0-9]\\{4,\\}:' /etc/passwd | awk -F: '$3 >= 1000 && $1 != \"nobody\" {print $1}'")
	out, err := cmd.Output()

	if err != nil {
		return nil, fmt.Errorf("error getting users list: %v", err)
	}

	var userList []models.User

	users := strings.Split(string(out), "\n")

	for _, username := range users {
		if username != "" {
			var user models.User

			uid, err := GetUID(username)

			if err != nil {
				return nil, fmt.Errorf("error getting UID: %v", err)
			}

			createdAt, err := GetCreatedAt(username)

			if err != nil {
				return nil, fmt.Errorf("error getting creation date: %v", err)
			}

			user.ID = uid
			user.Username = username
			user.Password = ""
			user.CreatedAt = createdAt

			userList = append(userList, user)
		}
	}

	return userList, nil
}

func GetUID(username string) (int, error) {
	cmd := exec.Command("id", "-u", username)
	out, err := cmd.Output()

	if err != nil {
		return 0, fmt.Errorf("error getting UID: %v", err)
	}

	uid, err := strconv.Atoi(strings.TrimSpace(string(out)))

	if err != nil {
		return 0, fmt.Errorf("error converting UID to int: %v", err)
	}

	return uid, nil
}

func GetCreatedAt(username string) (string, error) {
	fileInfo, err := os.Stat(fmt.Sprintf("/home/%s", username))
	if err != nil {
		return "", fmt.Errorf("error getting created at: %v", err)
	}

	createdAt := fileInfo.ModTime()
	return createdAt.Format("2006-01-02 15:04:05"), nil
}

func AuthenticateUser(username string, password string) bool {
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	_, err := ssh.Dial("tcp", "localhost:22", config)
	return err == nil
}

func CreateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("error creating JWT: %v", err)
	}

	return tokenString, nil
}
