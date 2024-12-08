package entity

import "context"

const RoleKey = iota

const (
	RoleAnonym = iota
	RoleUser
	RoleAdmin
)

func GetUserRole(ctx context.Context) int {
	v, ok := ctx.Value(RoleKey).(int)
	if !ok {
		return 0
	}

	return v
}
