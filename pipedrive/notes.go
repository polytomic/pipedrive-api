package pipedrive

import (
	"fmt"
	"net/http"
	"time"
)

type NotesService service

type JsonTime time.Time

type Note struct {
	Id                       int      `json:"id,omitempty"`
	UserId                   int      `json:"user_id,omitempty"`
	DealId                   int      `json:"deal_id,omitempty"`
	PersonId                 int      `json:"person_id,omitempty"`
	OrgId                    int      `json:"org_id,omitempty"`
	Content                  string   `json:"content,omitempty"`
	AddTime                  JsonTime `json:"add_time,omitempty"`
	UpdateTime               JsonTime `json:"update_time,omitempty"`
	ActiveFlag               bool     `json:"active_flag,omitempty"`
	PinnedToDealFlag         bool     `json:"pinned_to_deal_flag,omitempty"`
	PinnedToPersonFlag       bool     `json:"pinned_to_person_flag,omitempty"`
	PinnedToOrganizationFlag bool     `json:"pinned_to_organization_flag,omitempty"`
	LastUpdateUserId         int      `json:"last_update_user_id,omitempty"`
}

type Notes struct {
	Success        bool           `json:"success,omitempty"`
	Data           []Note         `json:"data,omitempty"`
	AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

type SingleNote struct {
	Success bool `json:"success,omitempty"`
	Data    Note `json:"data,omitempty"`
}

// Returns all notes
// https://developers.pipedrive.com/docs/api/v1/#!/Notes/get_notes
func (s *NotesService) List() (*Notes, *http.Response, error) {
	uri := s.client.CreateRequestUrl("/notes")
	req, err := s.client.NewRequest("GET", uri, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *Notes

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Returns details about a specific note.
// https://developers.pipedrive.com/docs/api/v1/#!/Notes/get_notes_id
func (s *NotesService) GetById(id int) (*SingleNote, *http.Response, error) {
	uri := s.client.CreateRequestUrl(fmt.Sprintf("/notes/%v", id))
	req, err := s.client.NewRequest("GET", uri, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *SingleNote

	resp, err := s.client.Do(req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}
