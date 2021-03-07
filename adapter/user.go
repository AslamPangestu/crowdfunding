package adapter

import "crowdfunding/entity"

//RegsiterAdapter : Adapter Register
func RegsiterAdapter(user entity.User, token string) entity.RegisterResponse {
	return entity.RegisterResponse{
		ID:         user.ID,
		Name:       user.Name,
		Username:   user.Username,
		Email:      user.Email,
		Occupation: user.Occupation,
		Token:      token,
	}
}

//LoginAdapter : Adapter Register
func LoginAdapter(user entity.User, token string) entity.LoginResponse {
	return entity.LoginResponse{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}
}