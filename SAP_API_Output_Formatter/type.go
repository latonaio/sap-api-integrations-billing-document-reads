package sap_api_output_formatter

type BillingDocument struct {
	ConnectionKey   string `json:"connection_key"`
	Result          bool   `json:"result"`
	RedisKey        string `json:"redis_key"`
	Filepath        string `json:"filepath"`
	APISchema       string `json:"api_schema"`
	BillingDocument string `json:"billing_document"`
	Deleted         bool   `json:"deleted"`
}

type Header struct {
	BillingDocument            string `json:"BillingDocument"`
	BillingDocumentType        string `json:"BillingDocumentType"`
	SDDocumentCategory         string `json:"SDDocumentCategory"`
	BillingDocumentCategory    string `json:"BillingDocumentCategory"`
	CreationDate               string `json:"CreationDate"`
	LastChangeDate             string `json:"LastChangeDate"`
	SalesOrganization          string `json:"SalesOrganization"`
	DistributionChannel        string `json:"DistributionChannel"`
	Division                   string `json:"Division"`
	BillingDocumentDate        string `json:"BillingDocumentDate"`
	BillingDocumentIsCancelled bool   `json:"BillingDocumentIsCancelled"`
	CancelledBillingDocument   string `json:"CancelledBillingDocument"`
	IsExportDelivery           string `json:"IsExportDelivery"`
	TotalNetAmount             string `json:"TotalNetAmount"`
	TransactionCurrency        string `json:"TransactionCurrency"`
	TaxAmount                  string `json:"TaxAmount"`
	TotalGrossAmount           string `json:"TotalGrossAmount"`
	CustomerPriceGroup         string `json:"CustomerPriceGroup"`
	IncotermsClassification    string `json:"IncotermsClassification"`
	CustomerPaymentTerms       string `json:"CustomerPaymentTerms"`
	PaymentMethod              string `json:"PaymentMethod"`
	CompanyCode                string `json:"CompanyCode"`
	AccountingDocument         string `json:"AccountingDocument"`
	ExchangeRateDate           string `json:"ExchangeRateDate"`
	ExchangeRateType           string `json:"ExchangeRateType"`
	DocumentReferenceID        string `json:"DocumentReferenceID"`
	SoldToParty                string `json:"SoldToParty"`
	PartnerCompany             string `json:"PartnerCompany"`
	PurchaseOrderByCustomer    string `json:"PurchaseOrderByCustomer"`
	CustomerGroup              string `json:"CustomerGroup"`
	Country                    string `json:"Country"`
	CityCode                   string `json:"CityCode"`
	Region                     string `json:"Region"`
	CreditControlArea          string `json:"CreditControlArea"`
	OverallBillingStatus       string `json:"OverallBillingStatus"`
	AccountingPostingStatus    string `json:"AccountingPostingStatus"`
	AccountingTransferStatus   string `json:"AccountingTransferStatus"`
	InvoiceListStatus          string `json:"InvoiceListStatus"`
	BillingDocumentListType    string `json:"BillingDocumentListType"`
	BillingDocumentListDate    string `json:"BillingDocumentListDate"`
	ToItem                     string `json:"to_Item"`
	ToHeaderPartner            string `json:"to_Partner"`
}

type HeaderPartner struct {
	BillingDocument string `json:"BillingDocument"`
	PartnerFunction string `json:"PartnerFunction"`
	Customer        string `json:"Customer"`
	Supplier        string `json:"Supplier"`
}

type Item struct {
	BillingDocument              string `json:"BillingDocument"`
	BillingDocumentItem          string `json:"BillingDocumentItem"`
	SalesDocumentItemCategory    string `json:"SalesDocumentItemCategory"`
	ReturnItemProcessingType     string `json:"ReturnItemProcessingType"`
	CreationDate                 string `json:"CreationDate"`
	Division                     string `json:"Division"`
	Material                     string `json:"Material"`
	InternationalArticleNumber   string `json:"InternationalArticleNumber"`
	Batch                        string `json:"Batch"`
	MaterialGroup                string `json:"MaterialGroup"`
	Plant                        string `json:"Plant"`
	StorageLocation              string `json:"StorageLocation"`
	BillingDocumentItemText      string `json:"BillingDocumentItemText"`
	BillingQuantity              string `json:"BillingQuantity"`
	BillingQuantityUnit          string `json:"BillingQuantityUnit"`
	NetAmount                    string `json:"NetAmount"`
	TransactionCurrency          string `json:"TransactionCurrency"`
	GrossAmount                  string `json:"GrossAmount"`
	PricingDate                  string `json:"PricingDate"`
	TaxAmount                    string `json:"TaxAmount"`
	MaterialPricingGroup         string `json:"MaterialPricingGroup"`
	MainItemMaterialPricingGroup string `json:"MainItemMaterialPricingGroup"`
	BusinessArea                 string `json:"BusinessArea"`
	ProfitCenter                 string `json:"ProfitCenter"`
	WBSElement                   string `json:"WBSElement"`
	ControllingArea              string `json:"ControllingArea"`
	ProfitabilitySegment         string `json:"ProfitabilitySegment"`
	OrderID                      string `json:"OrderID"`
	CostCenter                   string `json:"CostCenter"`
	ReferenceSDDocument          string `json:"ReferenceSDDocument"`
	ReferenceSDDocumentItem      string `json:"ReferenceSDDocumentItem"`
	MatlAccountAssignmentGroup   string `json:"MatlAccountAssignmentGroup"`
	SalesDocument                string `json:"SalesDocument"`
	SalesDocumentItem            string `json:"SalesDocumentItem"`
	SDDocumentReason             string `json:"SDDocumentReason"`
	ShippingPoint                string `json:"ShippingPoint"`
	ToItemPartner                string `json:"to_Partner"`
	ToItemPricingElement         string `json:"to_PricingElement"`
}

