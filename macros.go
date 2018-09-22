package zendesk

import (
	"fmt"
)

type Macro struct {
	ID          int    `json:"id"`
	Active      bool   `json:"active"`
	Description string `json:"description"`
	Position    int    `json:"position"`
	Title       string `json:"title"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`

	//Actions     Actions `json:"actions"`
	//Restriction object `json:"restriction"`
}

func (m *Macro) apply(ticketID int) {

}

func (zd *Zendesk) GetMacroApply(macroID int) (ResponseData, error) {
	endpoint := fmt.Sprintf("macros/%d/apply", macroID)
	return zd.get(endpoint)
}

func (zd *Zendesk) ApplyMacro(macroID int, ticketID int) (ResponseData, error) {

	r, getMacroErr := zd.GetMacroApply(macroID)

	if getMacroErr != nil {
		return r, getMacroErr
	}

	var ticketUpdates Ticket
	ticketUpdates = r.Result.Ticket
	ticketUpdates.ID = ticketID

	return zd.UpdateOneTicket(ticketUpdates)
}
