package request

type GrowBoxCreateRequest struct {
	Name        string `validate:"required min=1,max=100" json:"namebox"`
	Description string `validate:"required min=1,max=100" json:"description"`
	Count_fans  int64  `validate:"required min=1,max=100" json:"count_fans"`
	Filtration  bool   `validate:"required min=1,max=100" json:"filtration"`
	Dimensions  int64  `validate:"required min=1,max=100" json:"dimensions"`
	Automation  bool   `validate:"required min=1,max=100" json:"Automation"`
}
