package convert

import (
	"github.com/go-openapi/strfmt"
	"postgesql-cluster-console/internal/storage"
	"postgesql-cluster-console/models"
)

func EnvironmentToSwagger(env *storage.Environment) *models.ResponseEnvironment {
	return &models.ResponseEnvironment{
		CreatedAt:   strfmt.DateTime(env.CreatedAt),
		Description: env.Description,
		ID:          env.ID,
		Name:        env.Name,
		UpdatedAt: func() *strfmt.DateTime {
			if env.UpdatedAt == nil {
				return nil
			}
			updated := strfmt.DateTime(*env.UpdatedAt)

			return &updated
		}(),
	}
}

func EnvironmentsToSwagger(envs []storage.Environment) []*models.ResponseEnvironment {
	resp := make([]*models.ResponseEnvironment, 0, len(envs))
	for _, env := range envs {
		resp = append(resp, EnvironmentToSwagger(&env))
	}

	return resp
}
