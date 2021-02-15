package task

type PlayInfo struct {
	Code int64 `json:"code"`
	Data []struct {
		Cid       int64 `json:"cid"`
		Dimension struct {
			Height int64 `json:"height"`
			Rotate int64 `json:"rotate"`
			Width  int64 `json:"width"`
		} `json:"dimension"`
		Duration int64  `json:"duration"`
		From     string `json:"from"`
		Page     int64  `json:"page"`
		Part     string `json:"part"`
		Vid      string `json:"vid"`
		Weblink  string `json:"weblink"`
	} `json:"data"`
	Message string `json:"message"`
	TTL     int64  `json:"ttl"`
}

type DanmuContent struct {
	AppearTime    string
	Mode          string
	FontSize      string
	Color         string
	SendTimeStamp string
	DanmuPool     string
	UserHash      string
	DataBaseID    string
	Content       string
}

type I struct {
	Chatserver string `xml:"chatserver"`
	Chatid     string `xml:"chatid"`
	D          []D    `xml:"d"`
}
type D struct {
	P    string `xml:"p,attr"`
	Text string `xml:",chardata"`
}
