package model

type Reply struct {
	GroupId    int64  `json:"group_id"`
	Message    string `json:"message"`
	AutoEscape string `json:"auto_escape"`
}
