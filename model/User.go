package model

type UserInfo struct {
	UserID   string
	UserPW   string
	UserName string
}

func NewUserInfo(id string, pw string, name string) *UserInfo {
	return &UserInfo{
		UserID:   id,
		UserPW:   pw,
		UserName: name,
	}
}
