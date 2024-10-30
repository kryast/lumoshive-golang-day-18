package repository

import "lumoshive-golang-day-18/model"

func (r *OrdersRepositoryDB) GetMonthlyCustomer() ([]model.MonthlyCustomer, error) {
	query := `
		SELECT
			TO_CHAR(o.created_at, 'Month') AS month,
			COUNT(o.id) AS total_order,
			c.name AS customer_name
		FROM
			orders o
		JOIN
			customer c ON o.customer_id = c.id
		LEFT JOIN
			order_process op ON o.id = op.order_id
		GROUP BY
			month, c.name
		ORDER BY
			total_order DESC
		LIMIT 1;`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var MonthlyCustomer []model.MonthlyCustomer
	for rows.Next() {
		var customer model.MonthlyCustomer
		if err := rows.Scan(&customer.Month, &customer.TotalOrder, &customer.CustomerName); err != nil {
			return nil, err
		}
		MonthlyCustomer = append(MonthlyCustomer, customer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return MonthlyCustomer, nil
}
