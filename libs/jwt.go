package libs

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	AudienceApp = "a" // 接收方 app
	AudienceWap = "w" // 接收方 wap
	DefaultIss  = ``
	SecretKey   = `AuthCenter`
)

func NewJwt() *Jwt {
	return &Jwt{
		SecretKey: SecretKey,
	}
}

type Jwt struct {
	Audience    string `json:"aud,omitempty"` // aud 标识token的接收者.
	ExpiresAt   int64  `json:"exp,omitempty"` // exp 过期时间.通常与Unix UTC时间做对比过期后token无效
	Id          string `json:"jti,omitempty"` // jti 是自定义的id号
	IssuedAt    int64  `json:"iat,omitempty"` // iat 签名发行时间.
	Issuer      string `json:"iss,omitempty"` // iss 是签名的发行者.
	NotBefore   int64  `json:"nbf,omitempty"` // nbf 这条token信息生效时间.这个值可以不设置,但是设定后,一定要小于当前Unix UTC,否则token将会延迟生效.
	Subject     string `json:"sub"`           // sub 签名面向的用户(用户名)
	SecretKey   string `json:"secretKey"`     // jwt签名密钥
	TokenString string `json:"token"`         // 生成的token
}

// 创建 Token
func (j *Jwt) GenerateToken() (string, error) {

	jwtSvc := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	if j.Issuer == "" {
		j.Issuer = DefaultIss
	}
	if j.Audience == "" {
		j.Audience = AudienceWap
	}

	claims["iss"] = j.Issuer
	claims["exp"] = j.ExpiresAt
	claims["iat"] = time.Now().Unix() - 100
	claims["sub"] = j.Subject
	claims["aud"] = j.Audience
	claims["jti"] = j.Audience
	jwtSvc.Claims = claims
	token, err := jwtSvc.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}

	return token, err
}

// 解析token
func (j *Jwt) ParseToken() error {
	token, err := jwt.Parse(j.TokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})

	if err != nil {
		return fmt.Errorf("jwt parse error | err | %s", err.Error())
	}
	var (
		ok           bool
		audienceMaps = map[string]int{
			AudienceApp: 1,
			AudienceWap: 1,
		}
		claims jwt.MapClaims
	)
	if claims, ok = token.Claims.(jwt.MapClaims); !ok {
		return fmt.Errorf("token assert claims failed | claims | %+v ", claims)
	}
	if !token.Valid {
		return fmt.Errorf("token invalid")
	}

	if j.Subject, ok = claims["sub"].(string); !ok {
		return fmt.Errorf("claims assert subject to string failed | claims | %+v ", claims)
	}
	if j.Audience, ok = claims["aud"].(string); !ok {
		return fmt.Errorf("claims assert audience to string failed | claims | %+v ", claims)
	}

	// 判断token 受众类型  // 只有 app 和 wap 可以通过 manage 用户在后台 不在go项目中
	if _, ok = audienceMaps[j.Audience]; !ok {
		return fmt.Errorf("audience is invalid | audience | %s", j.Audience)
	}

	return nil
}

const (
	Subject = `sub`
	Object  = `object`
	Domain  = `domain`
)
