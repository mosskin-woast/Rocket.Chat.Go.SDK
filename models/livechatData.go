package models

// LivechatInitialData is the response to the livechat:getInitialData method
type LivechatInitialData struct {
	Enabled                    bool   `json:"enabled"`
	Title                      string `json:"title"`
	Color                      string `json:"color"`
	RegistrationForm           bool   `json:"registrationForm"`
	Room                       []byte `json:"room"`        // this is an object, @TODO implement struct
	Visitor                    []byte `json:"visitor"`     // this is an obect, @TODO implement struct
	Triggers                   []byte `json:"triggers"`    // this is an array, @TODO implement array/structs
	Departments                []byte `json:"departments"` // this is an array, @TODO implement array/structs
	AllowSwitchingDepartments  bool   `json:"allowSwitchingDepartments"`
	Online                     bool   `json:"online"`
	OfflineColor               string `json:"offlineColor"`
	OfflineMessage             string `json:"offlineMessage"`
	OfflineSuccessMessage      string `json:"offlineSuccessMessage"`
	OfflineUnavailableMessage  string `json:"offlineUnavailableMessage"`
	NameFieldRegistrationForm  bool   `json:"nameFieldRegistrationForm"`
	EmailFieldRegistrationForm bool   `json:"emailFieldRegistrationForm"`
	OfflineTitle               string `json:"offlineTitle"`
	Language                   string `json:"language"`
	Transcript                 bool   `json:"transcript"`
	TranscriptMessage          string `json:"transcriptMessage"`
	AgentData                  []byte `json:"agentData"` // this is an object, @TODO implement struct
}

// LivechatMessage is a message sent to a live chat
type LivechatMessage struct {
	ID      string `json:"id"`
	RoomID  string `json:"rid"`
	Message string `json:"msg"`
	Token   string `json:"token"`
}
