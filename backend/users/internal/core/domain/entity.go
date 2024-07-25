package domain


const (
	Admin = "admin"
	Customer = "user"
)




type (
	User struct {
		ID UserID
		Username string
		Password string
		Role string
		CreatedAt  time.Time
	}

	UserID struct {
		uuid.UUID
	}

	func ParseUserID(id string) (UserID, error) {
		if id == "" {
			return UserID{}, errors.New("id is empty string")
		}
		uid, err := uuid.Parse(id)
		if err!=nil {
			return UserID{}, err
		}
		if uid == uuid.Nil  {
			return UserID{}, errors.New("id is zero uuid")
		}
		return UserID{uid}, nil
	}
	func GenerateUserID(id string) (UserID, error) {
		id, err := uuid.NewRandom()
		if err!= nil {
			return UserID{}, fmt.Errorf("failed to generate vault id: %s", err.Error())
		}
		return UserID{id}, err
	}

	func (u UserID) IsZero() bool {
		return u.String() == uuid.Nil.String()
	}



	AccessToken struct {
		UserID UserID
		JWTToken AccessJWTToken
		ExpiresAt time.Time
		IssuedAt time.Time
		

	}
	AccessJWTToken struct {
		UserID UserID
		Role string
		Value string
	}

	RefreshtToken struct {
		ID UUID
		UserID UserID
		ExpiresAt time.Time
		IssuedAt time.Time
	}

	RefreshtTokenID struct {
		uuid.UUID
	}

	func ParseRefreshtTokenID(id string) (RefreshtTokenID, error) {
		if id == "" {
			return RefreshtTokenID{}, errors.New("id is empty string")
		}
		uid, err := uuid.Parse(id)
		if err!=nil {
			return RefreshtTokenID{}, err
		}
		if uid == uuid.Nil  {
			return RefreshtTokenID{}, errors.New("id is zero uuid")
		}
		return RefreshtTokenID{uid}, nil
	}
	func GenerateRefreshtTokenID(id string) (RefreshtTokenID, error) {
		id, err := uuid.NewRandom()
		if err!= nil {
			return RefreshtTokenID{}, fmt.Errorf("failed to generate vault id: %s", err.Error())
		}
		return RefreshtTokenID{id}, err
	}

	func (u RefreshtTokenID) IsZero() bool {
		return u.String() == uuid.Nil.String()
	}
	
)