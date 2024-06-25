package usermodel

import "time"

type User struct {
	Id                     int
	AppId                  int
	ChannelId              int
	ChannelUserId          string
	Name                   string
	Nickname               string
	BirthDate              string
	Location               string
	Gender                 string
	City                   string
	Email                  string
	EmailVerifiedAt        time.Time
	Phone                  string
	PhoneVerificationToken time.Time
	Password               string
	OldPassword            string
	Picture                string
	Status                 string
	ResetPasswordToken     string
	FcmToken               string
	RememberToken          string
	CreatedAt              time.Time
	UpdatedAt              time.Time
	DeletedAt              time.Time
	Client                 string
	Pin                    string
	IsEnabled              bool
	IpAddress              string
	Lang                   string
	IntegrationId          int
	ContactId              int
	Fields                 string
	IsLogin                bool
	CustomerTags           string
	BranchOffice           string
	CreditAppNumber        string
	JobTitle               string
	Address                string
	Company                string
	JobRole                string
	Website                string
	CompanyNumber          string
	JobOffice              string
	BlockDescription       string
	BlockReportedBy        string
	IsBlockApproved        bool
	BlockApprovedBy        int
	RequestBlockAt         time.Time
	ApproveBlockAt         time.Time
	TotalFollowUpRetries   int
	TempId                 string
	ResetPinToken          string
}

func (User) TableName() string {
	return "app.users"
}