type ItemPartner struct {
	BillingDocument     string `json:"BillingDocument"`
	BillingDocumentItem string `json:"BillingDocumentItem"`
	PartnerFunction     string `json:"PartnerFunction"`
	Customer            string `json:"Customer"`
	Supplier            string `json:"Supplier"`
}

type ToHeaderPartner struct {
	BillingDocument string `json:"BillingDocument"`
	PartnerFunction string `json:"PartnerFunction"`
	Customer        string `json:"Customer"`
	Supplier        string `json:"Supplier"`
}

type ToItem struct {
	BillingDocument              string `json:"BillingDocument"`
	BillingDocumentItem          string `json:"BillingDocumentItem"`
	SalesDocumentItemCategory    string `json:"SalesDocumentItemCategory"`
	ReturnItemProcessingType     string `json:"ReturnItemProcessingType"`
	CreationDate                 string `json:"CreationDate"`
	Division                     string `json:"Division"`
	Material                     string `json:"Material"`
	InternationalArticleNumber   string `json:"InternationalArticleNumber"`
	Batch                        string `json:"Batch"`
	MaterialGroup                string `json:"MaterialGroup"`
	Plant                        string `json:"Plant"`
	StorageLocation              string `json:"StorageLocation"`
	BillingDocumentItemText      string `json:"BillingDocumentItemText"`
	BillingQuantity              string `json:"BillingQuantity"`
	BillingQuantityUnit          string `json:"BillingQuantityUnit"`
	NetAmount                    string `json:"NetAmount"`
	TransactionCurrency          string `json:"TransactionCurrency"`
	GrossAmount                  string `json:"GrossAmount"`
	PricingDate                  string `json:"PricingDate"`
	TaxAmount                    string `json:"TaxAmount"`
	MaterialPricingGroup         string `json:"MaterialPricingGroup"`
	MainItemMaterialPricingGroup string `json:"MainItemMaterialPricingGroup"`
	BusinessArea                 string `json:"BusinessArea"`
	ProfitCenter                 string `json:"ProfitCenter"`
	WBSElement                   string `json:"WBSElement"`
	ControllingArea              string `json:"ControllingArea"`
	ProfitabilitySegment         string `json:"ProfitabilitySegment"`
	OrderID                      string `json:"OrderID"`
	CostCenter                   string `json:"CostCenter"`
	ReferenceSDDocument          string `json:"ReferenceSDDocument"`
	ReferenceSDDocumentItem      string `json:"ReferenceSDDocumentItem"`
	MatlAccountAssignmentGroup   string `json:"MatlAccountAssignmentGroup"`
	SalesDocument                string `json:"SalesDocument"`
	SalesDocumentItem            string `json:"SalesDocumentItem"`
	SDDocumentReason             string `json:"SDDocumentReason"`
	ShippingPoint                string `json:"ShippingPoint"`
	ToItemPartner                string `json:"to_Partner"`
	ToItemPricingElement         string `json:"to_PricingElement"`
}

type ToItemPartner struct {
	BillingDocument     string `json:"BillingDocument"`
	BillingDocumentItem string `json:"BillingDocumentItem"`
	PartnerFunction     string `json:"PartnerFunction"`
	Customer            string `json:"Customer"`
	Supplier            string `json:"Supplier"`
}

type ToItemPricingElement struct {
	BillingDocument             string `json:"BillingDocument"`
	BillingDocumentItem         string `json:"BillingDocumentItem"`
	PricingProcedureStep        string `json:"PricingProcedureStep"`
	PricingProcedureCounter     string `json:"PricingProcedureCounter"`
	ConditionType               string `json:"ConditionType"`
	PricingDateTime             string `json:"PricingDateTime"`
	ConditionCalculationType    string `json:"ConditionCalculationType"`
	ConditionBaseValue          string `json:"ConditionBaseValue"`
	ConditionRateValue          string `json:"ConditionRateValue"`
	ConditionCurrency           string `json:"ConditionCurrency"`
	ConditionQuantity           string `json:"ConditionQuantity"`
	ConditionQuantityUnit       string `json:"ConditionQuantityUnit"`
	ConditionCategory           string `json:"ConditionCategory"`
	PricingScaleType            string `json:"PricingScaleType"`
	ConditionRecord             string `json:"ConditionRecord"`
	ConditionSequentialNumber   string `json:"ConditionSequentialNumber"`
	TaxCode                     string `json:"TaxCode"`
	ConditionAmount             string `json:"ConditionAmount"`
	TransactionCurrency         string `json:"TransactionCurrency"`
	PricingScaleBasis           string `json:"PricingScaleBasis"`
	ConditionScaleBasisValue    string `json:"ConditionScaleBasisValue"`
	ConditionScaleBasisUnit     string `json:"ConditionScaleBasisUnit"`
	ConditionScaleBasisCurrency string `json:"ConditionScaleBasisCurrency"`
	ConditionIsManuallyChanged  bool   `json:"ConditionIsManuallyChanged"`
}
