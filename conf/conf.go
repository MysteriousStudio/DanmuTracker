package conf

type TrackerConf struct {
	ServerPort string `json:"server_port,omitempty"`
	AdminToken string `json:"admin_token,omitempty"`
}
