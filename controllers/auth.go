package controllers

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/olebedev/config"
	"golang.org/x/crypto/sha3"
	"gopkg.in/kataras/iris.v5"
)

// IsAuthenticated checks whether user is authenticated
func IsAuthenticated(ctx *iris.Context) {
	username := ctx.GetCookie("username")
	timestamp := ctx.GetCookie("timestamp")
	auth := ctx.GetCookie("auth")

	if username == "" || timestamp == "" || auth == "" {
		_ = ctx.Text(iris.StatusUnauthorized, "Unauthorised")
	} else {
		if checkHash(username, timestamp, auth) {
			ctx.Next()
		} else {
			_ = ctx.Text(iris.StatusForbidden, "Forbidden!!")
		}
	}
}

type loginJSON struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}

// Login sets the login cookie if successful
func Login(ctx *iris.Context) {
	loginInfo := loginJSON{}
	err := ctx.ReadJSON(&loginInfo)
	if err != nil {
		_ = ctx.Text(iris.StatusBadRequest, err.Error())
	} else {
		if loginInfo.verifyLoginInfo() {
			username := loginInfo.Username
			timestamp := fmt.Sprintf("%d", time.Now().Unix())
			secret := getSecret()
			hashValue := []byte(username + ":" + timestamp + ":" + secret)
			hasher := sha3.New256()
			hasher.Write(hashValue)
			sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
			ctx.SetCookieKV("username", username)
			ctx.SetCookieKV("timestamp", timestamp)
			ctx.SetCookieKV("auth", sha)
			_ = ctx.Text(iris.StatusOK, "")
		} else {
			_ = ctx.Text(iris.StatusNotFound, "")
		}
	}
}

func getSecret() string {

	cfg, _ := config.ParseYamlFile("config.yml")
	cfg.EnvPrefix("DASHBOARD")

	// Can be set using DASHBOARD_SECRET_VALUE environment variable
	secret, _ := cfg.String("secret.value")
	return secret

}

func (lj *loginJSON) verifyLoginInfo() bool {
	// TODO: Contact actual server
	return true
}

func checkHash(username string, timestamp string, auth string) bool {
	secret := getSecret()
	hashValue := []byte(username + ":" + timestamp + ":" + secret)
	hasher := sha3.New256()
	hasher.Write(hashValue)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return auth == sha
}

// SendInternalServer sends Internal Server message with error
func SendInternalServer(err error, ctx *iris.Context) {
	_ = ctx.JSON(iris.StatusInternalServerError, iris.Map{"error": err.Error()})
}
