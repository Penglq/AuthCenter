package libs

import "testing"

func TestJwt_CreateToken(t *testing.T) {
	jwt := Jwt{
		Issuer:    "auth",
		SecretKey: "35rm97g62l",
		NotBefore: 7,
		Audience:  AudienceApp,
		Subject:   "1326",
	}
	token, err := jwt.CreateToken()
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}

func TestJwt_ParseToken(t *testing.T) {
	jwt := Jwt{
		Issuer:      "auth",
		SecretKey:   "35rm97g62l",
		NotBefore:   7,
		Audience:    AudienceApp,
		Subject:     "1326",
		TokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhIiwiZXhwIjoiMTk3MC0wMS0wMVQwODowMDowMCswODowMCIsImlhdCI6MTU1NTkyODE1MCwiaXNzIjoiYXV0aCIsImp0aSI6ImEiLCJzdWIiOiIxMzI2In0._PbJBHRWO3fI2Evj6XB8B5r_VVPpO0gsbtUeRRhDnA8",
	}
	sub, err := jwt.ParseToken()
	if err != nil {
		t.Error(err)
	}
	t.Log(sub)
}
