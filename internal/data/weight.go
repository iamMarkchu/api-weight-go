package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/iamMarkchu/api-weight-go/internal/biz"
	"github.com/iamMarkchu/api-weight-go/internal/model/orm"
	"gorm.io/gorm"
)

type weightRepo struct {
	data *Data
	log  *log.Helper
}

func (r *weightRepo) GetListByMonth(ctx context.Context, uid uint64, month string) (list []orm.WeightModel, err error) {
	list = make([]orm.WeightModel, 0)
	err = r.data.db.Model(orm.WeightModel{}).Where("uid = ? AND status = ? AND date like ?", uid, 0, month+"%").Order("id desc").Find(&list).Error
	return
}

func (r *weightRepo) GetList(ctx context.Context, uid uint64, limit int32) (list []orm.WeightModel, err error) {
	list = make([]orm.WeightModel, 0)
	err = r.data.db.Model(orm.WeightModel{}).Where("uid = ? AND status = ?", uid, 0).Limit(int(limit)).Order("id desc").Find(&list).Error
	return
}

func (r *weightRepo) Record(ctx context.Context, data orm.WeightModel) (id uint64, err error) {
	err = r.data.db.Transaction(func(tx *gorm.DB) error {
		eData := orm.WeightModel{}
		err = tx.Model(orm.WeightModel{}).Where("`date` = ? AND uid = ?", data.Date, data.UID).First(&eData).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("db error")
		}
		if eData.ID != 0 {
			eData.Weight = data.Weight
			eData.Updated = data.Updated
			err = tx.Updates(&eData).Error
			id = eData.ID
		} else {
			err = tx.Model(orm.WeightModel{}).Create(&data).Error
			if err != nil {
				return errors.New("db error")
			}
			id = data.ID
		}
		return nil
	})
	return
}

func NewWeightRepo(data *Data, logger log.Logger) biz.WeightRepo {
	return &weightRepo{data: data, log: log.NewHelper(logger)}
}
