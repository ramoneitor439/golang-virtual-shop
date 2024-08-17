package userrolesrepository

import (
	"fmt"

	"mystore.com/domain/entities"
	"mystore.com/infrastructure/data"
)

func Add(userId uint64, roles ...*entities.Role) error {
	connection, conErr := data.CreatePostgresqlConnection()
	if conErr != nil {
		return conErr
	}

	connectionError := connection.Connect()
	if connectionError != nil {
		return connectionError
	}

	defer connection.Close()

	var sqlInsertion = "INSERT INTO public.user_roles (user_id, role_id) VALUES "

	for index, role := range roles {
		sqlInsertion += fmt.Sprintf("(%d, %d)", userId, role.Id)
		if index < len(roles)-1 {
			sqlInsertion += ","
		}
	}

	_, insertError := connection.Mute(sqlInsertion)

	return insertError
}
