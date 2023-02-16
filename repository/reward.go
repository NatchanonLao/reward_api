package repository

type Reward struct {
	CreateDate    string
	RewardName    string
	Total         int
	CustomerName  string
	Email         string
	PhoneNumber   string
	AddressDetail string
	SubDistrict   string
	District      string
	Province      string
	PostalCode    string
	CN            string
}

type RewardRepository interface {
	GetAllReward() ([]Reward, error)
}
