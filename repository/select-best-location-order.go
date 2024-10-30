package repository

import "lumoshive-golang-day-18/model"

func (r *OrdersRepositoryDB) GetBestLocationOrder() ([]model.BestLocationOrder, error) {
	query := `
		SELECT
			location,
			COUNT(*) AS total_orders
		FROM (
			SELECT
				pick_up AS location
			FROM
				orders
			UNION ALL
			SELECT
				destination AS location
			FROM
				orders
			) AS combined_locations
		GROUP BY
			location
		ORDER BY
			total_orders DESC;`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var BestLocation []model.BestLocationOrder
	for rows.Next() {
		var location model.BestLocationOrder
		if err := rows.Scan(&location.Location, &location.TotalOrders); err != nil {
			return nil, err
		}
		BestLocation = append(BestLocation, location)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return BestLocation, nil
}
