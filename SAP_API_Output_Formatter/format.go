package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-work-center-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
)

func ConvertToWorkCenter(raw []byte, l *logger.Logger) *WorkCenter {
	pm := &responses.WorkCenter{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		l.Error(err)
		return nil
	}
	if len(pm.D.Results) == 0 {
		l.Error("Result data is not exist.")
		return nil
	}
	if len(pm.D.Results) > 1 {
		l.Error("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &WorkCenter{
		WorkCenterInternalID         data.WorkCenterInternalID,
		WorkCenterTypeCode           data.WorkCenterTypeCode,
		WorkCenter                   data.WorkCenter,
		WorkCenterDesc               data.WorkCenter_desc,
		Plant                        data.Plant,
		WorkCenterCategoryCode       data.WorkCenterCategoryCode,
		WorkCenterResponsible        data.WorkCenterResponsible,
		SupplyArea                   data.SupplyArea,
		WorkCenterUsage              data.WorkCenterUsage,
		MatlCompIsMarkedForBackflush data.MatlCompIsMarkedForBackflush,
		WorkCenterLocation           data.WorkCenterLocation,
		CapacityInternalID           data.CapacityInternalID,
		CapacityCategoryCode         data.CapacityCategoryCode,
		ValidityStartDate            data.ValidityStartDate,
		ValidityEndDate              data.ValidityEndDate,
		WorkCenterIsToBeDeleted      data.WorkCenterIsToBeDeleted,
	}
}
