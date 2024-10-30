package repository

import "lumoshive-golang-day-18/model"

func (r *DriverRepositoryDB) GetBestDriver() ([]model.BestDriver, error) {
	query := `
		SELECT
    		d.name AS driver_name,
    	COUNT(op.order_id) AS total_orders
		FROM
    		order_process op
		JOIN
    		driver d ON op.driver_id = d.id
		GROUP BY
    		d.name
		ORDER BY
    		total_orders DESC
		LIMIT 1;`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var BestDriver []model.BestDriver
	for rows.Next() {
		var driver model.BestDriver
		if err := rows.Scan(&driver.DriverName, &driver.TotalOrders); err != nil {
			return nil, err
		}
		BestDriver = append(BestDriver, driver)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return BestDriver, nil
}
