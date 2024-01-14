package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
)

func main() {
	// Define the model for access control policies
	modelText := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && r.obj == p.obj && r.act == p.act
	`

	// Create a model from the text
	m, err := model.NewModelFromString(modelText)
	if err != nil {
		fmt.Println("Error creating model:", err)
		return
	}

	// Create an enforcer with the model and a sample policy file
	e, err := casbin.NewEnforcer(m, "path/to/policy.csv")
	if err != nil {
		fmt.Println("Error creating enforcer:", err)
		return
	}

	// Check if a user has permission to perform an action on an object
	sub := "alice"  // user
	obj := "data1"  // resource
	act := "read"   // action

	if res, err := e.Enforce(sub, obj, act); res {
		fmt.Printf("%s has permission to %s %s\n", sub, act, obj)
	} else {
		fmt.Printf("%s does not have permission to %s %s\n", sub, act, obj)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
