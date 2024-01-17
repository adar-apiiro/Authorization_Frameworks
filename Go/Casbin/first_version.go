package main

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/casbin/model"
	"github.com/casbin/casbin/persist"
	"github.com/casbin/casbin/rbac"
)

func main() {
	// Initialize Casbin enforcer
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
	m, err := model.NewModelFromString(modelText)
	if err != nil {
		panic(err)
	}

	// Create an example policy (you would typically load this from a persistent store)
	policy := [][]string{
		{"alice", "data1", "read"},
		{"bob", "data2", "write"},
	}

	// Initialize the adapter (you would choose an appropriate adapter based on your storage needs)
	adapter := persist.NewMemoryAdapter()
	for _, p := range policy {
		adapter.AddPolicy("", "p", p)
	}

	// Initialize Casbin enforcer with model and adapter
	enforcer, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		panic(err)
	}

	// Check authorization
	subject := "alice"
	object := "data1"
	action := "read"
	if enforcer.Enforce(subject, object, action) {
		fmt.Printf("%s has permission to %s %s\n", subject, action, object)
	} else {
		fmt.Printf("%s does not have permission to %s %s\n", subject, action, object)
	}

	// RBAC example: Adding roles and checking permissions
	enforcer.AddRoleForUser("alice", "admin")
	if enforcer.Enforce("alice", "data1", "read") {
		fmt.Printf("%s has permission to read %s\n", subject, object)
	} else {
		fmt.Printf("%s does not have permission to read %s\n", subject, object)
	}

	// Additional RBAC functionalities
	roleManager := enforcer.GetRoleManager().(rbac.RoleManager)
	roleManager.AddLink("admin", "user")  // Adding a role link
	usersWithRole, _ := roleManager.GetUsers("user") // Getting users with a role
	fmt.Printf("Users with role 'user': %v\n", usersWithRole)
}
