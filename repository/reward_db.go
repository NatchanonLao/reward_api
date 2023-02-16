package repository

import (
	"context"
	"database/sql"
	"time"
)

type rewardRepositoryDB struct {
	db *sql.DB
}

func NewRewardRepositoryDB(db *sql.DB) RewardRepository {
	return rewardRepositoryDB{db: db}
}

func (repo rewardRepositoryDB) GetAllReward() ([]Reward, error) {

	query := `SELECT createDate,rewardName,total,customerName,Email,phoneNumber,COALESCE(AddressDetail,'') as AddressDetail,
	COALESCE(SubDistrict,'') as SubDistrict,COALESCE(District,'') as District ,COALESCE(Province,'') as 
	Province, COALESCE(PostalCode,'') as PostalCode ,CN FROM BWAV_Reward_ReportRedeem ORDER BY createDate`

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rewards := []Reward{}
	for rows.Next() {
		var reward Reward

		err := rows.Scan(
			&reward.CreateDate,
			&reward.RewardName,
			&reward.Total,
			&reward.CustomerName,
			&reward.Email,
			&reward.PhoneNumber,
			&reward.AddressDetail,
			&reward.SubDistrict,
			&reward.District,
			&reward.Province,
			&reward.PostalCode,
			&reward.CN,
		)
		if err != nil {
			return nil, err
		}
		rewards = append(rewards, reward)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return rewards, nil
}
