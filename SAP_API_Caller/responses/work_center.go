package responses

type WorkCenter struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			WorkCenterInternalID         string `json:"WorkCenterInternalID"`
			WorkCenterTypeCode           string `json:"WorkCenterTypeCode"`
			WorkCenter                   string `json:"WorkCenter"`
			WorkCenterDesc               string `json:"WorkCenter_desc"`
			Plant                        string `json:"Plant"`
			WorkCenterCategoryCode       string `json:"WorkCenterCategoryCode"`
			WorkCenterResponsible        string `json:"WorkCenterResponsible"`
			SupplyArea                   string `json:"SupplyArea"`
			WorkCenterUsage              string `json:"WorkCenterUsage"`
			MatlCompIsMarkedForBackflush bool   `json:"MatlCompIsMarkedForBackflush"`
			WorkCenterLocation           string `json:"WorkCenterLocation"`
			CapacityInternalID           string `json:"CapacityInternalID"`
			CapacityCategoryCode         string `json:"CapacityCategoryCode"`
			ValidityStartDate            string `json:"ValidityStartDate"`
			ValidityEndDate              string `json:"ValidityEndDate"`
			WorkCenterIsToBeDeleted      bool   `json:"WorkCenterIsToBeDeleted"`
		} `json:"results"`
	} `json:"d"`
}
