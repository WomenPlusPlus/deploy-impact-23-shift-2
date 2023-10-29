package api

import (
	"context"
	"fmt"
	"shift/internal/entity"
)

func UserResourcePermission(ctx context.Context, id int) (bool, error) {
	kind, ok := ctx.Value(entity.ContextKeyKind).(string)
	if !ok {
		return false, fmt.Errorf("invalid value for user kind context")
	}
	
	if kind == entity.UserKindAdmin {
		return true, nil
	}

	userId, ok := ctx.Value(entity.ContextKeyUserId).(int)
	if !ok {
		return false, fmt.Errorf("invalid value for user id context")
	}
	
	return userId == id, nil
}
