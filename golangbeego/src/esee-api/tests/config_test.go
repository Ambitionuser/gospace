package test

import (
	"esee-api/config"
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	fmt.Println("niaho1")
	user := config.BConfig.String("mysql.ods.user")
	fmt.Println(user)
}
