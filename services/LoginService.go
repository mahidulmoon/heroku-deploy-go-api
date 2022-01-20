package services

type LoginService interface {
	LoginUser(email string, password string) bool
}
type loginInformation struct {
	email    string
	password string
}

func StaticLoginService() LoginService {
	return &loginInformation{
		email:    "mahidulmoon@gmail.com",
		password: "testing",
	}
}
func (info *loginInformation) LoginUser(email string, password string) bool {
	return info.email == email && info.password == password
}

// dynamic user
// func (info *loginInformation) LoginUser(email string, password string) bool {
// 	var model models.User
// 	result := model.Authenticate(email, password)
// 	//return info.email == email && info.password == password
// 	if result {
// 		return true
// 	} else {
// 		return false
// 	}
// }
