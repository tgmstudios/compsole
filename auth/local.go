package auth

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/BradHacker/compsole/ent"
	"github.com/BradHacker/compsole/ent/action"
	"github.com/BradHacker/compsole/ent/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserLoginVals struct {
	Username string `form:"username" json:"username" binding:"required" example:"admin"`
	Password string `form:"password" json:"password" binding:"required" example:"password123"`
}

// LocalLogin godoc
// @Summary Login with a local account
// @Schemes http https
// @Description Login with a local account
// @Tags Auth API
// @Accept json,mpfd
// @Param login body UserLoginVals true "User account details"
// @Produce json
// @Success 200 {object} ent.User
// @Header 200 {string} Cookie "`auth-cookie` contains the session token"
// @Router /auth/local/login [post]
func LocalLogin(client *ent.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hostname, ok := os.LookupEnv("GRAPHQL_HOSTNAME")
		if !ok {
			hostname = "localhost"
		}
		cookie_timeout := 60
		if env_value, exists := os.LookupEnv("COOKIE_TIMEOUT"); exists {
			if atio_value, err := strconv.Atoi(env_value); err == nil {
				cookie_timeout = atio_value
			}
		}
		secure_cookie := false
		if env_value, exists := os.LookupEnv("HTTPS_ENABLED"); exists {
			if env_value == "true" {
				secure_cookie = true
			}
		}
		var loginVals UserLoginVals
		username := ""
		password := ""

		if err := ctx.ShouldBind(&loginVals); err != nil {
			if secure_cookie {
				ctx.SetCookie("auth-cookie", "", 0, "/", hostname, true, true)
			} else {
				ctx.SetCookie("auth-cookie", "", 0, "/", hostname, false, false)
			}
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err})
			return
		} else {
			username = loginVals.Username
			password = loginVals.Password
		}

		// Get the client's IP address
		clientIp, err := ForContextIp(ctx)
		if err != nil {
			logrus.Warnf("failed to get IP from gin context: %v", err)
		}

		entUser, err := client.User.Query().Where(
			user.And(
				user.UsernameEQ(username),
				user.ProviderEQ(user.ProviderLOCAL),
			),
		).Only(ctx)
		if ent.IsNotFound(err) {
			err = client.Action.Create().
				SetIPAddress(clientIp).
				SetType(action.TypeFAILED_SIGN_IN).
				SetMessage(fmt.Sprintf("user \"%s\" does not exists", username)).
				SetActionToUser(entUser).
				Exec(ctx)
			if err != nil {
				logrus.Warn("failed to create FAILED_SIGN_IN action: %v", err)
			}
		}
		if err != nil {
			if secure_cookie {
				ctx.SetCookie("auth-cookie", "", 0, "/", hostname, true, true)
			} else {
				ctx.SetCookie("auth-cookie", "", 0, "/", hostname, false, false)
			}
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}
		// Compare the stored hashed password, with the hashed version of the password that was received
		if err = bcrypt.CompareHashAndPassword([]byte(entUser.Password), []byte(password)); err != nil {
			err = client.Action.Create().
				SetIPAddress(clientIp).
				SetType(action.TypeFAILED_SIGN_IN).
				SetMessage(fmt.Sprintf("wrong password for user \"%s\"", username)).
				SetActionToUser(entUser).
				Exec(ctx)
			if err != nil {
				logrus.Warn("failed to create FAILED_SIGN_IN action: %v", err)
			}
			if secure_cookie {
				ctx.SetCookie("auth-cookie", "", 0, "/", hostname, true, true)
			} else {
				ctx.SetCookie("auth-cookie", "", 0, "/", hostname, false, false)
			}
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}

		expiresAt := time.Now().Add(time.Minute * time.Duration(cookie_timeout)).Unix()

		claims := &CompsoleJWTClaims{
			IssuedAt: time.Now().Unix(),
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expiresAt,
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		jwtKey, exists := os.LookupEnv("JWT_SECRET")
		if !exists {
			if secure_cookie {
				ctx.SetCookie("auth-cookie", "", 0, "/", hostname, true, true)
			} else {
				ctx.SetCookie("auth-cookie", "", 0, "/", hostname, false, false)
			}
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Error signing token"})
			return
		}

		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			if secure_cookie {
				ctx.SetCookie("auth-cookie", "", 0, "/", hostname, true, true)
			} else {
				ctx.SetCookie("auth-cookie", "", 0, "/", hostname, false, false)
			}
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Error signing token"})
			return
		}

		_, err = client.Token.Create().SetTokenToUser(entUser).SetExpireAt(expiresAt).SetToken(tokenString).Save(ctx)
		if err != nil {
			if secure_cookie {
				ctx.SetCookie("auth-cookie", "", 0, "/", hostname, true, true)
			} else {
				ctx.SetCookie("auth-cookie", "", 0, "/", hostname, false, false)
			}
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Error updating token"})
			return
		}

		if secure_cookie {
			ctx.SetCookie("auth-cookie", tokenString, cookie_timeout*60, "/", hostname, true, true)
		} else {
			ctx.SetCookie("auth-cookie", tokenString, cookie_timeout*60, "/", hostname, false, false)
		}

		// Successful sign-in
		err = client.Action.Create().
			SetIPAddress(clientIp).
			SetType(action.TypeSIGN_IN).
			SetMessage(fmt.Sprintf("user \"%s\" has signed in successfully", username)).
			SetActionToUser(entUser).
			Exec(ctx)
		if err != nil {
			logrus.Warn("failed to create SIGN_IN action: %v", err)
		}

		entUser.Password = ""
		ctx.JSON(200, entUser)

		ctx.Next()
	}
}
