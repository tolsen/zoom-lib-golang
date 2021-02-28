package metrics

import (
	"fmt"

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
		PageCount     int                  `json:"page_count"`
		PageSize      int                  `json:"page_size"`
		TotalRecords  int                  `json:"total_records"`
		NextPageToken string               `json:"next_page_token"`
		Participants  []ListParticipantQos `json:"participants"`
	}

	ListParticipantQos struct {
		UserId    string                      `json:"user_id"`
		UserName  string                      `json:"user_name"`
		IpAddress string                      `json:"ip_address"`
		Location  string                      `json:"location"`
		JoinTime  *zoom.Time                  `json:"join_time"`
		LeaveTime *zoom.Time                  `json:"leave_time"`
		PcName    string                      `json:"pc_name"`
		Version   string                      `json:"version"`
		UserQos   []ListParticipantQosUserQos `json:"user_qos"`
	}

	ListParticipantQosUserQos struct {
		DateTime    *zoom.Time `json:"date_time"`
		AudioInput  AudioQos   `json:"audio_input"`
		AudioOutput AudioQos   `json:"audio_output"`
		VideoInput  VideoQos   `json:"video_input"`
		VideoOutput VideoQos   `json:"video_output"`

		// not interested
		//AsInput VideoQos `json:"as_input"`
		//AsOutput VideoQos `json:"as_output"`

		// more stuff I am not interested in not here
	}

	AudioQos struct {
		Bitrate string `json:"bitrate"`
		Latency string `json:"latency"`
		Jitter  string `json:"jitter"`
		AvgLoss string `json:"avg_loss"`
		MaxLoss string `json:"max_loss"`
	}

	VideoQos struct {
		Resolution string `json:"resolution"`
		FrameRate  string `json:"frame_rate"`
		AudioQos
	}
)

const (
	// GET /metrics/meetings/{meetingId}/participants/qos
	ListParticipantsQosPath = "/metrics/meetings/%v/participants/qos"
)

func ListParticipantsQos(c *zoom.Client, opts ListParticipantsQosOptions) (ListParticipantsQosResponse, error) {
	var ret = ListParticipantsQosResponse{}
	return ret, c.RequestV2(zoom.RequestV2Opts{
		Method:        zoom.Get,
		Path:          fmt.Sprintf(ListParticipantsQosPath, opts.MeetingId),
		URLParameters: &opts,
		Ret:           &ret,
	})
}
