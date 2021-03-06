package util

import (
    "fmt"
    "golang.org/x/crypto/bcrypt"
    "net/http"
)


func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func GetHashingCost(hashedPassword []byte) int {
    cost, _ := bcrypt.Cost(hashedPassword) // 为了简单忽略错误处理
    return cost
}

func PassWordHashingHandler(w http.ResponseWriter, r *http.Request) {
    password := "secret"
    hash, _ := HashPassword(password) // 为了简单忽略错误处理

    fmt.Fprintln(w,"Password:", password)
    fmt.Fprintln(w, "Hash:    ", hash)

    match := CheckPasswordHash(password, hash)
    fmt.Fprintln(w,"Match:   ", match)

    cost := GetHashingCost([]byte(hash))
    fmt.Fprintln(w,"Cost:    ", cost)

}