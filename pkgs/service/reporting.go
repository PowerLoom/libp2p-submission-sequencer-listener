package service

import (
	"Listen/pkgs/common"
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

var ReportingInstance *ReportingService

type ReportingService struct {
	url    string
	client *http.Client
}

func InitializeReportingService(url string, timeout time.Duration) {
	ReportingInstance = &ReportingService{
		url: url, client: &http.Client{Timeout: timeout},
	}
	go ReportingInstance.StartListening()
}

func (s *ReportingService) StartListening() {
	for issue := range common.SequencerAlertsChannel {
		s.SendFailureNotification(issue)
	}
}

// sendPostRequest sends a POST request to the specified URL
func (s *ReportingService) SendFailureNotification(issue common.SequencerAlert) {
	jsonData, err := json.Marshal(issue)
	if err != nil {
		log.Errorln("Unable to marshal notification: ", issue)
		return
	}
	req, err := http.NewRequest("POST", s.url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Errorln("Error creating request: ", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Send the request
	resp, err := s.client.Do(req)
	if err != nil {
		log.Errorf("Error sending request for issue %s: %s\n", issue.String(), err)
		// Handle error in case of failure
		return
	}
	defer resp.Body.Close()

	// Here you can handle response or further actions
	log.Debugln("Reporting service response status: ", resp.Status)
}
