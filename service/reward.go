package service

type RewardResponse struct {
	CreateDate    string `json:"create_date"`
	RewardName    string `json:"reward_name"`
	Total         int    `json:"total"`
	CustomerName  string `json:"customer_name"`
	Email         string `json:"email"`
	PhoneNumber   string `json:"phone_number"`
	AddressDetail string `json:"address_detail"`
	SubDistrict   string `json:"sub_district"`
	District      string `json:"district"`
	Province      string `json:"province"`
	PostalCode    string `json:"postal_code"`
	CN            string `json:"customer_id"`
}

type RewardService interface {
	GetAllReward() ([]RewardResponse, error)
}
