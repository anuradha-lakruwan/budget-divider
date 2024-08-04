package user

import (
	"net/http"

	"github.com/rbretecher/expense-tracker/internal/service"
)

type DeleteArgs struct {
	ID int `json:"id" validate:"required"`
}

func (s *UserService) Delete(r *http.Request, args *DeleteArgs, reply *service.NoReply) error {
	result, err := s.db.Exec(`
		DELETE
		FROM users
		WHERE id = $1
	`, args.ID)

	return service.HandleDelete(result, err)
}
