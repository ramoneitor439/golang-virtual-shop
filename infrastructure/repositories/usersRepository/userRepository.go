package usersRepository

import (
	"errors"

	"mystore.com/domain/entities"
	"mystore.com/infrastructure/data"
	userRolesRepository "mystore.com/infrastructure/repositories/userRolesRepository"
)

func Add(user *entities.User) error {
	connection, conErr := data.CreatePostgresqlConnection()
	if conErr != nil {
		return conErr
	}

	connectionError := connection.Connect()
	if connectionError != nil {
		return connectionError
	}

	defer connection.Close()

	connection.StartTransaction()

	_, execErr := connection.Mute("INSERT INTO public.users (email, first_name, last_name, password_hash, is_active) "+
		"VALUES ($1, $2, $3, $4, $5)", user.Email, user.FirstName, user.LastName, user.PasswordHash, user.IsActive)

	if execErr != nil {
		connection.RollbackTransaction()
		return execErr
	}

	rows, queryErr := connection.Query("SELECT id FROM public.users WHERE email = $1", user.Email)
	if queryErr != nil {
		connection.RollbackTransaction()
		return queryErr
	}

	if !rows.Next() {
		connection.RollbackTransaction()
		return errors.New("user was not correctly registered")
	}

	var userId uint64
	if scanErr := rows.Scan(&userId); scanErr != nil {
		connection.RollbackTransaction()
		return scanErr
	}

	userRolesErr := userRolesRepository.Add(userId, user.Roles...)
	if userRolesErr != nil {
		connection.RollbackTransaction()
		return userRolesErr
	}

	connection.CommitTransaction()

	return nil
}

func GetAll() ([]*entities.User, error) {
	connection, conErr := data.CreatePostgresqlConnection()
	if conErr != nil {
		return nil, conErr
	}

	connectionError := connection.Connect()
	if connectionError != nil {
		return nil, connectionError
	}

	defer connection.Close()

	rows, queryErr := connection.Query("SELECT u.id, u.email, u.first_name, u.last_name, u.password, u.is_active, r.id, r.name, r.normalized_name " +
		"FROM public.users u " +
		"INNER JOIN user_roles ur " +
		"ON ur.user_id = u.id " +
		"INNER JOIN roles r " +
		"ON r.id = ur.role_id " +
		"GROUP BY u.id")

	if queryErr != nil {
		return nil, queryErr
	}

	var usersMap = make(map[uint64]*entities.User)

	for rows.Next() {
		var user entities.User
		var role entities.Role
		err := rows.Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName, &user.PasswordHash, &user.IsActive, &role.Id, &role.Name, &role.NormalizedName)
		if err != nil {
			return nil, err
		}

		if _, exists := usersMap[user.Id]; exists {
			usersMap[user.Id].Roles = append(usersMap[user.Id].Roles, &role)
			continue
		}

		usersMap[user.Id] = &entities.User{
			BaseEntity:      entities.BaseEntity{Id: user.Id},
			AuditableEntity: entities.AuditableEntity{},
			Email:           user.Email,
			PasswordHash:    user.PasswordHash,
			FirstName:       user.FirstName,
			LastName:        user.LastName,
			IsActive:        user.IsActive,
			Roles:           []*entities.Role{&role},
		}
	}

	var users []*entities.User

	for _, user := range usersMap {
		users = append(users, user)
	}

	return users, nil
}

