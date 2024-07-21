package jwt

import (
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	var claim = Claims{Uid: 0, Username: "test"}
	tokenStr, _ := GenerateToken(&claim)
    fmt.Println(tokenStr)
}

func TestParseToken(t *testing.T) {
	// var claim = Claims{Uid: 0, Username: "username"}
	// tokenStr, _ := GenerateToken(&claim)
	// fmt.Println(tokenStr)
    tokenStr := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjE1NzA0MTYsInVpZCI6MCwidXNlcm5hbWUiOiJ1c2VybmFtZSJ9.lDTpPJBs-4xJvXkJ4_g_sPR3xZJsorOBnbwly4K_FDM"
	cl, err := JwtVerify(tokenStr)
	if (err != nil) {
		fmt.Println(err)
		return
	}
	fmt.Println(cl.Uid, cl.Username)
}