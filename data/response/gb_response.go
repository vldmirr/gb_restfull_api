package response

type GrowBoxResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Count_fans  int64  `json:"count_fans"`
	Filtration  bool   `json:"filtration"`
	Dimensions  int64  `json:"dimensions"`
	Automation  bool   `json:"Automation"`
}
