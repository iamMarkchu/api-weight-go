package biz

import (
	"api-weight-go/api/weight"
	"api-weight-go/internal/model/orm"
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type WeightRepo interface {
	Record(ctx context.Context, data orm.WeightModel) (id uint64, err error)
	GetList(ctx context.Context, uid uint64, limit int32) (list []orm.WeightModel, err error)
	GetListByMonth(ctx context.Context, uid uint64, month string) (list []orm.WeightModel, err error)
}

type WeightUserCase struct {
	repo WeightRepo
	log  *log.Helper
}

func (uc *WeightUserCase) RecordWeight(ctx context.Context, req *weight.RecordWeightRequest, uid uint64) (err error) {
	uc.log.WithContext(ctx).Infof("param:%+v", req)
	// 校验用户是否合法
	// 转换成yyyymmdd
	dateStr := time.Unix(int64(req.Date), 0).Format("20060102")
	_, err = uc.repo.Record(ctx, orm.WeightModel{
		Date:    dateStr,
		Weight:  req.Weight,
		UID:     uid,
		Created: time.Now().Unix(),
		Updated: time.Now().Unix(),
	})
	if err != nil {
		return
	}
	return nil
}

func (uc *WeightUserCase) GetRecentWeight(ctx context.Context, req *weight.GetRecentWeightRequest, uid uint64) (res []*weight.WeightEntity, err error) {
	uc.log.WithContext(ctx).Infof("param:%+v", req)
	limit := req.GetLatest()
	list, err := uc.repo.GetList(ctx, uid, limit)
	if err != nil {
		return
	}
	res = uc.transData(list)
	return
}

func (uc *WeightUserCase) transData(list []orm.WeightModel) (res []*weight.WeightEntity) {
	res = make([]*weight.WeightEntity, 0)
	if len(list) == 0 {
		return
	}
	for _, item := range list {
		res = append(res, &weight.WeightEntity{
			Id:     item.ID,
			Date:   item.Date,
			Weight: item.Weight,
			Uid:    item.UID,
		})
	}
	return
}

func (uc *WeightUserCase) GetWeightByMonth(ctx context.Context, req *weight.GetWeightByMonthRequest, uid uint64) (res []*weight.WeightEntity, err error) {
	uc.log.WithContext(ctx).Infof("param:%+v", req)
	month := req.GetMonth()
	list, err := uc.repo.GetListByMonth(ctx, uid, month)
	if err != nil {
		return
	}
	res = uc.transData(list)
	return
}

func NewWeightUserCase(repo WeightRepo, logger log.Logger) *WeightUserCase {
	return &WeightUserCase{repo: repo, log: log.NewHelper(logger)}
}
