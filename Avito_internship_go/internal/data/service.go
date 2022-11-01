package data

import (
	"errors"
)

func (s ServiceModel) Get(id int64) (*Service, error) {
	print("get service")
	if id < 1 {
		return nil, errors.New("incorrect id")
	}

	query := `
			SELECT *
			FROM service
			WHERE service_id = $1
		`

	var service Service
	err := s.DB.QueryRow(query, id).Scan(
		&service.ServiceId,
		&service.ServiceName,
		&service.ServicePrice,
	)

	if err != nil {
		println("service dont found")
	}

	return &service, nil
}

func (s ServiceModel) Create(name string, price float64) {
	query := `
		INSERT INTO service (service_name, service_price)
		VALUES ($1, $2)`

	s.DB.QueryRow(query, name, price).Scan()
}
