package pipedrive

import (
	"fmt"
	"net/http"
)

type PersonsService service

type Person struct {
	ID        int `json:"id"`
	CompanyID int `json:"company_id"`
	OwnerID   struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		HasPic     bool   `json:"has_pic"`
		PicHash    string `json:"pic_hash"`
		ActiveFlag bool   `json:"active_flag"`
		Value      int    `json:"value"`
	} `json:"owner_id"`
	OrgID                       interface{} `json:"org_id"`
	Name                        string      `json:"name"`
	FirstName                   string      `json:"first_name"`
	LastName                    string      `json:"last_name"`
	OpenDealsCount              int         `json:"open_deals_count"`
	RelatedOpenDealsCount       int         `json:"related_open_deals_count"`
	ClosedDealsCount            int         `json:"closed_deals_count"`
	RelatedClosedDealsCount     int         `json:"related_closed_deals_count"`
	ParticipantOpenDealsCount   int         `json:"participant_open_deals_count"`
	ParticipantClosedDealsCount int         `json:"participant_closed_deals_count"`
	EmailMessagesCount          int         `json:"email_messages_count"`
	ActivitiesCount             int         `json:"activities_count"`
	DoneActivitiesCount         int         `json:"done_activities_count"`
	UndoneActivitiesCount       int         `json:"undone_activities_count"`
	ReferenceActivitiesCount    int         `json:"reference_activities_count"`
	FilesCount                  int         `json:"files_count"`
	NotesCount                  int         `json:"notes_count"`
	FollowersCount              int         `json:"followers_count"`
	WonDealsCount               int         `json:"won_deals_count"`
	RelatedWonDealsCount        int         `json:"related_won_deals_count"`
	LostDealsCount              int         `json:"lost_deals_count"`
	RelatedLostDealsCount       int         `json:"related_lost_deals_count"`
	ActiveFlag                  bool        `json:"active_flag"`
	Phone                       []struct {
		Value   string `json:"value"`
		Primary bool   `json:"primary"`
	} `json:"phone"`
	Email []struct {
		Value   string `json:"value"`
		Primary bool   `json:"primary"`
	} `json:"email"`
	FirstChar                       string      `json:"first_char"`
	UpdateTime                      string      `json:"update_time"`
	AddTime                         string      `json:"add_time"`
	VisibleTo                       string      `json:"visible_to"`
	PictureID                       interface{} `json:"picture_id"`
	NextActivityDate                interface{} `json:"next_activity_date"`
	NextActivityTime                interface{} `json:"next_activity_time"`
	NextActivityID                  interface{} `json:"next_activity_id"`
	LastActivityID                  int         `json:"last_activity_id"`
	LastActivityDate                string      `json:"last_activity_date"`
	TimelineLastActivityTime        interface{} `json:"timeline_last_activity_time"`
	TimelineLastActivityTimeByOwner interface{} `json:"timeline_last_activity_time_by_owner"`
	LastIncomingMailTime            interface{} `json:"last_incoming_mail_time"`
	LastOutgoingMailTime            interface{} `json:"last_outgoing_mail_time"`
	OrgName                         interface{} `json:"org_name"`
	OwnerName                       string      `json:"owner_name"`
	CcEmail                         string      `json:"cc_email"`
}

type PersonsRespose struct {
	Success        bool           `json:"success"`
	Data           []Person       `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

type PersonResponse struct {
	Success        bool           `json:"success"`
	Data           Person         `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

type PersonAddFollowerResponse struct {
	Success bool `json:"success"`
	Data    struct {
		UserID   int    `json:"user_id"`
		ID       int    `json:"id"`
		PersonID int    `json:"person_id"`
		AddTime  string `json:"add_time"`
	} `json:"data"`
}

type PersonCreateOptions struct {
	Name      string    `url:"name"`
	OwnerID   uint      `url:"owner_id"`
	OrgID     uint      `url:"org_id"`
	Email     string    `url:"email"`
	Phone     string    `url:"phone"`
	VisibleTo VisibleTo `url:"visible_to"`
	AddTime   Timestamp `url:"add_time"`
}

type PersonUpdateOptions struct {
	Name      string    `url:"name"`
	OwnerID   uint      `url:"owner_id"`
	OrgID     uint      `url:"org_id"`
	Email     string    `url:"email"`
	Phone     string    `url:"phone"`
	VisibleTo VisibleTo `url:"visible_to"`
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/get_persons
func (s *PersonsService) List() (*PersonsRespose, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/persons", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *PersonsRespose

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/post_persons_id_followers
func (s *PersonsService) AddFollower(id uint, userID uint) (*PersonAddFollowerResponse, *Response, error) {
	uri := fmt.Sprintf("/persons/%v/followers", id)
	req, err := s.client.NewRequest(http.MethodPost, uri, nil, struct {
		UserID uint `url:"user_id"`
	}{
		userID,
	})

	if err != nil {
		return nil, nil, err
	}

	var record *PersonAddFollowerResponse

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/post_persons
func (s *PersonsService) Create(opt *PersonCreateOptions) (*PersonResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodPost, "/persons", nil, struct {
		Name      string    `url:"name"`
		OwnerID   uint      `url:"owner_id"`
		OrgID     uint      `url:"org_id"`
		Email     string    `url:"email"`
		Phone     string    `url:"phone"`
		VisibleTo VisibleTo `url:"visible_to"`
		AddTime   string    `url:"add_time"`
	}{
		opt.Name,
		opt.OwnerID,
		opt.OrgID,
		opt.Email,
		opt.Phone,
		opt.VisibleTo,
		opt.AddTime.FormatFull(),
	})

	if err != nil {
		return nil, nil, err
	}

	var record *PersonResponse

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/put_persons_id
func (s *PersonsService) Update(id uint, opt *PersonUpdateOptions) (*PersonResponse, *Response, error) {
	uri := fmt.Sprintf("/persons/%v", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, nil, opt)

	if err != nil {
		return nil, nil, err
	}

	var record *PersonResponse

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/put_persons_id_merge
func (s *PersonsService) Merge(id uint, mergeWithID uint) (*PersonResponse, *Response, error) {
	uri := fmt.Sprintf("/persons/%v/merge", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, nil, struct {
		MergeWithID uint `url:"merge_with_id"`
	}{
		mergeWithID,
	})

	if err != nil {
		return nil, nil, err
	}

	var record *PersonResponse

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/delete_persons_id_followers_follower_id
func (s *PersonsService) DeleteFollower(id uint, followerID uint) (*Response, error) {
	uri := fmt.Sprintf("/persons/%v/followers/%v", id, followerID)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/delete_persons_id
func (s *PersonsService) Delete(id uint) (*Response, error) {
	uri := fmt.Sprintf("/persons/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/delete_persons_id_picture
func (s *PersonsService) DeletePicture(id uint) (*Response, error) {
	uri := fmt.Sprintf("/persons/%v/picture", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/delete_persons
func (s *PersonsService) DeleteMultiple(ids []int) (*Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, "/persons", &DeleteMultipleOptions{
		Ids: arrayToString(ids, ","),
	}, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
