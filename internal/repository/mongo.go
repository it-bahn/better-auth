package repository

import (
	"context"
)

type Repository interface {
	Create(*context.Context) map[string]interface{}
	Update(ctx *context.Context, userID string) map[string]interface{}
	Delete(ctx *context.Context, userID string) map[string]interface{}
	FindOne(ctx *context.Context, userID string) map[string]interface{}
}

/**
A function to create a new repository
*/
