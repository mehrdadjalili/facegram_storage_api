package responsemodel

//RequestLimiterInfo struct
type RequestLimiterInfo struct {
	TimeType  string `json:"time_type"`
	LimitTime int    `json:"limit_time"`
}
