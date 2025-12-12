package board

import (
	"github.com/Nebuska/neblab/shared/appError"
	"github.com/Nebuska/neblab/shared/appError/errorCodes"
)

type Service interface {
	CreateBoard(userId uint, board Board) (Board, error)
	GetBoard(userId uint, boardId uint) (Board, error)
	GetBoards(userId uint) ([]Board, error)
	DeleteBoard(userId uint, boardId uint) error
}

type boardService struct {
	repo Repository
}

func NewBoardService(repo Repository) Service {
	return &boardService{repo: repo}
}

func (s *boardService) CreateBoard(userId uint, board Board) (Board, error) {
	return s.repo.CreateBoard(userId, board)
}

func (s *boardService) GetBoard(userId uint, boardId uint) (Board, error) {
	hasAccess, err := s.repo.UserHasAccessToBoard(userId, boardId)
	if err != nil {
		return Board{}, err
	}
	if !hasAccess {
		return Board{}, appError.New(errorCodes.Forbidden,
			"BoardService",
			"user does not have access to board")
	}
	return s.repo.GetBoard(boardId)
}

func (s *boardService) GetBoards(userId uint) ([]Board, error) {
	return s.repo.GetUsersBoards(userId)
}

func (s *boardService) DeleteBoard(userId uint, boardId uint) error {
	hasAccess, err := s.repo.UserHasAccessToBoard(userId, boardId)
	if err != nil {
		return err
	}
	if !hasAccess {
		return appError.New(errorCodes.Forbidden,
			"BoardService",
			"user does not have access to board")
	}
	return s.repo.DeleteBoard(boardId)
}
