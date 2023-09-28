package mysql

const (
	sqlApiRequestTable = "requests"
)

type sqlApiRequest struct {
	ID string `db:"id"`
}
