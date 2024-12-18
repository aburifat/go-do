package storage

import (
	"errors"
	"sync"
	"time"

	"github.com/aburifat/go-do/models"
)

type MemoryStorage struct {
	todos map[string]*models.Todo
	mu    sync.RWMutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		todos: make(map[string]*models.Todo),
	}
}

func (s *MemoryStorage) GetTodos() []*models.Todo {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todos := make([]*models.Todo, 0, len(s.todos))

	for _, todo := range s.todos {
		todos = append(todos, todo)
	}

	return todos
}

func (s *MemoryStorage) GetTodoByID(id string) (*models.Todo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todo, exists := s.todos[id]

	if !exists {
		return nil, errors.New("todo not found")
	}
	return todo, nil
}

func (s *MemoryStorage) CreateTodo(todo *models.Todo) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo.ID = generateID()
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	s.todos[todo.ID] = todo
}

func (s *MemoryStorage) UpdateTodo(id string, model *models.Todo) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo, exists := s.todos[id]

	if !exists {
		return errors.New("no record found with the id")
	} else if id != model.ID {
		return errors.New("corrapted data found on request")
	}

	todo.Title = model.Title
	todo.Completed = model.Completed
	todo.UpdatedAt = time.Now()
	s.todos[id] = todo

	return nil
}

func (s *MemoryStorage) DeleteTodo(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.todos[id]; !exists {
		return errors.New("todo not found")
	}
	delete(s.todos, id)
	return nil
}

func generateID() string {
	return time.Now().Format("20060102150405")
}
