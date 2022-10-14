package utils

import (
	b "echoapp/models"
	"net/http"
	"strconv"
    "fmt"

	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	u := &b.User{
		ID: b.Seq,
	}

    fmt.Println("oqipwjkedqwpiejdqwe",u)
    // bu kisimda bind yapio ne gelirse artk
	if err := c.Bind(u); err != nil {
		return err
	}
	b.Users[u.ID] = u
	b.Seq++
	return c.JSON(http.StatusCreated, u)
}

func GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, b.Users[id])
}

func UpdateUser(c echo.Context) error {
	u := new(b.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	b.Users[id].Name = u.Name
	return c.JSON(http.StatusOK, b.Users[id])
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(b.Users, id)
	return c.NoContent(http.StatusNoContent)
}

func GetAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, b.Users)
}
