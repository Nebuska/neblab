package board

import (
	"testing"

	"github.com/Nebuska/task-tracker/internal/boardUser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestBoardService_CreateBoard(t *testing.T) {
	var test = struct {
		userId uint
		board  Board
	}{
		userId: 1,
		board: Board{
			Name: "Test Board",
		},
	}

	repo := new(mockBoardRepository)
	repo.On("CreateBoard", test.userId, test.board).Return(test.board, nil)

	service := NewBoardService(repo)

	board, err := service.CreateBoard(test.userId, test.board)
	assert.NoError(t, err)
	assert.Equal(t, test.board, board)
}

func TestBoardService_GetBoard(t *testing.T) {
	var test = struct {
		userId uint
		board  Board
	}{
		userId: 1,
		board: Board{
			Model: gorm.Model{ID: 1},
			Name:  "Test Board",
			BoardUsers: []boardUser.BoardUser{
				{
					UserID:  1,
					BoardID: 1,
					Role:    "Creator",
				},
			},
		},
	}

	repo := new(mockBoardRepository)
	repo.On("GetBoard", test.userId, test.board).Return(test.board, nil)

	service := NewBoardService(repo)
	board, err := service.GetBoard(test.userId, test.board.ID)
	assert.NoError(t, err)
	assert.Equal(t, test.board, board)
}

func TestBoardService_GetBoards(t *testing.T) {
	var test = struct {
		userId uint
		boards []Board
	}{
		userId: 1,
		boards: []Board{},
	}

	repo := new(mockBoardRepository)
	repo.On("GetBoards", test.userId).Return(test.boards, nil)

	service := NewBoardService(repo)
	boards, err := service.GetBoards(test.userId)
	assert.NoError(t, err)
	assert.Equal(t, test.boards, boards)
}

func TestBoardService_DeleteBoard(t *testing.T) {
	var test = struct {
		boardId     uint
		userId      uint
		expectError bool
	}{
		boardId:     1,
		userId:      1,
		expectError: false,
	}

	repo := new(mockBoardRepository)
	repo.On("DeleteBoard", test.userId, test.boardId).Return(test.expectError).Once()

	service := NewBoardService(repo)
	err := service.DeleteBoard(test.userId, test.boardId)

	assert.Equal(t, test.expectError, err != nil)
}

type mockBoardRepository struct {
	Repository
	mock.Mock
}

func (m *mockBoardRepository) UserHasAccessToBoard(userID, boardID uint) (bool, error) {
	if userID == 1 {
		return true, nil
	}
	return userID == boardID, nil
}

func (m *mockBoardRepository) GetBoard(boardId uint) (Board, error) {
	args := m.Called(boardId)
	return args.Get(0).(Board), args.Error(1)
}

func (m *mockBoardRepository) GetUsersBoards(userID uint) ([]Board, error) {
	args := m.Called(userID)
	return args.Get(0).([]Board), args.Error(1)
}

func (m *mockBoardRepository) CreateBoard(OwnerId uint, board Board) (Board, error) {
	args := m.Called(OwnerId, board)
	return args.Get(0).(Board), args.Error(1)
}

func (m *mockBoardRepository) UpdateBoard(board Board) (Board, error) {
	args := m.Called(board)
	return args.Get(0).(Board), args.Error(1)
}

func (m *mockBoardRepository) DeleteBoard(boardID uint) error {
	args := m.Called(boardID)
	return args.Error(0)
}
