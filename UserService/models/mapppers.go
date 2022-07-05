package models

func (admin *Admin) ToDTO() UserDTO {
	return UserDTO{
		Email:     admin.Email,
		Username:  admin.Username,
		Firstname: admin.Firstname,
		Lastname:  admin.Lastname,
		Role:      string(admin.Role),
	}
}

func (admin *Driver) ToDTO() UserDTO {
	return UserDTO{
		Email:     admin.Email,
		Username:  admin.Username,
		Firstname: admin.Firstname,
		Lastname:  admin.Lastname,
		Role:      string(admin.Role),
	}
}

func (admin *Passenger) ToDTO() UserDTO {
	return UserDTO{
		Email:     admin.Email,
		Username:  admin.Username,
		Firstname: admin.Firstname,
		Lastname:  admin.Lastname,
		Role:      string(admin.Role),
	}
}
