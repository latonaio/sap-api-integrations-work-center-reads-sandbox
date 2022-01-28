package main

import (
	sap_api_caller "sap-api-integrations-work-center-reads/SAP_API_Caller"
	"sap-api-integrations-work-center-reads/sap_api_input_reader"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

func main() {
	l := logger.NewLogger()
	fr := sap_api_input_reader.NewFileReader()
	inoutSDC := fr.ReadSDC("./Inputs//SDC_Work_Center_sample.json")
	caller := sap_api_caller.NewSAPAPICaller(
		"https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata4/sap/", l, // Work Center は URL の フォーマットが若干異なる
	)

	accepter := inoutSDC.Accepter
	if len(accepter) == 0 || accepter[0] == "All" {
		accepter = []string{
			"WorkCenter",
		}
	}

	caller.AsyncGetWorkCenter(
		inoutSDC.WorkCenter.WorkCenterInternalID,
		inoutSDC.WorkCenter.WorkCenterTypeCode,
		accepter,
	)
}
