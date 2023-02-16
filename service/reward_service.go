package service

import (
	"dcom_service/repository"
	"log"
)

type rewardService struct {
	rewardRepo repository.RewardRepository
	logger     *log.Logger
}

func NewRewardService(rewardRepo repository.RewardRepository, logger *log.Logger) RewardService {
	return rewardService{rewardRepo: rewardRepo, logger: logger}
}

func (s rewardService) GetAllReward() ([]RewardResponse, error) {
	rewards, err := s.rewardRepo.GetAllReward()
	if err != nil {
		s.logger.Printf("Error : %v", err)
		return nil, err
	}

	rewardResponses := []RewardResponse{}

	for _, reward := range rewards {
		rewardResponse := RewardResponse{
			CreateDate:    reward.CreateDate,
			RewardName:    reward.RewardName,
			Total:         reward.Total,
			CustomerName:  reward.CustomerName,
			Email:         reward.Email,
			PhoneNumber:   reward.PhoneNumber,
			AddressDetail: reward.AddressDetail,
			SubDistrict:   reward.SubDistrict,
			District:      reward.District,
			Province:      reward.Province,
			PostalCode:    reward.PostalCode,
			CN:            reward.CN,
		}
		rewardResponses = append(rewardResponses, rewardResponse)
	}

	return rewardResponses, nil
}
