package jwt

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken(AppKey, AppSecret, "chenxingyu")
	fmt.Println(token)
	assert.Nil(t, err)
}

func TestParseToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBcHBLZXkiOiIzYzZlMGI4YTljMTUyMjRhODIyOGI5YTk4Y2ExNTMxZCIsIkFwcFNlY3JldCI6IjZhYmViZGZkY2EwNzg5OWZiZDAxMGNiZGNiZWIzNjNmIiwiVXNlck5hbWUiOiJjaGVueGluZ3l1IiwiZXhwIjoxNjQ1OTgzMTE3LCJpc3MiOiJjeHkifQ.b4Ds_9U2_-r78h9QyY5U4nbR9ndYA2cSiJoZVRe6efo"
	info, err := ParseToken(token)
	fmt.Println(info, err)
}
