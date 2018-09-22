package zendesk

import (
	"fmt"
)

type NewTicketComment struct {
	Body     string `json:"body"`
	HTMLBody string `json:"html_body"`
	Public   bool   `json:"public"`
	AuthorID int    `json:"author_id"`
}

type Attachment struct {
	URL        string `json:"url"`
	ID         int    `json:"id"`
	FileName   string `json:"file_name"`
	ContentURL string `json:"content_url"`
}

type TicketComment struct {
	ID          int          `json:"id"`
	Type        string       `json:"type"`
	Body        string       `json:"body"`
	HTMLBody    string       `json:"html_body"`
	PlainBody   string       `json:"plain_body"`
	Public      bool         `json:"public"`
	AuthorID    int          `json:"author_id"`
	CreatedAt   string       `json:"created_at"`
	Attachments []Attachment `json:"attachments"`

	// TODO: Support these
	// Via         object `json:"via"`
	// Metadata    object `json:"metadata"`
}

func (zd *Zendesk) GetTicketComments(ticketID int) (ResponseData, error) {
	endpoint := fmt.Sprintf("tickets/%d/comments", ticketID)
	return zd.get(endpoint)
}
