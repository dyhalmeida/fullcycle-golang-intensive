package database

import (
	"database/sql"

	"github.com/dyhalmeida/fullcycle-golang-intensive/internal/order/entity"
	_ "github.com/go-sql-driver/mysql"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, errorPrepareStmt := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")

	if errorPrepareStmt != nil {
		return errorPrepareStmt
	}

	_, errorExec := stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)

	if errorExec != nil {
		return errorExec
	}
	return nil
}
