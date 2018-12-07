package service

type UserManager struct{
	Users map[string]string
}

func (um *UserManager) InitializeService(){
	um.Users = make(map[string]string)
}

/*
func (um *UserManager) AddUser(username, password string){

	var exists bool
	_,exists = um.Users[username]

	um.Users[username] = password
}*/