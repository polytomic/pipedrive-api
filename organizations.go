package pipedrive

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

// OrganizationsService handles organization related
// methods of the Pipedrive API.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Organizations
type OrganizationsService service

// Organization represents a Pipedrive organization.
type Organization struct {
	ID        int `json:"id"`
	CompanyID int `json:"company_id"`
	OwnerID   struct {
		ID         int         `json:"id"`
		Name       string      `json:"name"`
		Email      string      `json:"email"`
		HasPic     interface{} `json:"has_pic"`
		PicHash    string      `json:"pic_hash"`
		ActiveFlag bool        `json:"active_flag"`
		Value      int         `json:"value"`
	} `json:"owner_id"`
	Name                            string                 `json:"name"`
	OpenDealsCount                  int                    `json:"open_deals_count"`
	RelatedOpenDealsCount           int                    `json:"related_open_deals_count"`
	ClosedDealsCount                int                    `json:"closed_deals_count"`
	RelatedClosedDealsCount         int                    `json:"related_closed_deals_count"`
	EmailMessagesCount              int                    `json:"email_messages_count"`
	PeopleCount                     int                    `json:"people_count"`
	ActivitiesCount                 int                    `json:"activities_count"`
	DoneActivitiesCount             int                    `json:"done_activities_count"`
	UndoneActivitiesCount           int                    `json:"undone_activities_count"`
	ReferenceActivitiesCount        int                    `json:"reference_activities_count"`
	FilesCount                      int                    `json:"files_count"`
	NotesCount                      int                    `json:"notes_count"`
	FollowersCount                  int                    `json:"followers_count"`
	WonDealsCount                   int                    `json:"won_deals_count"`
	RelatedWonDealsCount            int                    `json:"related_won_deals_count"`
	LostDealsCount                  int                    `json:"lost_deals_count"`
	RelatedLostDealsCount           int                    `json:"related_lost_deals_count"`
	ActiveFlag                      bool                   `json:"active_flag"`
	CategoryID                      interface{}            `json:"category_id"`
	PictureID                       interface{}            `json:"picture_id"`
	CountryCode                     interface{}            `json:"country_code"`
	FirstChar                       string                 `json:"first_char"`
	UpdateTime                      string                 `json:"update_time"`
	AddTime                         string                 `json:"add_time"`
	VisibleTo                       interface{}            `json:"visible_to"`
	NextActivityDate                string                 `json:"next_activity_date"`
	NextActivityTime                interface{}            `json:"next_activity_time"`
	NextActivityID                  int                    `json:"next_activity_id"`
	LastActivityID                  int                    `json:"last_activity_id"`
	LastActivityDate                string                 `json:"last_activity_date"`
	TimelineLastActivityTime        interface{}            `json:"timeline_last_activity_time"`
	TimelineLastActivityTimeByOwner interface{}            `json:"timeline_last_activity_time_by_owner"`
	Address                         string                 `json:"address"`
	AddressSubpremise               string                 `json:"address_subpremise"`
	AddressStreetNumber             string                 `json:"address_street_number"`
	AddressRoute                    string                 `json:"address_route"`
	AddressSublocality              string                 `json:"address_sublocality"`
	AddressLocality                 string                 `json:"address_locality"`
	AddressAdminAreaLevel1          string                 `json:"address_admin_area_level_1"`
	AddressAdminAreaLevel2          string                 `json:"address_admin_area_level_2"`
	AddressCountry                  string                 `json:"address_country"`
	AddressPostalCode               string                 `json:"address_postal_code"`
	AddressFormattedAddress         string                 `json:"address_formatted_address"`
	OwnerName                       string                 `json:"owner_name"`
	CcEmail                         string                 `json:"cc_email"`
	CustomFields                    map[string]interface{} `json:"-"`
}

type _Org Organization

func (o *Organization) UnmarshalJSON(b []byte) error {
	obj := _Org{}
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

	*o = Organization(obj)

	return nil
}

func (o Organization) String() string {
	return Stringify(o)
}

