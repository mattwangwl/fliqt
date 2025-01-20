package migration

import (
	"context"
	"fliqt/internal/model"
	"fliqt/internal/repository"
)

func newSeed() *Seed {
	return &Seed{
		dao: repository.New(),
	}
}

type Seed struct {
	dao *repository.Repository
}

func (s *Seed) ExecAll(ctx context.Context) error {
	if err := s.Employee(ctx); err != nil {
		return err
	}

	return nil

}

func (s *Seed) Employee(ctx context.Context) error {
	models := []model.Employee{
		{Name: "Alice", Position: "Manager", Status: 1},
		{Name: "Bob", Position: "Developer", Status: 1},
		{Name: "Charlie", Position: "Developer", Status: 1},
		{Name: "David", Position: "Developer", Status: 0},
		{Name: "Eve", Position: "Designer", Status: 1},
	}

	return s.dao.Employee().Create(ctx, models...)
}
