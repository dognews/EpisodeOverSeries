package emby

// Constants
var serverIp = "NONE"
var apiKey = "NONE"

const pageLength = 50
const BufferDefaultSize = 1048576 // (1MB) ^2 aligned

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
