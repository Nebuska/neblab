package board

import (
	"github.com/Nebuska/neblab/shared/appError"
	"github.com/Nebuska/neblab/tasker/internal/boardUser"

	"gorm.io/gorm"
)

type Repository interface {
	UserHasAccessToBoard(userID, boardID uint) (bool, error)
	GetBoard(boardId uint) (Board, error)
	GetUsersBoards(userID uint) ([]Board, error)
	CreateBoard(OwnerId uint, board Board) (Board, error)
	UpdateBoard(board Board) (Board, error)
	DeleteBoard(boardID uint) error
}

type boardRepository struct {
	DB *gorm.DB
}

func newBoardRepository(db *gorm.DB) Repository {
	return &boardRepository{db}
}

func (repo *boardRepository) UserHasAccessToBoard(userID, boardID uint) (bool, error) {
	var count int64
	err := repo.DB.Table("board_users").
		Where("board_id = ? AND user_id = ?", boardID, userID).
		Count(&count).Error
	return count > 0, appError.FromGormError(err)
}

func (repo *boardRepository) GetBoard(id uint) (Board, error) {
	var board Board
	err := repo.DB.Preload("Tasks").First(&board, id).Error
	return board, appError.FromGormError(err)
}

func (repo *boardRepository) GetUsersBoards(userId uint) ([]Board, error) {
	var boards []Board
	err := repo.DB.Preload("Tasks").
		Joins("JOIN board_users ON board_users.board_id = boards.id").
		Where("board_users.user_id = ?", userId).Find(&boards).Error
	return boards, appError.FromGormError(err)
}

func (repo *boardRepository) CreateBoard(OwnerId uint, board Board) (Board, error) {
	err := repo.DB.Create(&board).Error
	if err != nil {
		repo.DB.Delete(&board) //todo : it might not exist and it should hard delete
		return Board{}, appError.FromGormError(err)
	}
	err = repo.DB.Create(&boardUser.BoardUser{
		UserID:  OwnerId,
		BoardID: board.ID,
		Role:    "",
	}).Error
	return board, appError.FromGormError(err)
}

func (repo *boardRepository) UpdateBoard(board Board) (Board, error) {
	err := repo.DB.Save(&board).Error
	return board, appError.FromGormError(err)
}

func (repo *boardRepository) DeleteBoard(boardID uint) error {
	err := repo.DB.Delete(&Board{}, boardID).Error
	return appError.FromGormError(err)
}
