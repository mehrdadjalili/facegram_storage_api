package utilitymodel

type UserBlock struct {
	TimeType    string `json:"time_type"`
	TimeNumber  int64  `json:"time_number"`
	BlockReason string `json:"block_reason"`
}
