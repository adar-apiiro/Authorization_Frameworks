package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist/file-adapter"
)

func main() {
	// Define the Casbin model
	modelText := `
	[request_definition]
	r = sub, obj, act

	[policy_definition]
	p = sub, obj, act

	[role_definition]
	g = _, _

	[policy_effect]
	e = some(where (p.eft == allow))

	[matchers]
	m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
	`

	// Create an Enforcer
	m, err := model.NewModelFromString(modelText)
	if err != nil {
		panic(err)
	}

	adapter := fileadapter.NewAdapter("path/to/policy.csv")

	e, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		panic(err)
	}

	// Load the policy from the adapter
	err = e.LoadPolicy()
	if err != nil {
		panic(err)
	}

	// Add roles, users, and permissions
	e.AddGroupingPolicy("alice", "admin")
	e.AddPolicy("admin", "data1", "read")
	e.AddPolicy("admin", "data2", "read")
	e.AddPolicy("bob", "data2", "write")

	// Check permissions
	hasPermission, err := e.Enforce("alice", "data1", "read")
	if err != nil {
		panic(err)
	}

	if hasPermission {
		fmt.Println("Alice has permission to read data1")
	} else {
		fmt.Println("Alice does not have permission to read data1")
	}
}
