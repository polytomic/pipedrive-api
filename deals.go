package pipedrive

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

// DealService handles deals related
// methods of the Pipedrive API.
//
// Pipedrive API dcos: https://developers.pipedrive.com/docs/api/v1/#!/Deals
type DealService service

// Deal represents a Pipedrive deal.
type Deal struct {
	ID            int `json:"id"`
	CreatorUserID *struct {
		ID         int         `json:"id"`
		Name       string      `json:"name"`
		Email      string      `json:"email"`
		HasPic     interface{} `json:"has_pic"`
		PicHash    string      `json:"pic_hash"`
		ActiveFlag bool        `json:"active_flag"`
		Value      int         `json:"value"`
	} `json:"creator_user_id"`
	UserID *struct {
		ID         int         `json:"id"`
		Name       string      `json:"name"`
		Email      string      `json:"email"`
		HasPic     interface{} `json:"has_pic"`
		PicHash    string      `json:"pic_hash"`
		ActiveFlag bool        `json:"active_flag"`
		Value      int         `json:"value"`
	} `json:"user_id"`
	PersonID *struct {
		Name  string `json:"name"`
		Email []struct {
			Value   string `json:"value"`
			Primary bool   `json:"primary"`
		} `json:"email"`
		Phone []struct {
			Value   string `json:"value"`
			Primary bool   `json:"primary"`
		} `json:"phone"`
		Value int `json:"value"`
	} `json:"person_id"`
	OrgID *struct {
		Name        string      `json:"name"`
		PeopleCount int         `json:"people_count"`
		OwnerID     int         `json:"owner_id"`
		Address     interface{} `json:"address"`
		CcEmail     string      `json:"cc_email"`
		Value       int         `json:"value"`
	} `json:"org_id"`
	StageID                  int                    `json:"stage_id"`
	Title                    string                 `json:"title"`
	Value                    float64                `json:"value"`
	Currency                 string                 `json:"currency"`
	AddTime                  string                 `json:"add_time"`
	UpdateTime               string                 `json:"update_time"`
	StageChangeTime          string                 `json:"stage_change_time"`
	Active                   bool                   `json:"active"`
	Deleted                  bool                   `json:"deleted"`
	Status                   string                 `json:"status"`
	Probability              interface{}            `json:"probability"`
	NextActivityDate         interface{}            `json:"next_activity_date"`
	NextActivityTime         interface{}            `json:"next_activity_time"`
	NextActivityID           interface{}            `json:"next_activity_id"`
	LastActivityID           int                    `json:"last_activity_id"`
	LastActivityDate         string                 `json:"last_activity_date"`
	LostReason               string                 `json:"lost_reason"`
	VisibleTo                interface{}            `json:"visible_to"`
	CloseTime                string                 `json:"close_time"`
	PipelineID               int                    `json:"pipeline_id"`
	WonTime                  interface{}            `json:"won_time"`
	FirstWonTime             interface{}            `json:"first_won_time"`
	LostTime                 string                 `json:"lost_time"`
	ProductsCount            int                    `json:"products_count"`
	FilesCount               int                    `json:"files_count"`
	NotesCount               int                    `json:"notes_count"`
	FollowersCount           int                    `json:"followers_count"`
	EmailMessagesCount       int                    `json:"email_messages_count"`
	ActivitiesCount          int                    `json:"activities_count"`
	DoneActivitiesCount      int                    `json:"done_activities_count"`
	UndoneActivitiesCount    int                    `json:"undone_activities_count"`
	ReferenceActivitiesCount int                    `json:"reference_activities_count"`
	ParticipantsCount        int                    `json:"participants_count"`
	ExpectedCloseDate        interface{}            `json:"expected_close_date"`
	LastIncomingMailTime     interface{}            `json:"last_incoming_mail_time"`
	LastOutgoingMailTime     interface{}            `json:"last_outgoing_mail_time"`
	StageOrderNr             int                    `json:"stage_order_nr"`
	PersonName               interface{}            `json:"person_name"`
	OrgName                  interface{}            `json:"org_name"`
	NextActivitySubject      interface{}            `json:"next_activity_subject"`
	NextActivityType         interface{}            `json:"next_activity_type"`
	NextActivityDuration     interface{}            `json:"next_activity_duration"`
	NextActivityNote         interface{}            `json:"next_activity_note"`
	FormattedValue           string                 `json:"formatted_value"`
	RottenTime               interface{}            `json:"rotten_time"`
	WeightedValue            float64                `json:"weighted_value"`
	FormattedWeightedValue   string                 `json:"formatted_weighted_value"`
	OwnerName                string                 `json:"owner_name"`
	CcEmail                  string                 `json:"cc_email"`
	OrgHidden                bool                   `json:"org_hidden"`
	PersonHidden             bool                   `json:"person_hidden"`
	CustomFields             map[string]interface{} `json:"-"`
}

