package task

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type mockTaskRepository struct {
	Repository
	mock.Mock
}

func (m *mockTaskRepository) CreateTask(task Task) (Task, error) {
	return task, nil
}

func (m *mockTaskRepository) GetTasksByFilter(filter Filter) ([]Task, error) {
	if filter.PageSize == 0 {
		return nil, errors.New("page size must be greater than 0")
	}
	if filter.SortBy == "" {
		return nil, errors.New("sort by must be set")
	}
	return nil, nil
}

func (m *mockTaskRepository) GetTaskById(taskId uint) (Task, error) {
	args := m.Called(taskId)
	return args.Get(0).(Task), args.Error(1)
}

func (m *mockTaskRepository) UserHasAccessToBoard(userID, boardID uint) (bool, error) {
	if userID == 1 {
		return true, nil
	}
	return userID == boardID, nil
}

func TestTaskService_CreateTask(t *testing.T) {
	tests := []struct {
		name      string
		Task      Task
		UserID    uint
		wantError bool
	}{
		{
			name: "CreateTask success on correct task",
			Task: Task{
				Name:        "Test 1",
				Description: "This is a test task",
				Status:      "",
				BoardID:     1,
			},
			UserID:    1,
			wantError: false,
		},
		{
			name: "CreateTask fails on trying with no board id",
			Task: Task{
				Name:        "Test 2",
				Description: "This is a test task",
				Status:      "",
			},
			UserID:    1,
			wantError: true,
		},
		{
			name: "CreateTask fails on trying with a user without permission",
			Task: Task{
				Name:        "Test 3",
				Description: "This is a test task",
				Status:      "",
				BoardID:     1,
			},
			UserID:    2,
			wantError: true,
		},
	}
	repo := new(mockTaskRepository)
	service := NewTaskService(repo)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, err := service.CreateTask(test.UserID, test.Task)
			if !test.wantError {
				assert.NoError(t, err, test.name)
			} else {
				assert.Error(t, err, test.name)
			}
		})
	}
}

func TestTaskService_GetTaskById(t *testing.T) {
	tests := []struct {
		name      string
		TaskId    uint
		UserID    uint
		wantError bool
	}{
		{
			name:      "GetTaskById success",
			TaskId:    1,
			UserID:    1,
			wantError: false,
		},
		{
			name:      "GetTaskById fails on user with no permission",
			TaskId:    2,
			UserID:    2,
			wantError: true,
		},
		{
			name:      "GetTaskById fails on trying with no task",
			TaskId:    3,
			UserID:    1,
			wantError: true,
		},
	}

	repo := new(mockTaskRepository)
	repo.On("GetTaskById", uint(3)).Return(Task{}, "no task found")
	repo.On("GetTaskById", mock.Anything).Return(correctTask, nil)

	service := NewTaskService(repo)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, err := service.GetTaskById(test.UserID, test.TaskId)
			if test.wantError {
				assert.Error(t, err, test.name)
			} else {
				assert.NoError(t, err, test.name)
			}
		})
	}
}

func TestTaskService_GetTasksByFilter(t *testing.T) {
	repo := new(mockTaskRepository)
	service := NewTaskService(repo)

	_, err := service.GetTasksByFilter(1, Filter{})
	assert.NoError(t, err, "GetTasksByFilter should not return an error")

}

var correctTask = Task{
	Model:       gorm.Model{ID: 1},
	Name:        "Test 1",
	Description: "This is a test task",
	Status:      "",
	BoardID:     1,
}
