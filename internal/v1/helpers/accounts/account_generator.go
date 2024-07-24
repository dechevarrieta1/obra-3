package accountshelpersv1

import (
	middlewaresv1 "github.com/dechevarrieta1/obra-3/internal/v1/middlewares"
	accountsmodelsv1 "github.com/dechevarrieta1/obra-3/internal/v1/models/accounts"
)

func GenerateAccountWithJWT(acc accountsmodelsv1.AccountUserRequest) (accountsmodelsv1.AccountUserResponse, error) {
	jwt, err := middlewaresv1.GenerateJWT(acc.Username)
	if err != nil {
		return accountsmodelsv1.AccountUserResponse{}, err
	}
	accountResponse := accountsmodelsv1.AccountUserResponse{
		Name:     acc.Name,
		LastName: acc.LastName,
		Role:     acc.Role,
		Username: acc.Username,
		Email:    acc.Email,
		JWT:      jwt,
	}
	return accountResponse, nil
}
