package service

import (
	"context"
	"fliqt/internal/model"
	"testing"
)

func generateTestEmployee() *Employee {
	return &Employee{
		employeeDao: &testEmployeeDao{},
	}
}

type testEmployeeDao struct {
}

func (dao *testEmployeeDao) Create(ctx context.Context, models ...model.Employee) error {
	return nil
}

func (dao *testEmployeeDao) Update(ctx context.Context, cond model.EmployeeUpdateCond) error {
	return nil
}

func (dao *testEmployeeDao) Get(ctx context.Context, id int64) (*model.Employee, error) {
	return &model.Employee{
		ID:       1,
		Name:     "test",
		Position: "test",
		Status:   1,
	}, nil
}

func (dao *testEmployeeDao) List(ctx context.Context, cond model.EmployeeListCond) ([]*model.Employee, error) {
	return nil, nil
}

func TestEmployee_List(t *testing.T) {
	impl := generateTestEmployee()
	type args struct {
		cond model.EmployeeListCond
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "get all",
			args: args{
				cond: model.EmployeeListCond{
					Name:     nil,
					Position: nil,
					Status:   nil,
				},
			},
			wantErr: false,
		},
		{
			name: "test position",
			args: args{
				cond: model.EmployeeListCond{
					Name:     nil,
					Position: []string{},
					Status:   nil,
				},
			},
			wantErr: false,
		},
		{
			name: "test position",
			args: args{
				cond: model.EmployeeListCond{
					Name:     nil,
					Position: []string{"Manager", "Developer", "Designer"},
					Status:   nil,
				},
			},
			wantErr: false,
		},
		{
			name: "test position error",
			args: args{
				cond: model.EmployeeListCond{
					Name:     nil,
					Position: []string{"Project Manager"},
					Status:   nil,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := impl.List(context.Background(), tt.args.cond)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
