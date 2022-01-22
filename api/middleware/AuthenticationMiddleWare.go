package middleware

import (
	"encoding/json"
	"facegram_file_server/config"
	"facegram_file_server/model/usermodel"
	"facegram_file_server/pkg/_crypto"
	"facegram_file_server/pkg/apiresponse"
	"facegram_file_server/pkg/utility"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func DocumentRouteCheckAccess(c *gin.Context) {
	c.Next()
}

func UploadRouteCheckAccess(c *gin.Context) {
	if token := getAccessToken(c); token != "" {
		var claims utility.JwtClaims
		key := *config.GetExternalKeyConfig()
		jwtKey := []byte(key)
		tokenInfo, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if !tokenInfo.Valid {
			apiresponse.SendUnAuthorize(c)
		} else if err != nil {
			if err == jwt.ErrSignatureInvalid {
				apiresponse.SendUnAuthorize(c)
			} else {
				apiresponse.SendInternalError(c)
			}
		} else {
			userData, err := _crypto.AesDecrypt(claims.UserData, key)
			if err != nil {
				apiresponse.SendInternalError(c)
			} else {
				if user, err := userClaims(userData); err == nil {
					if exist, err := checkExistSession(user.Sid); err == nil {
						if exist {
							if user.Kind == "unprotected" {
								userModel := usermodel.UserInfo{
									UserID:    user.UserID,
									PageID:    user.PageID,
									Role:      user.Role,
									Token:     token,
									SessionID: user.Sid,
								}
								err := saveUserInfo(userModel, c)
								if err != nil {
									apiresponse.SendInternalError(c)
								} else {
									c.Next()
								}
							} else {
								apiresponse.SendUnAuthorize(c)
							}
						} else {
							apiresponse.SendUnAuthorize(c)
						}
					} else {
						apiresponse.SendInternalError(c)
					}
				} else {
					apiresponse.SendInternalError(c)
				}
			}
		}
	} else {
		apiresponse.SendInvalidInputError(c)
	}
}

func getAccessToken(c *gin.Context) string {

	if token, ok := c.Request.Header["token"]; ok {
		return token[0]
	}

	return ""
}

func saveUserInfo(data usermodel.UserInfo, c *gin.Context) error {
	model, e := json.Marshal(data)
	if e != nil {
		return e
	}
	c.Set("user_info", string(model))
	return nil
}

func userClaims(data string) (usermodel.UserTokenInfo, error) {
	var model usermodel.UserTokenInfo
	e0 := json.Unmarshal([]byte(data), &model)
	if e0 != nil {
		return model, e0
	}
	return model, nil
}

func checkExistSession(id string) (bool, error) {
	return true, nil
}
