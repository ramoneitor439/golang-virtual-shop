package rolesRepository

import (
	"fmt"

	"mystore.com/domain/entities"
	"mystore.com/infrastructure/data"
)

func GetByName(normalizedNames ...string) ([]*entities.Role, error) {
	connection, conErr := data.CreatePostgresqlConnection()
	if conErr != nil {
		return nil, conErr
	}

	connectionError := connection.Connect()
	if connectionError != nil {
		return nil, connectionError
	}

	defer connection.Close()

	var sqlQuery = "SELECT id, name, normalized_name FROM public.roles WHERE normalized_name IN ("
	for index, name := range normalizedNames {
		sqlQuery += fmt.Sprintf("'%s')", name)

		if index < len(normalizedNames)-1 {
			sqlQuery += ","
		}
	}

	rows, queryError := connection.Query(sqlQuery)
	if queryError != nil {
		return nil, queryError
	}

	var roles []*entities.Role
	for rows.Next() {
		var role entities.Role
		scanErr := rows.Scan(&role.Id, &role.Name, &role.NormalizedName)
		if scanErr != nil {
			return nil, scanErr
		}

		roles = append(roles, &role)
	}

	return roles, nil
}
