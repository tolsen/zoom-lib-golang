package metrics

import (
	"fmt"
	"github.com/himalayan-institute/zoom-lib-golang"
)

type (
	ListParticipantsOptions struct {
		MeetingId     string          `url"-""`
		Type          ListMeetingType `url:"type,omitempty"`
		PageSize      *int            `url:"page_size,omitempty"` // Default: 30, Maximum: 300
		NextPageToken string          `url:"next_page_token,omitempty"`
		IncludeFields string          `url:"include_fields,omitempty"`
	}

	ListParticipantsResponse struct {
		PageCount      int               `json:"page_count"`
		PageSize       int               `json:"page_size"`
		TotalRecords   int               `json:"total_records"`
		NextPageToken  string            `json:"next_page_token"`
		Participants   []ListParticipant `json:"participants"`
		TrackingFields []TrackingField   `json:"tracking_fields"`
	}

	ListParticipant struct {
		ID                 string     `json:"id"`
		UserId             string     `json:"user_id"`
		UserName           string     `json:"user_name"`
		IpAddress          string     `json:"ip_address"`
		Location           string     `json:"location"`
		NetworkType        string     `json:"network_type"`
		DataCenter         string     `json:"data_center"`
		ConnectionType     string     `json:"connection_type"`
		JoinTime           *zoom.Time `json:"join_time"`
		LeaveTime          *zoom.Time `json:"leave_time"`
		ShareApplication   bool       `json:"share_application"`
		ShareDesktop       bool       `json:"share_desktop"`
		ShareWhiteboard    bool       `json:"share_whiteboard"`
		Recording          bool       `json:"recording"`
		PcName             string     `json:"pc_name"`
		Version            string     `json:"version"`
		InRoomParticipants int        `json:"in_room_participants"`
		LeaveReason        string     `json:"leave_reason"`
		Email              string     `json:"email"`
		RegistrantId       string     `json:"registrant_id"`
		AudioQuality       string     `json:"audio_quality"`
		VideoQuality      string     `json:"video_quality"`
		ScreenShareQuality string     `json:"screen_share_quality"`

		// Zoom has disabled the following fields as of 1/31/2021 for privacy reasons
		// Device             string     `json:"device"`
		// Microphone         string     `json:"microphone"`
		// Speaker            string     `json:"speaker"`
		// Camera             string     `json:"camera"`
		// Domain             string     `json:"domain"`
		// MacAddr            string     `json:"mac_addr"`
		// HarddiskId         string     `json:"harddisk_id"`
	}
)

const (
	// GET /metrics/meetings/{meetingId}/participants
	ListParticipantsPath = "/metrics/meetings/%v/participants"
)

func ListParticipants(c *zoom.Client, opts ListParticipantsOptions) (ListParticipantsResponse, error) {
	var ret = ListParticipantsResponse{}
	return ret, c.RequestV2(zoom.RequestV2Opts{
		Method:        zoom.Get,
		Path:          fmt.Sprintf(ListParticipantsPath, opts.MeetingId),
		URLParameters: &opts,
		Ret:           &ret,
	})
}
