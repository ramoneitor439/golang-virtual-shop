package auth

import (
	"errors"

	"mystore.com/domain/constants/roles"
	"mystore.com/domain/entities"
	authDtos "mystore.com/dtos/authDtos"
	"mystore.com/infrastructure/repositories/rolesRepository"
	"mystore.com/infrastructure/repositories/usersRepository"
	jwtbuilder "mystore.com/infrastructure/security/jwtBuilder"
)

func SignUp(email string, firstName string, lastName string, password string) error {
	if alreadyExists := usersRepository.Exists(email); alreadyExists != nil {
		return alreadyExists
	}

	roles, rolesErr := rolesRepository.GetByName(roles.USER)
	if rolesErr != nil {
		return rolesErr
	}

	user := &entities.User{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		IsActive:  false,
		Roles:     roles,
	}

	user.SetPassword(password)

	addError := usersRepository.Add(user)
	if addError != nil {
		return addError
	}

	return nil
}

func SignIn(email string, password string) (*authDtos.SignInResponse, error) {
	user, userError := usersRepository.GetByEmail(email)
	if userError != nil {
		return nil, userError
	}

	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	if !user.CheckPassword(password) {
		return nil, errors.New("invalid credentials")
	}

	response, err := jwtbuilder.CreateAccessToken(user)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func Update(userId uint64, request authDtos.UpdateRequest) error {
	user := &entities.User{
		BaseEntity: entities.BaseEntity{Id: userId},
		Email:      request.Email,
		FirstName:  request.FirstName,
		LastName:   request.LastName,
	}

	err := usersRepository.Update(user)

	return err
}
