package graph

import (
	"context"
	"log"

	"github.com/BigOplO/GO_INTERVIEW/graph/model"
	"github.com/BigOplO/GO_INTERVIEW/internal/db"
)

type Resolver struct{}

// get all the data from database
func (r *queryResolver) Breads(ctx context.Context) ([]*model.Bread, error) {

	log.Println("Fetching all breads")
	breads, err := db.GetAllBreads()
	if err != nil {
		log.Printf("Error fetching breads: %v", err)
		return nil, err
	}

	var res []*model.Bread
	for _, b := range breads {
		res = append(res, &model.Bread{
			ID:        b.ID,
			Name:      b.Name,
			CreatedAt: b.CreatedAt,
		})
	}
	log.Println("Returning all breads")
	return res, nil
}

// get the data by specific ID
func (r *queryResolver) Bread(ctx context.Context, id string) (*model.Bread, error) {

	log.Printf("Fetching bread with id %s", id)
	bread, err := db.GetBreadByID(id)
	if err != nil {
		log.Printf("Error fetching bread: %v", err)
		return nil, err
	}

	return &model.Bread{
		ID:        bread.ID,
		Name:      bread.Name,
		CreatedAt: bread.CreatedAt,
	}, nil
}
