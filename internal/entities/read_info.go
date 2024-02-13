package entities

import "encoding/json"

type ReadInfo struct {
	Status           ReadingState `json:"status"`
	LastChapterRead  int          `json:"lastChapterRead"`
	WeeklyReleaseDay WeekDay      `json:"releaseDay"`
	ReadIn           []ReadIn     `json:"readIn"`
}

func (r *ReadInfo) MarshalJSON() ([]byte, error) {
	urls := make([]string, len(r.ReadIn))
	for i, x := range r.ReadIn {
		urls[i] = x.Url
	}

	return json.Marshal(
		&struct {
			Status           ReadingState `json:"status"`
			LastChapterRead  int          `json:"lastChapterRead"`
			WeeklyReleaseDay WeekDay      `json:"releaseDay"`
			ReadIn           []string     `json:"readIn"`
		}{
			Status:           r.Status,
			LastChapterRead:  r.LastChapterRead,
			WeeklyReleaseDay: r.WeeklyReleaseDay,
			ReadIn:           urls,
		},
	)

}
