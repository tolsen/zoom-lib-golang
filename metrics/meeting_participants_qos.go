package metrics

import (
	"github.com/himalayan-institute/zoom-lib-golang"
)

type (

	ListParticipantsQosOptions struct {
		MeetingId     string          `url"-""`
		Type          ListMeetingType `url:"type,omitempty"`
		PageSize      *int            `url:"page_size,omitempty"` // Default: 1, Maximum: 10
		NextPageToken string          `url:"next_page_token,omitempty"`
	}

	ListParticipantsQosResponse struct {
		PageCount     int               `json:"page_count"`
		PageSize      int               `json:"page_size"`
		TotalRecords  int               `json:"total_records"`
		NextPageToken string            `json:"next_page_token"`
		Participants  []ListParticipantQos `json:"participants"`
	}

	ListParticipantQos struct {
		UserId             string     `json:"user_id"`
		UserName           string     `json:"user_name"`
		IpAddress          string     `json:"ip_address"`
		Location           string     `json:"location"`
		JoinTime           *zoom.Time `json:"join_time"`
		LeaveTime          *zoom.Time `json:"leave_time"`
		PcName             string     `json:"pc_name"`
		Version            string     `json:"version"`
		UserQos []ListParticipantQosUserQos `json:"user_qos"`
	}

	ListParticipantQosUserQos struct {

	}



)