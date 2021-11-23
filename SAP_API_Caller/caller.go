package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

func (c *SAPAPICaller) AsyncGetWorkCenter(WorkCenterInternalID, WorkCenterTypeCode, ValidityEndDate string) {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		c.WorkCenter(WorkCenterInternalID, WorkCenterTypeCode, ValidityEndDate)
		wg.Done()
	}()
	wg.Wait()
}

func (c *SAPAPICaller) WorkCenter(WorkCenterInternalID, WorkCenterTypeCode, ValidityEndDate string) {
	res, err := c.callWorkCenterSrvAPIRequirement("WorkCenterCapacity(WorkCenterInternalID='{WorkCenterInternalID}',WorkCenterTypeCode='{WorkCenterTypeCode}',CapacityCategoryAllocation='{CapacityCategoryAllocation}',CapacityInternalID='{CapacityInternalID}')/_Header", WorkCenterInternalID, WorkCenterTypeCode, ValidityEndDate)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) callWorkCenterSrvAPIRequirement(api, WorkCenterInternalID, WorkCenterTypeCode, ValidityEndDate string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "api_work_center/srvd_a2x/sap/workcenter/0001", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "WorkCenterInternalID, WorkCenterTypeCode, ValidityEndDate")
	params.Add("$filter", fmt.Sprintf("WorkCenterInternalID eq '%s' and WorkCenterTypeCode eq '%s' and ValidityEndDate eq '%s'", WorkCenterInternalID, WorkCenterTypeCode, ValidityEndDate))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}