package routes

import (
	"github.com/aburifat/go-do/handlers"
	"github.com/aburifat/go-do/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter(storage *storage.MemoryStorage) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/todos", handlers.HandleGetTodos(storage))
	r.Post("/todos", handlers.HandleCreateTodo(storage))
	r.Get("/todos/{id}", handlers.HandleGetTodoByID(storage))
	r.Put("/todos/{id}", handlers.HandleUpdateTodo(storage))
	r.Delete("/todos/{id}", handlers.HandleDeleteTodo(storage))

	return r
}
