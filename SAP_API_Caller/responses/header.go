package responses

type Header struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
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
			ToHeaderPartner            struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Partner"`
			ToItem struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_Item"`
		} `json:"results"`
	} `json:"d"`
}
