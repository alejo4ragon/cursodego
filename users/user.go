package users

import (
	"fmt"
	"godesde0/modelos"
	"time"
)

func AltaUsuario() {
	user := new(modelos.User)
	user.AddUser(1, "Johan", time.Now(), true)

	fmt.Println(user)
}
