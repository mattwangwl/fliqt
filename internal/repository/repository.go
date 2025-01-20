package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fliqt/internal/database"
	"fliqt/internal/model"
	"fliqt/internal/rediscli"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

func New() *Repository {
	return &Repository{
		db:       database.New(),
		redisCli: rediscli.New(),
	}
}

type Repository struct {
	db       *database.Database
	redisCli *rediscli.RedisClient
}

func (repo *Repository) Employee() IEmployee {
	return &employee{
		repository: repo,
	}
}

type IEmployee interface {
	Create(ctx context.Context, models ...model.Employee) error
	Update(ctx context.Context, cond model.EmployeeUpdateCond) error
	Get(ctx context.Context, id int64) (*model.Employee, error)
	List(ctx context.Context, cond model.EmployeeListCond) ([]*model.Employee, error)
}

type employee struct {
	repository *Repository
	model      model.Employee
}

func (dao *employee) Create(ctx context.Context, models ...model.Employee) error {
	if err := dao.repository.db.Create(models); err != nil {
		return err.Error
	}
	return nil
}

func (dao *employee) Update(ctx context.Context, cond model.EmployeeUpdateCond) error {
	updates := make(map[string]interface{})
	if cond.Name != nil {
		updates["name"] = cond.Name
	}
	if cond.Position != nil {
		updates["position"] = cond.Position
	}
	if cond.Status != nil {
		updates["status"] = cond.Status
	}

	if err := dao.repository.db.Model(dao.model).
		Where("`id` = ?", cond.ID).
		Updates(updates).Error; err != nil {
		return err
	}

	if _, err := dao.repository.redisCli.Del(ctx, fmt.Sprintf("employee:%d", cond.ID)).Result(); err != nil {
		return err
	}

	return nil
}

func (dao *employee) Get(ctx context.Context, id int64) (*model.Employee, error) {
	var (
		model model.Employee
	)

	if val, err := dao.repository.redisCli.Get(ctx, fmt.Sprintf("employee:%d", id)).Result(); err != nil {
		if !errors.Is(err, redis.Nil) {
			log.Printf("failed to get employee model from redis: %v", err)
		}
	} else {
		if err = json.Unmarshal([]byte(val), &model); err != nil {
			return nil, err
		} else {
			return &model, nil
		}
	}

	if err := dao.repository.db.Model(dao.model).
		Where("`id` = ?", id).First(&model).Error; err != nil {
		return nil, err
	}
	func() {
		buf, err := json.Marshal(model)
		if err != nil {
			log.Printf("failed to marshal employee model: %v", err)
		}
		if _, err := dao.repository.redisCli.Set(ctx, fmt.Sprintf("employee:%d", id), buf, 30*time.Second).Result(); err != nil {
			log.Printf("failed to set employee model to redis: %v", err)
		}
	}()

	return &model, nil
}

func (dao *employee) List(ctx context.Context, cond model.EmployeeListCond) ([]*model.Employee, error) {
	db := dao.repository.db.Model(dao.model)

	if cond.Name != nil {
		db = db.Where("`name` = ?", cond.Name)
	}
	if len(cond.Position) > 0 {
		db = db.Where("`position` in ?", cond.Position)
	}
	if cond.Status != nil {
		db = db.Where("`status` = ?", cond.Status)
	}

	var result []*model.Employee
	if err := db.Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
