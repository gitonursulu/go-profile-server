package main

import (
	"fmt"
)

func main() {
	//left := &predicate.ValueNode{Value: true}
	//right := &predicate.ValueNode{Value: false}

	//expressionTree := &predicate.PredicateNode{
	//	Operator: "AND",
	//	Left:     left,
	//	Right:    right,
	//}

	//result := expressionTree.Evaluate()
	fmt.Println("Sonuç:")

	// e := echo.New()

	// e.POST("/profile/create", create)

	// e.POST("/users", func(c echo.Context) error {
	// 	user := new(User)
	// 	if err := c.Bind(user); err != nil {
	// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// 	}

	// 	// JSON objesini işleme
	// 	fmt.Println("Gelen Kullanıcı:", user)

	// 	return c.String(http.StatusOK, "Kullanıcı kaydedildi")
	// })

	// e.Logger.Fatal(e.Start(":1323"))
}

// // e.POST("/save", save)
// func create(c echo.Context) error {
// 	// Get name and email
// 	name := c.FormValue("name")
// 	email := c.FormValue("email")
// 	println("done")
// 	return c.String(http.StatusOK, "name:"+name+", email:"+email)
// }
