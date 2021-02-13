package metrics

import "github.com/himalayan-institute/zoom-lib-golang"

type (
	ListMeetingType string

	ListMeetingsOptions struct {
		Type          ListMeetingType `url:"type,omitempty"`
		From          string          `url:"from,omitempty"`
		To            string          `url:"to,omitempty"`
		PageSize      *int            `url:"page_size,omitempty"` // Default: 30, Maximum: 300
		NextPageToken string          `url:"next_page_token,omitempty"`
		IncludeFields string          `url:"include_fields,omitempty"`
	}

	ListMeetingsResponse struct {
		From           string          `json:"from"`
		To             string          `json:"to"`
		PageCount      int             `json:"page_count"`
		PageSize       int             `json:"page_size"`
		TotalRecords   int             `json:"total_records"`
		NextPageToken  string          `json:"next_page_token"`
		Meetings       []ListMeeting   `json:"meetings"`
		TrackingFields []TrackingField `json:"tracking_fields"`
	}

	ListMeeting struct {
		UUID         string     `json:"uuid"`
		ID           int        `json:"id"`
		Topic        string     `json:"topic"`
		Host         string     `json:"host"`
		Email        string     `json:"email"`
		UserType     string     `json:"user_type"`
		StartTime    *zoom.Time `json:"start_time"`
		EndTime      *zoom.Time `json:"end_time"`
		Duration     string     `json:"duration"`
		Participants int        `json:"participants"`
		HasRecording bool       `json:"has_recording"`
	}

	TrackingField struct {
		Field string `json:"field"`
		Value string `json:"value"`
	}
)

const (
	ListMeetingsPath = "/metrics/meetings"

	ListMeetingTypePast    ListMeetingType = "past"
	ListMeetingTypePastOne ListMeetingType = "pastOne"
	ListMeetingTypeLive    ListMeetingType = "live" // DEFAULT
)

func ListMeetings(c *zoom.Client, opts ListMeetingsOptions) (ListMeetingsResponse, error) {
	var ret = ListMeetingsResponse{}
	return ret, c.RequestV2(zoom.RequestV2Opts{
		Method:        zoom.Get,
		Path:          ListMeetingsPath,
		URLParameters: &opts,
		Ret:           &ret,
	})
}
