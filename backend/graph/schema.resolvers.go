package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/solaoi/jms/graph/generated"
	"github.com/solaoi/jms/graph/model"
)

func (r *mutationResolver) CreateTemplate(ctx context.Context, input model.NewTemplate) (*model.Template, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	template := model.Template{
		Title:     input.Title,
		Content:   input.Content,
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
	}
	r.DB.Create(&template)

	return &template, nil
}

func (r *queryResolver) Templates(ctx context.Context) ([]*model.Template, error) {
		templates := []*model.Template{}

		r.DB.Find(&templates)
	
		return templates, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
