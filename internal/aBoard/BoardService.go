package aBoard

import (
	"github.com/Nebuska/task-tracker/pkg/appError"
	"github.com/Nebuska/task-tracker/pkg/appError/errorCodes"
)

type Service interface {
	CreateBoardUsingUser(userId uint, board Board) (Board, error)
	GetBoardUsingUser(userId uint, boardId uint) (Board, error)
	GetBoardsUsingUser(userId uint) ([]Board, error)
	DeleteBoardUsingUser(userId uint, boardId uint) error
}

type boardService struct {
	repo Repository
}

func NewBoardService(repo Repository) Service {
	return &boardService{repo: repo}
}

func (s *boardService) CreateBoardUsingUser(userId uint, board Board) (Board, error) {
	return s.repo.CreateBoard(userId, board)
}

func (s *boardService) GetBoardUsingUser(userId uint, boardId uint) (Board, error) {
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

func (s *boardService) GetBoardsUsingUser(userId uint) ([]Board, error) {
	return s.repo.GetUsersBoards(userId)
}

func (s *boardService) DeleteBoardUsingUser(userId uint, boardId uint) error {
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

func (s *boardService) CreateBoard(board Board) (Board, error) {
	panic("implement me")
}

func (s *boardService) GetBoard(boardId uint) (Board, error) {
	panic("implement me")
}
