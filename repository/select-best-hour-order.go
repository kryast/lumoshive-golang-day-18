package repository

import "lumoshive-golang-day-18/model"

func (r *OrdersRepositoryDB) GetBestHourOrder() ([]model.BestHourOrder, error) {
	query := `
		SELECT
   			order_hour,
    		total_orders
		FROM (
    	SELECT
        	EXTRACT(HOUR FROM created_at) AS order_hour,
        	COUNT(*) AS total_orders
    	FROM
        	orders
    	GROUP BY
        	order_hour
		) AS hourly_orders
		ORDER BY
    		total_orders DESC
		LIMIT 1;`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var BestHour []model.BestHourOrder
	for rows.Next() {
		var hour model.BestHourOrder
		if err := rows.Scan(&hour.OrderHour, &hour.TotalOrders); err != nil {
			return nil, err
		}
		BestHour = append(BestHour, hour)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return BestHour, nil
}