type _Deal Deal

func (d Deal) String() string {
	return Stringify(d)
}

func (d *Deal) UnmarshalJSON(b []byte) error {
	obj := _Deal{}
	err := json.Unmarshal(b, &obj)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, &(obj.CustomFields))
	if err != nil {
		return err
	}

	typ := reflect.TypeOf(obj)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		jsonTag := strings.Split(field.Tag.Get("json"), ",")[0]
		if jsonTag != "" && jsonTag != "-" {
			delete(obj.CustomFields, jsonTag)
		}
	}

	*d = Deal(obj)

	return nil
}

// DealsResponse represents multiple deals response.
type DealsResponse struct {
	Success        bool           `json:"success,omitempty"`
	Data           []Deal         `json:"data,omitempty"`
	AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// DealResponse represents single deal response.
type DealResponse struct {
	Success        bool           `json:"success,omitempty"`
	Data           Deal           `json:"data,omitempty"`
	AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// ListUpdates about a deal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/get_deals_id_flow
func (s *DealService) ListUpdates(ctx context.Context, id int) (*DealsResponse, *Response, error) {
	uri := fmt.Sprintf("/deals/%v/flow", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *DealsResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Find deals by name.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/get_deals_find
func (s *DealService) Find(ctx context.Context, term string) (*DealsResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/deals/find", &SearchOptions{
		Term: term,
	}, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *DealsResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

type DealsListOptions struct {
	UserID     int    `url:"user_id,omitempty"`
	FilterID   int    `url:"filter_id,omitempty"`
	StageID    int    `url:"stage_id,omitempty"`
	Status     string `url:"status,omitempty"`
	Start      int    `url:"start,omitempty"`
	Limit      int    `url:"limit,omitempty"`
	Sort       string `url:"sort,omitempty"`
	OwnedByYou int    `url:"owned_by_you,omitempty"`
}

// List deals.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/get_deals
func (s *DealService) List(ctx context.Context, opts *DealsListOptions) (*DealsResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/deals", opts, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *DealsResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

type DealSearchParams struct {
	Term           string   `url:"term,omitempty"`
	Fields         []string `url:"fields,omitempty,comma"`
	ExactMatch     bool     `url:"exact_match,omitempty"`
	OrganizationID int      `url:"organization_id,omitempty"`
	PersonID       int      `url:"person_id,omitempty"`
	Status         string   `url:"status,omitempty"`
	IncludeFields  string   `url:"include_fields,omitempty"`
	Start          int      `url:"start,omitempty"`
	Limit          int      `url:"limit,omitempty"`
}

type DealSearchResult struct {
	ResultScore float64 `json:"result_score"`
	Item        Deal    `json:"item"`
}

type DealsSearchResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Items []DealSearchResult `json:"items"`
	} `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

// Search for deal(s)
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/Deals#searchDeals
func (s *DealService) Search(ctx context.Context, searchParams DealSearchParams) (*DealsSearchResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/deals/search", searchParams, nil)

	if err != nil {
		return nil, nil, err
	}

	var response *DealsSearchResponse
	resp, err := s.client.Do(ctx, req, &response)

	if err != nil {
		return nil, resp, err
	}

	return response, resp, nil
}

// Duplicate a deal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/post_deals_id_duplicate
func (s *DealService) Duplicate(ctx context.Context, id int) (*DealResponse, *Response, error) {
	uri := fmt.Sprintf("/deals/%v/duplicate", id)
	req, err := s.client.NewRequest(http.MethodPost, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *DealResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// DealsMergeOptions specifices the optional parameters to the
// DealService.Merge method.
type DealsMergeOptions struct {
	MergeWithID uint `url:"merge_with_id,omitempty"`
}

// Merge two deals.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/put_deals_id_merge
func (s *DealService) Merge(ctx context.Context, id int, opt *DealsMergeOptions) (*Response, error) {
	uri := fmt.Sprintf("/deals/%v/merge", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, nil, opt)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// DealsUpdateOptions specifices the optional parameters to the
// DealService.Update method.
type DealsUpdateOptions struct {
	Title          string `json:"title,omitempty"`
	Value          string `json:"value,omitempty"`
	Currency       string `json:"currency,omitempty"`
	UserID         uint   `json:"user_id,omitempty"`
	PersonID       uint   `json:"person_id,omitempty"`
	OrganizationID uint   `json:"org_id,omitempty"`
	StageID        uint   `json:"stage_id,omitempty"`
	Status         string `json:"status,omitempty"`
	LostReason     string `json:"lost_reason,omitempty"`
	VisibleTo      uint   `json:"visible_to,omitempty"`

	CustomFields map[string]interface{} `json:"-"`
}

func (d DealsUpdateOptions) MarshalJSON() ([]byte, error) {
	fields := map[string]interface{}{}
	for k, v := range d.CustomFields {
		fields[k] = v
	}

	if d.Title != "" {
		fields["title"] = d.Title
	}
	if d.Value != "" {
		fields["value"] = d.Value
	}
	if d.Currency != "" {
		fields["currency"] = d.Currency
	}
	if d.UserID != 0 {
		fields["user_id"] = d.UserID
	}
	if d.PersonID != 0 {
		fields["person_id"] = d.PersonID
	}
	if d.OrganizationID != 0 {
		fields["org_id"] = d.OrganizationID
	}
	if d.StageID != 0 {
		fields["stage_id"] = d.StageID
	}
	if d.Status != "" {
		fields["status"] = d.Status
	}
	if d.LostReason != "" {
		fields["lost_reason"] = d.LostReason
	}
	if d.VisibleTo != 0 {
		fields["visible_to"] = d.VisibleTo
	}

	return json.Marshal(fields)
}

// Update a deal.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/put_deals_id
func (s *DealService) Update(ctx context.Context, id int, opt *DealsUpdateOptions) (*DealResponse, *Response, error) {
	uri := fmt.Sprintf("/deals/%v", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var updated *DealResponse
	resp, err := s.client.Do(ctx, req, &updated)
	if err != nil {
		return nil, resp, err
	}
	return updated, resp, nil
}

// DeleteFollower of a deal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/delete_deals_id_followers_follower_id
func (s *DealService) DeleteFollower(ctx context.Context, id int, followerID int) (*Response, error) {
	uri := fmt.Sprintf("/deals/%v/followers/%v", id, followerID)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// DeleteMultiple deletes deals in bulk.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/delete_deals
func (s *DealService) DeleteMultiple(ctx context.Context, ids []int) (*Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, "/deals", &DeleteMultipleOptions{
		Ids: arrayToString(ids, ","),
	}, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// DeleteParticipant deletes participant in a deal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/delete_deals_id_participants_deal_participant_id
func (s *DealService) DeleteParticipant(ctx context.Context, dealID int, participantID int) (*Response, error) {
	uri := fmt.Sprintf("/deals/%v/participants/%v", dealID, participantID)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Delete a deal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/delete_deals_id
func (s *DealService) Delete(ctx context.Context, id int) (*Response, error) {
	uri := fmt.Sprintf("/deals/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// DeleteAttachedProduct deletes attached product.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/delete_deals_id_products_product_attachment_id
func (s *DealService) DeleteAttachedProduct(ctx context.Context, dealID int, productAttachmentID int) (*Response, error) {
	uri := fmt.Sprintf("/deals/%v/products/%v", dealID, productAttachmentID)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// DealCreateOptions specifices the optional parameters to the
// DealsService.Create method.
type DealCreateOptions struct {
	Title       string    `json:"title,omitempty"`
	Value       string    `json:"value,omitempty"`
	Currency    string    `json:"currency,omitempty"`
	UserID      uint      `json:"user_id,omitempty"`
	PersonID    uint      `json:"person_id,omitempty"`
	OrgID       uint      `json:"org_id,omitempty"`
	StageID     uint      `json:"stage_id,omitempty"`
	Status      string    `json:"status,omitempty"`
	Probability uint      `json:"probability,omitempty"`
	LostReason  string    `json:"lost_reason,omitempty"`
	AddTime     Timestamp `json:"add_time,omitempty"`
	VisibleTo   VisibleTo `json:"visible_to,omitempty"`

	CustomFields map[string]interface{} `json:"-"`
}

func (d DealCreateOptions) MarshalJSON() ([]byte, error) {
	fields := map[string]interface{}{}
	for k, v := range d.CustomFields {
		fields[k] = v
	}

	if d.Title != "" {
		fields["title"] = d.Title
	}
	if d.Value != "" {
		fields["value"] = d.Value
	}
	if d.Currency != "" {
		fields["currency"] = d.Currency
	}
	if d.UserID != 0 {
		fields["user_id"] = d.UserID
	}
	if d.PersonID != 0 {
		fields["person_id"] = d.PersonID
	}
	if d.OrgID != 0 {
		fields["org_id"] = d.OrgID
	}
	if d.StageID != 0 {
		fields["stage_id"] = d.StageID
	}
	if d.Status != "" {
		fields["status"] = d.Status
	}
	if d.Probability != 0 {
		fields["probability"] = d.Probability
	}
	if d.LostReason != "" {
		fields["lost_reason"] = d.LostReason
	}
	if !d.AddTime.IsZero() {
		fields["add_time"] = d.AddTime.FormatFull()
	}
	if d.VisibleTo != 0 {
		fields["visible_to"] = d.VisibleTo
	}

	return json.Marshal(fields)
}

// Create a new deal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/post_deals
func (s *DealService) Create(ctx context.Context, opt *DealCreateOptions) (*DealResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/deals", nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *DealResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}
