package domain

type (
	UserOutDTO struct {
		ID UserID
		Username string
		Role string
		CreatedAt  time.Time
	}

	UserCreateInDTO struct {
		Username string
		RawPassword string
		Role string
	}

	UserCreateOutDTO struct {
		UserOutDTO 
	}
	
	UserGetInDTO struct {
		UserID UserID
	}

	UserGetOutDTO struct {
		UserOutDTO
	}

	UserUpdateInDTO struct {
		UserID UserID
		Username string
		Role string
	}
		
	UserUpdateOutDTO struct {
		UserOutDTO
	}

	UserDeleteInDTO struct {
		UserID UserID
		RawPassword string
	}	
	
	UserDeleteOutDTO struct {
		status bool
	}
)



