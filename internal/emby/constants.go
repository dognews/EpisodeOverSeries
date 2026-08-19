package emby

// Constants
const serverIp = "http://alpinehome-host:8096"
const apiKey = "REMOVED"
const pageLength = 50

// Emby Types
type itemEntry struct {
	Id        string            `json:"id"`
	ImageTags map[string]string `json:"imagetags"`
}

type itemsRequest struct {
	Items []itemEntry `json:"items"`
}

// Custom Types
type ImageData struct {
	Id      string
	Primary string
	Thumb   string
}
