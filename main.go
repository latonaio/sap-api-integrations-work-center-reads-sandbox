package main

import (
	sap_api_caller "sap-api-integrations-work-center-reads/SAP_API_Caller"
	"sap-api-integrations-work-center-reads/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs//SDC_Work_Center_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
        "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata4/sap/", l,  // Work Center は URL の フォーマットが若干異なる ※要検証
    )

    caller.AsyncGetWorkCenter(
        inoutSDC.WorkCenter.WorkCenterInternalID,
        inoutSDC.WorkCenter.WorkCenterTypeCode,
        inoutSDC.WorkCenter.ValidityEndDate,
    )
}
