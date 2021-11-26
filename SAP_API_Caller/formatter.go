package sap_api_caller

type WorkCenterReads struct {
	 ConnectionKey    string `json:"connection_key"`
	 Result           bool   `json:"result"`
	 RedisKey         string `json:"redis_key"`
	 Filepath         string `json:"filepath"`
	 APISchema        string `json:"api_schema"`
	 WorkCenter       string `json:"work_center_code"`
	 Deleted          string `json:"deleted"`
}

type WorkCenter struct {
     WorkCenterInternalID         int    `json:"WorkCenterInternalID"`
     WorkCenterTypeCode           string `json:"WorkCenterTypeCode"`
     WorkCenter                   string `json:"WorkCenter"`
     WorkCenterDesc               string `json:"WorkCenter_desc"`
     Plant                        string `json:"Plant"`
     WorkCenterCategoryCode       string `json:"WorkCenterCategoryCode"`
     WorkCenterResponsible        string `json:"WorkCenterResponsible"`
     SupplyArea                   string `json:"SupplyArea"`
     WorkCenterUsage              string `json:"WorkCenterUsage"`
     MatlCompIsMarkedForBackflush string `json:"MatlCompIsMarkedForBackflush"`
     WorkCenterLocation           string `json:"WorkCenterLocation"`
     CapacityInternalID           int    `json:"CapacityInternalID"`
     CapacityCategoryCode         string `json:"CapacityCategoryCode"`
     ValidityStartDate            string `json:"ValidityStartDate"`
     ValidityEndDate              string `json:"ValidityEndDate"`
     WorkCenterIsToBeDeleted      string `json:"WorkCenterIsToBeDeleted"`
}
