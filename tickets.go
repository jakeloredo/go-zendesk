package zendesk

import (
	"fmt"
)

// Field is a struct for custom_fields array
type Field struct {
	ID    int         `json:"id,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// NewTicket contains most of the writeable fields of a ticket
type NewTicket struct {
	Comment             TicketComment `json:"comment,omitempty"`
	ExternalID          string        `json:"external_id,omitempty"`
	Type                string        `json:"type,omitempty"`
	Subject             string        `json:"subject,omitempty"`
	RawSubject          string        `json:"raw_subject,omitempty"`
	Priority            string        `json:"priority,omitempty"`
	Status              string        `json:"status,omitempty"`
	Recipient           string        `json:"recipient,omitempty"`
	RequesterID         int           `json:"requester_id,omitempty"`
	SubmitterID         int           `json:"submitter_id,omitempty"`
	AssigneeID          int           `json:"assignee_id,omitempty"`
	OrganizationID      int           `json:"organization_id,omitempty"`
	GroupID             int           `json:"group_id,omitempty"`
	ForumTopicID        int           `json:"forum_topic_id,omitempty"`
	ProblemID           int           `json:"problem_id,omitempty"`
	DueAt               string        `json:"due_at,omitempty"`
	Tags                []string      `json:"tags,omitempty"`
	CustomFields        []Field       `json:"custom_fields,omitempty"`
	ViaFollowupSourceID int           `json:"via_followup_source_id,omitempty"`
	TicketFormID        int           `json:"ticket_form_id,omitempty"`
	BrandID             int           `json:"brand_id,omitempty"`

	//CollaboratorIds     array  `json:"collaborator_ids"`
	//Collaborators       array  `json:"collaborators"`
	//FollowerIds         array  `json:"follower_ids"`
	//MacroIds            array  `json:"macro_ids"`
}

// Ticket contains most of the fields returned by /tickets/{id}.json endpoint
type Ticket struct {
	ID                  int      `json:"id,omitempty"`
	URL                 string   `json:"url,omitempty"`
	ExternalID          string   `json:"external_id,omitempty"`
	Type                string   `json:"type,omitempty"`
	Subject             string   `json:"subject,omitempty"`
	RawSubject          string   `json:"raw_subject,omitempty"`
	Description         string   `json:"description,omitempty"`
	Priority            string   `json:"priority,omitempty"`
	Status              string   `json:"status,omitempty"`
	Recipient           string   `json:"recipient,omitempty"`
	RequesterID         int      `json:"requester_id,omitempty"`
	SubmitterID         int      `json:"submitter_id,omitempty"`
	AssigneeID          int      `json:"assignee_id,omitempty"`
	OrganizationID      int      `json:"organization_id,omitempty"`
	GroupID             int      `json:"group_id,omitempty"`
	ForumTopicID        int      `json:"forum_topic_id,omitempty"`
	ProblemID           int      `json:"problem_id,omitempty"`
	HasIncidents        bool     `json:"has_incidents,omitempty"`
	DueAt               string   `json:"due_at,omitempty"`
	Tags                []string `json:"tags,omitempty"`
	CustomFields        []Field  `json:"custom_fields,omitempty"`
	ViaFollowupSourceID int      `json:"via_followup_source_id,omitempty"`
	TicketFormID        int      `json:"ticket_form_id,omitempty"`
	BrandID             int      `json:"brand_id,omitempty"`
	AllowChannelback    bool     `json:"allow_channelback,omitempty"`
	IsPublic            bool     `json:"is_public,omitempty"`
	CreatedAt           string   `json:"created_at,omitempty"`
	UpdatedAt           string   `json:"updated_at,omitempty"`

	//POST only: MacroIds            array  `json:"macro_ids"`
	//Via                 Via    `json:"via"`
	//SatisfactionRating  object `json:"satisfaction_rating"`
	//SharingAgreementIds array  `json:"sharing_agreement_ids"`
	//FollowupIds         array  `json:"followup_ids"`
	//CollaboratorIds     array  `json:"collaborator_ids"`
	//Collaborators       array  `json:"collaborators"`
	//FollowerIds         array  `json:"follower_ids"`
}

func (zd *Zendesk) GetOneTicket(ticketID int) (ResponseData, error) {
	// Concatenate ticket ID to "tickets/" make the endpoint string
	endpoint := fmt.Sprintf("/tickets/%d", ticketID)
	return zd.get(endpoint)
}

func (zd *Zendesk) CreateOneTicket(ticketData NewTicket) (ResponseData, error) {
	endpoint := "tickets"
	payload := map[string]interface{}{"ticket": ticketData}
	return zd.post(endpoint, payload)
}

// SolveTicket attempts to solve a Zendesk ticket
// It assigns the ticket to the agent whose username is in the Zendesk struct
func (zd *Zendesk) SolveTicket(ticketID int) (ResponseData, error) {

	payload := make(map[string]interface{})
	payload["ticket"] = map[string]string{"assignee_email": zd.Username, "status": "solved"}

	endpoint := fmt.Sprintf("/tickets/%d", ticketID)

	return zd.put(endpoint, payload)
}

func (zd *Zendesk) UpdateOneTicket(ticketData Ticket) (ResponseData, error) {
	endpoint := fmt.Sprintf("/tickets/%d", ticketData.ID)
	payload := map[string]interface{}{"ticket": ticketData}
	payload["ticket"] = ticketData
	return zd.put(endpoint, payload)
}
