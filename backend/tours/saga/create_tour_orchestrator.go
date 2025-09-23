package orchestrator

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"database-example/clients"
	"database-example/model"
	"database-example/repo"
)

type CreateTourOrchestrator struct {
	tourRepo   repo.TourRepo
	blogClient clients.BlogClient
	logger     *log.Logger
}

func NewCreateTourOrchestrator(tourRepo repo.TourRepo, blogClient clients.BlogClient, logger *log.Logger) *CreateTourOrchestrator {
	return &CreateTourOrchestrator{
		tourRepo:   tourRepo,
		blogClient: blogClient,
		logger:     logger,
	}
}

func (o *CreateTourOrchestrator) Execute(ctx context.Context, tour *model.Tour, authorID string) error {

	if tour.ID.IsZero() {
		tour.ID = primitive.NewObjectID()
	}
	if tour.Status == "" {
		tour.Status = "draft"
	}
	tour.AuthorID = authorID

	if err := o.tourRepo.Create(ctx, tour); err != nil {
		return fmt.Errorf("orchestrator: failed to create tour: %w", err)
	}
	o.logger.Printf("orchestrator: tour created %s", tour.ID.Hex())

	if err := o.blogClient.CreateBlogPost(ctx, tour, authorID); err != nil {
		o.logger.Printf("orchestrator: blog creation failed, compensating by deleting tour %s: %v", tour.ID.Hex(), err)

		if derr := o.tourRepo.Delete(ctx, tour.ID); derr != nil {
			o.logger.Printf("orchestrator: compensation delete failed for tour %s: %v", tour.ID.Hex(), derr)
			return fmt.Errorf("orchestrator: blog creation failed: %v; compensation delete failed: %v", err, derr)
		}

		return fmt.Errorf("orchestrator: blog creation failed: %w", err)
	}

	o.logger.Printf("orchestrator: saga finished successfully for tour %s", tour.ID.Hex())
	return nil
}
