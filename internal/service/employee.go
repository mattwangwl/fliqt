package service

import (
	"context"
	"errors"
	"fliqt/internal/model"
	"fliqt/internal/repository"
)

func NewEmployee() *Employee {
	dao := repository.New()
	return &Employee{
		employeeDao: dao.Employee(),
	}
}

type Employee struct {
	employeeDao repository.IEmployee
}

func (e *Employee) Get(ctx context.Context, id int64) (*model.Employee, error) {
	result, err := e.employeeDao.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (e *Employee) List(ctx context.Context, cond model.EmployeeListCond) ([]*model.Employee, error) {
	if cond.Status != nil && *cond.Status != 0 && *cond.Status != 1 {
		return nil, errors.New("status must be 0 or 1")
	}
	if len(cond.Position) > 0 {
		for _, p := range cond.Position {
			if p != "Manager" && p != "Developer" && p != "Designer" {
				return nil, errors.New("position must be Manager, Developer, or Designer")
			}
		}
	}

	result, err := e.employeeDao.List(ctx, cond)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (e *Employee) Create(ctx context.Context, model model.Employee) error {
	if model.Status != 0 && model.Status != 1 {
		return errors.New("status must be 0 or 1")
	}
	if model.Position != "Manager" && model.Position != "Developer" && model.Position != "Designer" {
		return errors.New("position must be Manager, Developer, or Designer")
	}
	if model.Name == "" {
		return errors.New("name is required")
	}

	if err := e.employeeDao.Create(ctx, model); err != nil {
		return err
	}
	return nil
}

func (e *Employee) Update(ctx context.Context, cond model.EmployeeUpdateCond) error {
	if cond.Status != nil && *cond.Status != 0 && *cond.Status != 1 {
		return errors.New("status must be 0 or 1")
	}
	if cond.Position != nil && *cond.Position != "Manager" && *cond.Position != "Developer" && *cond.Position != "Designer" {
		return errors.New("position must be Manager, Developer, or Designer")
	}
	if cond.Name != nil && *cond.Name == "" {
		return errors.New("name is required")
	}

	if err := e.employeeDao.Update(ctx, cond); err != nil {
		return err
	}
	return nil
}

func (e *Employee) UpdateStatus(ctx context.Context, id int64, status int) error {
	if status != 0 && status != 1 {
		return errors.New("status must be 0 or 1")
	}

	if err := e.employeeDao.Update(ctx, model.EmployeeUpdateCond{
		ID:     id,
		Status: &status,
	}); err != nil {
		return err
	}

	return nil
}
