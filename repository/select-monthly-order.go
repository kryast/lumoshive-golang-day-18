package repository

import "lumoshive-golang-day-18/model"

func (r *OrdersRepositoryDB) GetMontlyOrder() ([]model.MontlyOrder, error) {
	query := `
		SELECT 
			TO_CHAR(o.created_at, 'Month') AS month,
			COUNT(o.id) as total_order
		FROM
			order_process op

		FULL JOIN
			orders o ON op.order_id = o.id

		GROUP BY
			month;`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var MontlyOrder []model.MontlyOrder
	for rows.Next() {
		var order model.MontlyOrder
		if err := rows.Scan(&order.Month, &order.TotalOrder); err != nil {
			return nil, err
		}
		MontlyOrder = append(MontlyOrder, order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return MontlyOrder, nil
}
