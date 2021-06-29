package realtime

import (
	"encoding/json"

	"github.com/Jeffail/gabs"
	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
)

// GetInitialLivechatData returns the result of livechat:getInitialData
func (c Client) GetInitialLivechatData(visitorToken string) (*models.LivechatInitialData, error) {
	rawResponse, err := c.ddp.Call("livechat:getInitialData", visitorToken)
	if err != nil {
		return nil, err
	}

	document, err := gabs.Consume(rawResponse.(map[string]interface{}))
	if err != nil {
		return nil, err
	}

	return initialLivechatDataFromDocument(document), nil
}

// RegisterGuest registers a livechat guest
func (c Client) RegisterGuest(guestInfo *models.VisitorInfo) (*models.Visitor, error) {
	rawData, err := c.ddp.Call("livechat:registerGuest", guestInfo)
	if err != nil {
		return nil, err
	}

	document, err := gabs.Consume(rawData.(map[string]interface{}))
	if err != nil {
		return nil, err
	}

	return visitorFromDocument(document), nil
}

// SendLivechatMessage sends a live chat message
func (c Client) SendLivechatMessage(msg *models.LivechatMessage) error {
	_, err := c.ddp.Call("sendMessageLivechat", msg)
	return err
}

func initialLivechatDataFromDocument(doc *gabs.Container) *models.LivechatInitialData {
	allowSwitchDepartments, _ := doc.Path("allowSwitchDepartments").Data().(bool)
	nameFieldRegistrationForm, _ := doc.Path("nameFieldRegistratinForm").Data().(bool)
	emailFieldRegistrationForm, _ := doc.Path("emailFieldRegistratinForm").Data().(bool)

	return &models.LivechatInitialData{
		Enabled:                    doc.Path("enabled").Data().(bool),
		Title:                      doc.Path("title").String(),
		Color:                      doc.Path("color").String(),
		RegistrationForm:           doc.Path("registrationForm").Data().(bool),
		Room:                       doc.Path("room").Bytes(),
		Visitor:                    doc.Path("visitor").Bytes(),
		Triggers:                   doc.Path("triggers").Bytes(),
		Departments:                doc.Path("departments").Bytes(),
		AllowSwitchingDepartments:  allowSwitchDepartments,
		Online:                     doc.Path("online").Data().(bool),
		OfflineColor:               doc.Path("offlineColor").String(),
		OfflineMessage:             doc.Path("offlineMessage").String(),
		OfflineSuccessMessage:      doc.Path("offlineSuccessMessage").String(),
		OfflineUnavailableMessage:  doc.Path("offlineUnavailableMessage").String(),
		NameFieldRegistrationForm:  nameFieldRegistrationForm,
		EmailFieldRegistrationForm: emailFieldRegistrationForm,
		OfflineTitle:               doc.Path("offlineTitle").String(),
		Language:                   doc.Path("language").String(),
		Transcript:                 doc.Path("transcript").Data().(bool),
		TranscriptMessage:          doc.Path("transacriptMessage").String(),
		AgentData:                  doc.Path("agentData").Bytes(),
	}
}

func visitorFromDocument(doc *gabs.Container) *models.Visitor {
	visitorEmailsBytes := doc.Path("visitor.visitorEmails").Bytes()
	visitorEmails := []models.VisitorEmail{}
	json.Unmarshal(visitorEmailsBytes, &visitorEmails)
	return &models.Visitor{
		UserID:        doc.Path("userId").String(),
		Name:          doc.Path("visitor.name").String(),
		Token:         doc.Path("visitor.token").String(),
		Username:      doc.Path("visitor.username").String(),
		VisitorEmails: visitorEmails,
	}
}
