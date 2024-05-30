package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

type Account struct {
	gorm.Model
	// Username string `json:"username" gorm:"column:username"`
	Username string `json:"username" gorm:"column:username; unique:true"`
	Password string `json:"password" gorm:"column:password;"`
	Name     string `json:"name" gorm:"column:name;"`
}

type JWTClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func middleHandleToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("user", nil)
		if ctx.Request.Header["Authorization"] == nil {
			fmt.Printf("\n----> no authorize\n")
			ctx.Set("user", nil)
		} else {
			bearerToken := ctx.Request.Header["Authorization"][0]
			parseToken := strings.Split(bearerToken, " ")
			if parseToken == nil || len(parseToken) != 2 {
				fmt.Printf("\n----> token is not valid 1\n")
				ctx.Set("user", nil)
			}
			token, err := VerifyToken(parseToken[1])
			if err != nil {
				fmt.Printf("\n----> token is not valid 2: %v\n", err)
				ctx.Set("user", nil)
			}
			claims, ok := token.Claims.(*JWTClaims)

			if ok && claims != nil {
				if float64(time.Now().Unix()) > float64(claims.ExpiresAt.Unix()) {
					fmt.Printf("\n----> token is expire\n")
					ctx.Set("user", nil)
				} else {
					var user Account
					if err := Db.Where(&Account{Username: claims.Username}).First(&user).Error; err != nil {
						ctx.Set("user", nil)
					} else {
						fmt.Printf("\n----> set %v\n", user.ID)
						ctx.Set("user", user.ID)
					}
				}
			}
		}
		ctx.Next()
	}
}

func middlePrivateRoute(ctx *gin.Context) {
	userId := ctx.MustGet("user")
	if userId == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		ctx.Abort()
	} else {
		ctx.Next()
	}

}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error for load env")
	}
}

func ConnectDatabase() {
	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	user := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")

	// set up postgres sql to open it.
	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbName, password)

	DB, err := gorm.Open(postgres.Open(psqlSetup), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: ")
	} else {
		Db = DB
		Db.AutoMigrate(&Account{})
		fmt.Printf("\n---------------Connect DB successful---------------\n")

	}

}

func (account *Account) EncriptPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	account.Password = string(bytes)
	return nil
}

func (account *Account) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
}

func (account *Account) RegisterToken() (string, error) {
	expirationTime := time.Now().Add(time.Hour)
	claims := &JWTClaims{
		Username: account.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaim.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	// Parse the token using the secret key
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil

}

func createAccount(ctx *gin.Context) {
	var user Account
	ctx.ShouldBindJSON(&user)

	if err := user.EncriptPassword(user.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if err := Db.Table("accounts").Create(&user).Error; err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(201, gin.H{
		"data": user,
	})
}

func userLogin(ctx *gin.Context) {
	var loginData struct {
		Username string `json: "username"`
		Password string `json: "password"`
	}
	// var loginData LoginData
	ctx.ShouldBindJSON(&loginData)
	var user Account
	if err := Db.Where(&Account{Username: loginData.Username}).First(&user).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "username is not correct",
		})
		return
	}
	if err := user.VerifyPassword(loginData.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "password is not correct",
		})
		return
	}

	token, err := user.RegisterToken()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Token error",
		})
		return
	}

	ctx.JSON(200, gin.H{"token": token})
}

func getAccounts(ctx *gin.Context) {
	var account []Account
	Db.Find(&account)
	ctx.JSON(http.StatusOK, account)
}

func getProfile(ctx *gin.Context) {

}

func main() {
	LoadEnv()
	router := gin.Default()
	ConnectDatabase()
	api := router.Group("/api")
	api.Use(middleHandleToken())
	api.GET("/", middlePrivateRoute, getAccounts)
	api.POST("/create", createAccount)
	api.POST("/login", userLogin)

	router.Run(":9000") // listen and serve on 0.0.0.0:8080
}
