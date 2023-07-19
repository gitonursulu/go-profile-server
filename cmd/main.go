package main

import (
	"fmt"
	"profile/pkg/predicate"
)

func main() {
	// ((true AND false) OR (true AND NOT false)) AND (true OR false)
	expressionTree := &predicate.PredicateNode{
		Operator: "AND",
		Left: &predicate.PredicateNode{
			Operator: "OR",
			Left: &predicate.PredicateNode{
				Operator: ">",
				Left:     &predicate.ValueNode{Value: "age", ValueSpec: predicate.Column},
				Right:    &predicate.ValueNode{Value: 20},
			},
			Right: &predicate.PredicateNode{
				Operator: "AND",
				Left:     &predicate.ValueNode{Value: true},
				Right: &predicate.PredicateNode{
					Operator: "=",
					Left:     &predicate.ValueNode{Value: "name", ValueSpec: predicate.Column},
					Right:    &predicate.ValueNode{Value: "onur"},
				},
			},
		},
		Right: &predicate.PredicateNode{
			Operator: "OR",
			Left:     &predicate.ValueNode{Value: true},
			Right:    &predicate.ValueNode{Value: false},
		},
	}

	result := expressionTree.Evaluate()

	fmt.Println("Sonuç:", result)

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
