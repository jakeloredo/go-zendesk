package zendesk

import (
	"fmt"
)

type ExecutedView struct {
	Rows []ViewRow `json:"rows"`
}

type ViewRow struct {
	TicketID int    `json:"ticket_id"`
	Subject  string `json:"subject"`
	Created  string `json:"created"`
}

func (zd *Zendesk) ExecuteView(viewID int) (ResponseData, error) {
	// TODO: Change this to get each page until next_page is null
	endpoint := fmt.Sprintf("views/%d/execute.json?per_page=500", viewID)
	return zd.get(endpoint)
}

func (zd *Zendesk) ExecuteView2(viewID int) (PaginatedResponseData, error) {
	result := PaginatedResponseData{}
	// TODO: Change this to get each page until next_page is null
	perPage := 100
	pageNo := 1
	endpoint := fmt.Sprintf("views/%d/execute.json?page=%d&per_page=%d", viewID, pageNo, perPage)

	page, err := zd.get(endpoint)
	if err != nil {
		return result, err
	}
	
	nPages := page.Count / perPage
	result.PageCount = nPages
	result.Pages = make([]ResponseData, nPages)
	result.Pages[0] = page

	for i := 2; i <= nPages; i++ {
		endpoint = fmt.Sprintf("views/%d/execute.json?page=%d&per_page=%d", viewID, i, perPage)
		page, err = zd.get(endpoint)
		if err != nil {
			return result, err
		}
		result.Pages[i-1] = page
	}

	return result, nil
}