func GetByEmail(email string) (*entities.User, error) {
	connection, conErr := data.CreatePostgresqlConnection()
	if conErr != nil {
		return nil, conErr
	}

	connectionError := connection.Connect()
	if connectionError != nil {
		return nil, connectionError
	}

	defer connection.Close()

	rows, queryErr := connection.Query("SELECT u.id, u.email, u.first_name, u.last_name, u.password_hash, u.is_active, r.id, r.name, r.normalized_name "+
		"FROM public.users u "+
		"INNER JOIN public.user_roles ur "+
		"ON ur.user_id = u.id "+
		"INNER JOIN public.roles r "+
		"ON r.id = ur.role_id "+
		"WHERE email = $1", email)

	if queryErr != nil {
		return nil, queryErr
	}

	userMap := make(map[uint64]*entities.User)
	var user *entities.User

	for rows.Next() {
		var user entities.User
		var role entities.Role
		err := rows.Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName, &user.PasswordHash, &user.IsActive, &role.Id, &role.Name, &role.NormalizedName)
		if err != nil {
			return nil, err
		}

		if _, exists := userMap[user.Id]; exists {
			userMap[user.Id].Roles = append(userMap[user.Id].Roles, &role)
			continue
		}

		userMap[user.Id] = &entities.User{
			BaseEntity:      entities.BaseEntity{Id: user.Id},
			AuditableEntity: entities.AuditableEntity{},
			Email:           user.Email,
			PasswordHash:    user.PasswordHash,
			FirstName:       user.FirstName,
			LastName:        user.LastName,
			IsActive:        user.IsActive,
			Roles:           []*entities.Role{&role},
		}
	}

	for _, value := range userMap {
		user = value
		break
	}

	return user, nil
}

func GetById(id uint64) (*entities.User, error) {
	connection, conErr := data.CreatePostgresqlConnection()
	if conErr != nil {
		return nil, conErr
	}

	connectionError := connection.Connect()
	if connectionError != nil {
		return nil, connectionError
	}

	defer connection.Close()

	rows, queryErr := connection.Query("SELECT id, email, first_name, last_name, password, is_active, r.id, r.name, r.normalized_name "+
		"FROM public.users u "+
		"INNER JOIN public.user_roles ur "+
		"ON ur.user_id = u.id "+
		"INNER JOIN public.roles r "+
		"ON r.id = ur.role_id "+
		"WHERE id = $1", id)

	if queryErr != nil {
		return nil, queryErr
	}

	userMap := make(map[uint64]*entities.User)
	var user *entities.User

	for rows.Next() {
		var user entities.User
		var role entities.Role
		err := rows.Scan(&user.Id, &user.Email, &user.FirstName, &user.LastName, &user.PasswordHash, &user.IsActive, &role.Id, &role.Name, &role.NormalizedName)
		if err != nil {
			return nil, err
		}

		if _, exists := userMap[user.Id]; exists {
			userMap[user.Id].Roles = append(userMap[user.Id].Roles, &role)
			continue
		}

		userMap[user.Id] = &entities.User{
			BaseEntity:      entities.BaseEntity{Id: user.Id},
			AuditableEntity: entities.AuditableEntity{},
			Email:           user.Email,
			PasswordHash:    user.PasswordHash,
			FirstName:       user.FirstName,
			LastName:        user.LastName,
			IsActive:        user.IsActive,
			Roles:           []*entities.Role{&role},
		}
	}

	for _, value := range userMap {
		user = value
		break
	}

	return user, nil
}

func Delete(id uint64) error {
	connection, conErr := data.CreatePostgresqlConnection()
	if conErr != nil {
		return conErr
	}

	connectionError := connection.Connect()
	if connectionError != nil {
		return connectionError
	}

	defer connection.Close()

	_, execErr := connection.Mute("DELETE FROM public.users WHERE id = $1", id)
	if execErr != nil {
		return execErr
	}

	return nil
}

func Update(user *entities.User) error {
	connection, conErr := data.CreatePostgresqlConnection()
	if conErr != nil {
		return conErr
	}

	connectionError := connection.Connect()
	if connectionError != nil {
		return connectionError
	}

	defer connection.Close()

	_, execErr := connection.Mute("UPDATE public.users SET "+
		"email = $1,"+
		"first_name = $2,"+
		"last_name = $3 "+
		"WHERE id = $4", user.Email, user.FirstName, user.LastName, user.Id)
	if execErr != nil {
		return execErr
	}

	return nil
}

func Exists(email string) error {
	connection, conErr := data.CreatePostgresqlConnection()
	if conErr != nil {
		return conErr
	}

	connectionError := connection.Connect()
	if connectionError != nil {
		return connectionError
	}

	defer connection.Close()

	rows, queryErr := connection.Query("SELECT 1 FROM public.users WHERE email = $1", email)
	if queryErr != nil {
		return queryErr
	}

	if rows.Next() {
		return errors.New("email is already taken")
	}

	return nil
}