// OrganizationsResponse represents multiple organizations response.
type OrganizationsResponse struct {
	Success        bool           `json:"success"`
	Data           []Organization `json:"data"`
	AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// OrganizationResponse represents single organization response.
type OrganizationResponse struct {
	Success        bool           `json:"success"`
	Data           Organization   `json:"data"`
	AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

type OrganizationsListOptions struct {
	UserID    int    `url:"user_id,omitempty"`
	FilterID  int    `url:"filter_id,omitempty"`
	FirstChar string `url:"first_char,omitempty"`
	Start     int    `url:"start,omitempty"`
	Limit     int    `url:"limit,omitempty"`
	Sort      string `url:"sort,omitempty"`
}

// List all organizations.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Organizations/get_organizations
func (s *OrganizationsService) List(ctx context.Context, opts *OrganizationsListOptions) (*OrganizationsResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/organizations", opts, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *OrganizationsResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

type OrganizationSearchParams struct {
	Term       string   `url:"term,omitempty"`
	Fields     []string `url:"fields,omitempty,comma"`
	ExactMatch bool     `url:"exact_match,omitempty"`
	Status     string   `url:"status,omitempty"`
	Start      int      `url:"start,omitempty"`
	Limit      int      `url:"limit,omitempty"`
}

type OrganizationSearchResult struct {
	ResultScore float64
	Item        Organization
}

type OrganizationsSearchResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Items []OrganizationSearchResult `json:"items"`
	} `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

// Search for Organization(s)
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/Organizations#searchOrganizations
func (s *OrganizationsService) Search(ctx context.Context, searchParams OrganizationSearchParams) (*OrganizationsSearchResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/organizations/search", searchParams, nil)

	if err != nil {
		return nil, nil, err
	}

	var response *OrganizationsSearchResponse
	resp, err := s.client.Do(ctx, req, &response)

	if err != nil {
		return nil, resp, err
	}

	return response, resp, nil
}

// OrganizationUpdateOptions specifices the optional parameters to the
// OrganizationUpdateOptions.Update method.
type OrganizationUpdateOptions struct {
	Name      string    `json:"name,omitempty"`
	OwnerID   uint      `json:"owner_id,omitempty"`
	VisibleTo VisibleTo `json:"visible_to,omitempty"`
	Address   string    `json:"address,omitempty"`

	CustomFields map[string]interface{} `json:"-"`
}

func (d OrganizationUpdateOptions) MarshalJSON() ([]byte, error) {
	fields := map[string]interface{}{}
	for k, v := range d.CustomFields {
		fields[k] = v
	}

	if d.Name != "" {
		fields["name"] = d.Name
	}
	if d.OwnerID != 0 {
		fields["owner_id"] = d.OwnerID
	}
	if d.Address != "" {
		fields["address"] = d.Address
	}
	if d.VisibleTo != 0 {
		fields["visible_to"] = d.VisibleTo
	}

	return json.Marshal(fields)
}

// Update a specific person.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Organizations/put_persons_id
func (s *OrganizationsService) Update(ctx context.Context, id int, opt *OrganizationUpdateOptions) (*OrganizationResponse, *Response, error) {
	uri := fmt.Sprintf("/organizations/%v", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *OrganizationResponse
	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Merge organizations.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/organizations/_id_/merge
func (s *OrganizationsService) Merge(ctx context.Context, id int, mergeWithID int) (*OrganizationResponse, *Response, error) {
	uri := fmt.Sprintf("/organizations/%v/merge", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, nil, struct {
		MergeWithID int `url:"merge_with_id"`
	}{
		mergeWithID,
	})

	if err != nil {
		return nil, nil, err
	}

	var record *OrganizationResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// DeleteFollower deletes a follower from an organization.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Organizations/delete_organizations_id_followers_follower_id
func (s *OrganizationsService) DeleteFollower(ctx context.Context, id int, followerID int) (*Response, error) {
	uri := fmt.Sprintf("/organizations/%v/followers/%v", id, followerID)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Delete marks an organization as deleted.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Organizations/delete_organizations_id
func (s *OrganizationsService) Delete(ctx context.Context, id int) (*Response, error) {
	uri := fmt.Sprintf("/organizations/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// DeleteMultiple deletes multiple organizations in bulk.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Organizations/delete_organizations
func (s *OrganizationsService) DeleteMultiple(ctx context.Context, ids []int) (*Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, "/organizations", &DeleteMultipleOptions{
		Ids: arrayToString(ids, ","),
	}, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// OrganizationCreateOptions specifices the optional parameters to the
// OrganizationsService.Create method.
type OrganizationCreateOptions struct {
	Name      string    `json:"name"`
	OwnerID   uint      `json:"owner_id"`
	VisibleTo VisibleTo `json:"visible_to"`
	AddTime   Timestamp `json:"add_time"`
	Label     uint      `json:"label"`

	CustomFields map[string]interface{} `json:"-"`
}

func (d OrganizationCreateOptions) MarshalJSON() ([]byte, error) {
	fields := map[string]interface{}{}
	for k, v := range d.CustomFields {
		fields[k] = v
	}

	if d.Name != "" {
		fields["name"] = d.Name
	}
	if d.OwnerID != 0 {
		fields["owner_id"] = d.OwnerID
	}
	if !d.AddTime.IsZero() {
		fields["add_time"] = d.AddTime.FormatFull()
	}
	if d.VisibleTo != 0 {
		fields["visible_to"] = d.VisibleTo
	}

	return json.Marshal(fields)
}

// Create a new organizations.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Organizations/post_organizations
func (s *OrganizationsService) Create(ctx context.Context, opt *OrganizationCreateOptions) (*OrganizationResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/organizations", nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *OrganizationResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}
