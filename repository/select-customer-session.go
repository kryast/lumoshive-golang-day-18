package repository

import "lumoshive-golang-day-18/model"

func (r *SessionRepositoryDB) GetCustomerSession() ([]model.CustomerSession, error) {
	query := `
		SELECT
			CASE
				WHEN is_active THEN 'Logged In'
				ELSE 'Logged Out'
			END AS status,
			COUNT(*) AS total_customers
		FROM
			session
		GROUP BY
			is_active;`

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var CustomerSession []model.CustomerSession
	for rows.Next() {
		var status model.CustomerSession
		if err := rows.Scan(&status.Status, &status.TotalCustomers); err != nil {
			return nil, err
		}
		CustomerSession = append(CustomerSession, status)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return CustomerSession, nil
}
