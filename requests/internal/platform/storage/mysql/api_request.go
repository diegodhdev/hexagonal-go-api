package mysql

const (
	sqlApiRequestTable = "requests"
)

type sqlApiRequest struct {
	ID           string `db:"id"`
	Api          string `db:"api"`
	Mode         string `db:"mode"`
	ResponseType string `db:"response_type"`
	Request      string `db:"body"`
}
