package model

type UserInfo struct {
	UserID   string
	UserPW   string
	UserName string
}

type UserLoginInfo struct {
	UserID   string
	UserName string
}

func GetUserInfo(id string, name string) *UserLoginInfo {
	return &UserLoginInfo{
		UserID:   id,
		UserName: name,
	}
}

func NewUserInfo(id string, pw string, name string) *UserInfo {
	return &UserInfo{
		UserID:   id,
		UserPW:   pw,
		UserName: name,
	}
}
