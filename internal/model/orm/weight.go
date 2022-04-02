package orm

type WeightModel struct {
	ID      uint64
	Date    string
	Weight  float32
	UID     uint64
	Img     string
	Status  uint8
	Created int64
	Updated int64
}

func (WeightModel) TableName() string {
	return "user_weight"
}
