package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/parnurzeal/gorequest"
	"github.com/yashsriv/dashboard-http/config"
	"github.com/yashsriv/dashboard-http/models"

	"github.com/jlaffaye/ftp"
	"golang.org/x/crypto/sha3"
	"gopkg.in/kataras/iris.v5"
)

// IsAuthenticated checks whether user is authenticated
func IsAuthenticated(ctx *iris.Context) {
	username := ctx.RequestHeader("X-Username-Header")
	timestamp := ctx.RequestHeader("X-Timestamp-Header")
	auth := ctx.RequestHeader("X-Auth-Header")

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

// CurrentUser fetches info of current user
func CurrentUser(ctx *iris.Context) {
	username := ctx.RequestHeader("X-Username-Header")
	var users []models.User
	err := config.DatabaseConnection.Q().
		Where("username = ?", username).
		All(&users)
	if err != nil {
		SendInternalServer(err, ctx)
		return
	}
	if len(users) > 0 {
		_ = ctx.JSON(iris.StatusOK, users[0])
	} else {
		_ = ctx.Text(iris.StatusNotFound, "")
	}
}

type loginJSON struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}

type loginAuth struct {
	Username  string `json:"username"`
	Timestamp string `json:"timestamp"`
	Auth      string `json:"auth"`
}

type loginResponse struct {
	User models.User `json:"user"`
	Auth loginAuth   `json:"auth"`
}

// Login sets the login cookie if successful
func Login(ctx *iris.Context) {
	loginInfo := loginJSON{}
	err := ctx.ReadJSON(&loginInfo)
	if err != nil {
		_ = ctx.Text(iris.StatusBadRequest, err.Error())
	} else {
		if loginInfo.verifyLoginInfo() {
			var user models.User
			var users []models.User
			err := config.DatabaseConnection.Q().
				Where("username = ?", loginInfo.Username).
				All(&users)
			if err != nil {
				SendInternalServer(err, ctx)
				return
			}
			if len(users) > 0 {
				user = users[0]
			} else {
				request := gorequest.New().Get("https://search.pclub.in/api/student").
					Param("username", loginInfo.Username)

				resp, body, errs := request.End()
				if errs != nil {
					SendInternalServer(errs[0], ctx)
				}
				if resp.StatusCode == 200 {
					err = json.Unmarshal([]byte(body), &user)
					if err != nil {
						SendInternalServer(err, ctx)
						return
					}
					err = config.DatabaseConnection.Create(&user)
					if err != nil {
						SendInternalServer(err, ctx)
						return
					}
				}
			}
			username := loginInfo.Username
			timestamp := fmt.Sprintf("%d", time.Now().Unix())
			secret := config.CookieSecret
			hashValue := []byte(username + ":" + timestamp + ":" + secret)
			hasher := sha3.New256()
			hasher.Write(hashValue)
			sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
			auth := loginAuth{
				Username:  username,
				Timestamp: timestamp,
				Auth:      sha,
			}
			_ = ctx.JSON(iris.StatusOK, loginResponse{User: user, Auth: auth})
		} else {
			_ = ctx.Text(iris.StatusNotFound, "")
		}
	}
}

func (lj *loginJSON) verifyLoginInfo() bool {
	conn, err := ftp.Dial("webhome.cc.iitk.ac.in:21")
	if err != nil {
		panic(err)
	}
	defer conn.Quit()
	err = conn.Login(lj.Username, lj.Password)
	defer conn.Logout()
	if err != nil {
		iris.Logger.Println(err)
		return false
	}
	return true
}

func checkHash(username string, timestamp string, auth string) bool {
	secret := config.CookieSecret
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
