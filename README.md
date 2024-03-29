# sap-api-integrations-billing-document-reads
sap-api-integrations-billing-document-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 請求伝票データを取得するマイクロサービスです。    
sap-api-integrations-billing-document-reads には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-billing-document-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/OP_API_BILLING_DOCUMENT_SRV_0001/overview

## 動作環境  
sap-api-integrations-billing-document-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）　　

## クラウド環境での利用
sap-api-integrations-billing-document-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-billing-document-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_BILLING_DOCUMENT_SRV_0001/overview    
* APIサービス名(=baseURL): API_BILLING_DOCUMENT_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-billing-document-reads には、次の API をコールするためのリソースが含まれています。  

* A_BillingDocument（請求伝票 - ヘッダ）※請求伝票の詳細データを取得するために、ToItem、ToHeaderPartner、ToItemPartner、ToItemPricingElement、と合わせて利用されます。
* A_BillingDocumentPartner（請求伝票 - ヘッダ取引先）
* A_BillingDocumentItem（請求伝票 - 明細）※請求伝票明細の詳細データを取得するために、ToItemPartner、ToItemPricingElementと合わせて利用されます。
* A_BillingDocumentItemPartner（請求伝票 - 明細取引先）
* ToHeaderPartner（請求伝票 - ヘッダ取引先 ※To）
* ToItem（請求伝票 - 明細 ※To）
* ToItemPartner（請求伝票 - 明細取引先 ※To）
* ToItemPricingElement（請求伝票 - 明細価格条件）

## API への 値入力条件 の 初期値
sap-api-integrations-billing-document-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.BillingDocument.BillingDocument（請求伝票）
* inoutSDC.BillingDocument.HeaderPartner.PartnerFunction（ヘッダ取引先機能）
* inoutSDC.BillingDocument.BillingDocumentItem.BillingDocumentItem（請求伝票明細）
* inoutSDC.BillingDocument.BillingDocumentItem.ItemPartner.PartnerFunction（明細取引先機能）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。

```
	"api_schema": "SAPBillingDocumentReads",
	"accepter": ["Header"],	
	"billing_document": "90000000",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "SAPBillingDocumentReads",
	"accepter": ["All"],	
	"billing_document": "90000000",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetBillingDocument(billingDocument, headerPartnerFunction, billingDocumentItem, itemPartnerFunction string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(billingDocument)
				wg.Done()
			}()
		case "HeaderPartner":
			func() {
				c.HeaderPartner(billingDocument, headerPartnerFunction)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(billingDocument, billingDocumentItem)
				wg.Done()
			}()
		case "ItemPartner":
			func() {
				c.ItemPartner(billingDocument, billingDocumentItem, itemPartnerFunction)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```
## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 請求伝票 の ヘッダデータ が取得された結果の JSON の例です。  
以下の項目のうち、"BillingDocument" ～ "to_Partner" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-billing-document-reads/SAP_API_Caller/caller.go#L70",
	"function": "sap-api-integrations-billing-document-reads/SAP_API_Caller.(*SAPAPICaller).Header",
	"level": "INFO",
	"message": [
		{
			"BillingDocument": "90000000",
			"BillingDocumentType": "F2",
			"SDDocumentCategory": "M",
			"BillingDocumentCategory": "L",
			"CreationDate": "2022-09-15",
			"LastChangeDate": "",
			"SalesOrganization": "0001",
			"DistributionChannel": "01",
			"Division": "01",
			"BillingDocumentDate": "1998-02-01",
			"BillingDocumentIsCancelled": false,
			"CancelledBillingDocument": "",
			"IsExportDelivery": "",
			"TotalNetAmount": "1000.00",
			"TransactionCurrency": "EUR",
			"TaxAmount": "0.00",
			"TotalGrossAmount": "1000.00",
			"CustomerPriceGroup": "01",
			"IncotermsClassification": "FH",
			"CustomerPaymentTerms": "0001",
			"PaymentMethod": "",
			"CompanyCode": "0001",
			"AccountingDocument": "",
			"ExchangeRateDate": "1998-02-01",
			"ExchangeRateType": "",
			"DocumentReferenceID": "Test",
			"SoldToParty": "1",
			"PartnerCompany": "",
			"PurchaseOrderByCustomer": "Test",
			"CustomerGroup": "",
			"Country": "JP",
			"CityCode": "",
			"Region": "13",
			"CreditControlArea": "0001",
			"OverallBillingStatus": "B",
			"AccountingPostingStatus": "A",
			"AccountingTransferStatus": "F",
			"InvoiceListStatus": "",
			"BillingDocumentListType": "LR",
			"BillingDocumentListDate": "",
			"to_Item": "http://100.21.57.120:8080/sap/opu/odata/sap/API_BILLING_DOCUMENT_SRV/A_BillingDocument('90000000')/to_Item",
			"to_Partner": "http://100.21.57.120:8080/sap/opu/odata/sap/API_BILLING_DOCUMENT_SRV/A_BillingDocument('90000000')/to_Partner"
		}
	],
	"time": "2022-09-15T08:52:03+09:00"
}
```
