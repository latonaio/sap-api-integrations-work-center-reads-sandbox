# sap-api-integrations-work-center-reads  
sap-api-integrations-work-center-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 作業区 データを取得するマイクロサービスです。  
sap-api-integrations-work-center-reads には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-work-center-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/OP_WORKCENTER_0001/overview  

## 動作環境
sap-api-integrations-work-center-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。   
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。   
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須） 

## クラウド環境での利用  
sap-api-integrations-work-center-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-work-center-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_WORKCENTER_0001/overview  
* APIサービス名(=baseURL): api_work_center/srvd_a2x/sap/workcenter/0001

## 本レポジトリ に 含まれる API名
sap-api-integrations-work-center-reads には、次の API をコールするためのリソースが含まれています。  

* WorkCenterCapacity(WorkCenterInternalID='{WorkCenterInternalID}',WorkCenterTypeCode='{WorkCenterTypeCode}',CapacityCategoryAllocation='{CapacityCategoryAllocation}',CapacityInternalID='{CapacityInternalID}')/_Header（作業区マスタ）

## API への 値入力条件 の 初期値
sap-api-integrations-work-center-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

## SDC レイアウト

* inoutSDC.WorkCenter.WorkCenterInternalID（作業区内部ID）
* inoutSDC.WorkCenter.WorkCenterTypeCode（作業区タイプ）

## SAP API Business Hub における API サービス の バージョン と バージョン におけるデータレイアウトの相違

SAP API Business Hub における API サービス のうちの 殆どの API サービス のBASE URLのフォーマットは、"API_(リポジトリ名)_SRV" であり、殆どの API サービス 間 の データレイアウトは統一されています。   
従って、Latona および AION における リソースにおいても、データレイアウトが統一されています。    
一方、本レポジトリ に関わる API である Work Center のサービスは、BASE URLのフォーマットが他のAPIサービスと異なります。      
その結果、本レポジトリ内の一部のAPIのデータレイアウトが、他のAPIサービスのものと異なっています。  

#### BASE URLが "API_(リポジトリ名)_SRV" のフォーマットである API サービス の データレイアウト（=responses）  
BASE URLが "API_{リポジトリ名}_SRV" のフォーマットであるAPIサービスのデータレイアウト（=responses）は、例えば、次の通りです。  
```
type ToProductionOrderItem struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ManufacturingOrder             string      `json:"ManufacturingOrder"`
			ManufacturingOrderItem         string      `json:"ManufacturingOrderItem"`
			ManufacturingOrderCategory     string      `json:"ManufacturingOrderCategory"`
			ManufacturingOrderType         string      `json:"ManufacturingOrderType"`
			IsCompletelyDelivered          bool        `json:"IsCompletelyDelivered"`
			Material                       string      `json:"Material"`
			ProductionPlant                string      `json:"ProductionPlant"`
			Plant                          string      `json:"Plant"`
			MRPArea                        string      `json:"MRPArea"`
			QuantityDistributionKey        string      `json:"QuantityDistributionKey"`
			MaterialGoodsReceiptDuration   string      `json:"MaterialGoodsReceiptDuration"`
			StorageLocation                string      `json:"StorageLocation"`
			Batch                          string      `json:"Batch"`
			InventoryUsabilityCode         string      `json:"InventoryUsabilityCode"`
			GoodsRecipientName             string      `json:"GoodsRecipientName"`
			UnloadingPointName             string      `json:"UnloadingPointName"`
			MfgOrderItemPlndDeliveryDate   string      `json:"MfgOrderItemPlndDeliveryDate"`
			MfgOrderItemActualDeliveryDate string      `json:"MfgOrderItemActualDeliveryDate"`
			ProductionUnit                 string      `json:"ProductionUnit"`
			MfgOrderItemPlannedTotalQty    string      `json:"MfgOrderItemPlannedTotalQty"`
			MfgOrderItemPlannedScrapQty    string      `json:"MfgOrderItemPlannedScrapQty"`
			MfgOrderItemGoodsReceiptQty    string      `json:"MfgOrderItemGoodsReceiptQty"`
			MfgOrderItemActualDeviationQty string      `json:"MfgOrderItemActualDeviationQty"`
		} `json:"results"`
	} `json:"d"`
}

```

#### BASE URL が "api_work_center/srvd_a2x/sap/workcenter/0001" である Work Center の APIサービス の データレイアウト（=responses）  
BASE URL が "api_work_center/srvd_a2x/sap/workcenter/0001" である Work Center の APIサービス の データレイアウト（=responses）は、例えば、次の通りです。  

```
type WorkCenter struct {
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
}

```
このように、BASE URLが "API_(リポジトリ名)_SRV" のフォーマットである API サービス の データレイアウトと、 Work Center の データレイアウトは、D、Results、Metadata の配列構造を持っているか持っていないかという点が異なります。  