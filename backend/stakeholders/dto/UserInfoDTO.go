package dto


type UserInfoDTO struct {
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	ProfilePicture string `json:"profilePicture"`
	Biography      string `json:"biography"`
	Motto          string `json:"motto"`
}
