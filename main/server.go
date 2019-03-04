package main

import (
	"github.com/labstack/echo"
	"net/http"
)
// get test
// http://localhost:1323/users/evan
func getUser(c echo.Context) error{
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
// post test
// curl -F "name=evan" -F "age=18" http://localhost:1323/save
func save(c echo.Context)  error{
	name := c.FormValue("name")
	age := c.FormValue("age")
	return c.String(http.StatusOK, "name: " + name + ", age:" + age)
}
// http://localhost:1323/show?team=x-men&member=evan
func show(c echo.Context)  error{
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team: " + team + ", member:" + member)
}
// format test
// curl -F "name=evan" -F "age=18" http://localhost:1323/users
type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Age string `json:"age" xml:"age" form:"age" query:"age"`
}

func jsonUser(c echo.Context)  error{
	u := new(User)
	if err := c.Bind(u); err!=nil{
		return err
	}
	// json
	return c.JSON(http.StatusCreated, u)
	// xml
	// return c.XML(http.StatusCreated, u)
}

func main(){
	e := echo.New()
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.POST("/save", save)
	e.POST("/users", jsonUser)
	e.Logger.Fatal(e.Start(":1323"))
}



