package usermodel

type UserPageInfo struct {
	ID       string `json:"id"`
	UniqueID string `json:"unique_id"`
	Kind     string `json:"kind"`
}

type UserModel struct {
	UserID                 string                 `json:"user_id"`
	PageID                 string                 `json:"page_id"`
	Role                   string                 `json:"role"`
	Token                  string                 `json:"token"`
	SessionID              string                 `json:"session_id"`
	PublicKey              string                 `json:"public_key"`
	RequestApplicationInfo RequestApplicationInfo `json:"request_application_info"`
	Location               RequestLocationInfo    `json:"location"`
}

type UserTokenInfo struct {
	UserID string `json:"user_id"`
	PageID string `json:"page_id"`
	Sid    string `json:"sid"`
	Kind   string `json:"kind"`
	Role   string `json:"role"`
}

type RequestApplicationInfo struct {
	AppVersion int
	AppType    string
	Language   string
}

type RequestLocationInfo struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Country   string `json:"country"`
	State     string `json:"state"`
	City      string `json:"city"`
}

type UserInfo struct {
	UserID    string              `json:"user_id"`
	PageID    string              `json:"page_id"`
	Role      string              `json:"role"`
	Token     string              `json:"token"`
	SessionID string              `json:"session_id"`
	Location  RequestLocationInfo `json:"location"`
}
