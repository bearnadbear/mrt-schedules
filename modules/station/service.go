package station

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bearnadbear/mrt-schedules/cummon/client"
)

type Service interface {
	GetAllStation() (response []StationResponse, err error)
}

type service struct {
	client *http.Client
}

func (s *service) GetAllStation() (response []StationResponse, err error) {
	url := "https://jakartamrt.co.id/id/val/stasiuns?_type=query"

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return
	}

	var stations []Station
	err = json.Unmarshal(byteResponse, &stations)

	for _, item := range stations {
		response = append(response, StationResponse{
			Id:   item.Id,
			Name: item.Name,
		})
	}

	return
}

func NewService() Service {
	return &service{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
