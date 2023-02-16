package handler

import (
	"dcom_service/service"
	"encoding/json"
	"net/http"
)

type rewardHandler struct {
	rewardSrv service.RewardService
}

func NewRewardHandler(rewardSrv service.RewardService) rewardHandler {
	return rewardHandler{rewardSrv: rewardSrv}
}

func (h rewardHandler) GetAllReward(w http.ResponseWriter, r *http.Request) {
	rewards, err := h.rewardSrv.GetAllReward()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(rewards)
}
