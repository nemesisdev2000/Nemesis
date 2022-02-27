package UserDB

type user struct {
	username     string
	passwordhash string
	role         int
}

var userList = []user{
	{
		username:     "admin",
		passwordhash: "passw0rd",
		role:         1,
	},

	{
		username:     "dosxuz",
		passwordhash: "testpass",
		role:         0,
	},
}

func GetUserObject(username string) (user, bool) {
	for _, user := range userList {
		if user.username == username {
			return user, true
		}
	}
	return user{}, false
}

func (u *user) ValidatePasswordHash(pswdhash string) bool {
	return u.passwordhash == pswdhash
}

func AddUserObject(username string, passwordhash string, role int) bool {
	newUser := user{
		username:     username,
		passwordhash: passwordhash,
		role:         role,
	}

	//check if a user already exists

	for _, ele := range userList {
		if ele.username == username {
			return false
		}
	}
	userList = append(userList, newUser)
	return true
}
