package mysql

const (
	sqlApiRequestTable = "api_requests"
)

type sqlApiRequest struct {
	ID string `db:"id"`
}
