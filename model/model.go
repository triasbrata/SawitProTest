package model

type ErrorResponse struct {
	StatusCode int64  `json:"statusCode"`
	Success    bool   `json:"success"`
	Message    string `json:"message"`
}

type EstateRequest struct {
	Width  int `json:"width"`
	Length int `json:"length"`
}

type EstateResponseSuccess struct {
	ID string `json:"id"`
}

type TreeRequest struct {
	X      int64 `json:"x"`
	Y      int64 `json:"y"`
	Height int64 `json:"height"`
}

type TreeResponseSuccess struct {
	ID string `json:"id"`
}

type EstateStatsResponse struct {
	Count  int     `json:"count"`
	Max    int64   `json:"max"`
	Min    int64   `json:"min"`
	Median float64 `json:"median"`
}
type DronePlanResponseSuccess struct {
	Distance int64                        `json:"distance"`
	Rest     *DronePlanResponseSuccessRes `json:"rest,omitempty"`
}
type DronePlanResponseSuccessRes struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}
