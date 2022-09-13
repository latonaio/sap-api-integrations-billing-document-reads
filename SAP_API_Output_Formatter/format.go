package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-billing-document-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) ([]Header, error) {
	pm := &responses.Header{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	header := make([]Header, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		header = append(header, Header{
			BillingDocument:            data.BillingDocument,
			BillingDocumentType:        data.BillingDocumentType,
			SDDocumentCategory:         data.SDDocumentCategory,
			BillingDocumentCategory:    data.BillingDocumentCategory,
			CreationDate:               data.CreationDate,
			LastChangeDate:             data.LastChangeDate,
			SalesOrganization:          data.SalesOrganization,
			DistributionChannel:        data.DistributionChannel,
			Division:                   data.Division,
			BillingDocumentDate:        data.BillingDocumentDate,
			BillingDocumentIsCancelled: data.BillingDocumentIsCancelled,
			CancelledBillingDocument:   data.CancelledBillingDocument,
			IsExportDelivery:           data.IsExportDelivery,
			TotalNetAmount:             data.TotalNetAmount,
			TransactionCurrency:        data.TransactionCurrency,
			TaxAmount:                  data.TaxAmount,
			TotalGrossAmount:           data.TotalGrossAmount,
			CustomerPriceGroup:         data.CustomerPriceGroup,
			IncotermsClassification:    data.IncotermsClassification,
			CustomerPaymentTerms:       data.CustomerPaymentTerms,
			PaymentMethod:              data.PaymentMethod,
			CompanyCode:                data.CompanyCode,
			AccountingDocument:         data.AccountingDocument,
			ExchangeRateDate:           data.ExchangeRateDate,
			ExchangeRateType:           data.ExchangeRateType,
			DocumentReferenceID:        data.DocumentReferenceID,
			SoldToParty:                data.SoldToParty,
			PartnerCompany:             data.PartnerCompany,
			PurchaseOrderByCustomer:    data.PurchaseOrderByCustomer,
			CustomerGroup:              data.CustomerGroup,
			Country:                    data.Country,
			CityCode:                   data.CityCode,
			Region:                     data.Region,
			CreditControlArea:          data.CreditControlArea,
			OverallBillingStatus:       data.OverallBillingStatus,
			AccountingPostingStatus:    data.AccountingPostingStatus,
			AccountingTransferStatus:   data.AccountingTransferStatus,
			InvoiceListStatus:          data.InvoiceListStatus,
			BillingDocumentListType:    data.BillingDocumentListType,
			BillingDocumentListDate:    data.BillingDocumentListDate,
			ToHeaderPartner:            data.ToHeaderPartner.Deferred.URI,
			ToItem:                     data.ToItem.Deferred.URI,
		})
	}

	return header, nil
}

func ConvertToHeaderPartner(raw []byte, l *logger.Logger) ([]HeaderPartner, error) {
	pm := &responses.HeaderPartner{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to HeaderPartner. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	headerPartner := make([]HeaderPartner, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		headerPartner = append(headerPartner, HeaderPartner{
			BillingDocument: data.BillingDocument,
			PartnerFunction: data.PartnerFunction,
			Customer:        data.Customer,
			Supplier:        data.Supplier,
		})
	}

	return headerPartner, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) ([]Item, error) {
	pm := &responses.Item{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	item := make([]Item, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		item = append(item, Item{
			BillingDocument:              data.BillingDocument,
			BillingDocumentItem:          data.BillingDocumentItem,
			SalesDocumentItemCategory:    data.SalesDocumentItemCategory,
			ReturnItemProcessingType:     data.ReturnItemProcessingType,
			CreationDate:                 data.CreationDate,
			Division:                     data.Division,
			Material:                     data.Material,
			InternationalArticleNumber:   data.InternationalArticleNumber,
			Batch:                        data.Batch,
			MaterialGroup:                data.MaterialGroup,
			Plant:                        data.Plant,
			StorageLocation:              data.StorageLocation,
			BillingDocumentItemText:      data.BillingDocumentItemText,
			BillingQuantity:              data.BillingQuantity,
			BillingQuantityUnit:          data.BillingQuantityUnit,
			NetAmount:                    data.NetAmount,
			TransactionCurrency:          data.TransactionCurrency,
			GrossAmount:                  data.GrossAmount,
			PricingDate:                  data.PricingDate,
			TaxAmount:                    data.TaxAmount,
			MaterialPricingGroup:         data.MaterialPricingGroup,
			MainItemMaterialPricingGroup: data.MainItemMaterialPricingGroup,
			BusinessArea:                 data.BusinessArea,
			ProfitCenter:                 data.ProfitCenter,
			WBSElement:                   data.WBSElement,
			ControllingArea:              data.ControllingArea,
			ProfitabilitySegment:         data.ProfitabilitySegment,
			OrderID:                      data.OrderID,
			CostCenter:                   data.CostCenter,
			ReferenceSDDocument:          data.ReferenceSDDocument,
			ReferenceSDDocumentItem:      data.ReferenceSDDocumentItem,
			MatlAccountAssignmentGroup:   data.MatlAccountAssignmentGroup,
			SalesDocument:                data.SalesDocument,
			SalesDocumentItem:            data.SalesDocumentItem,
			SDDocumentReason:             data.SDDocumentReason,
			ShippingPoint:                data.ShippingPoint,
			ToItemPartner:                data.ToItemPartner.Deferred.URI,
			ToItemPricingElement:         data.ToItemPricingElement.Deferred.URI,
		})
	}

	return item, nil
}

func ConvertToItemPartner(raw []byte, l *logger.Logger) ([]ItemPartner, error) {
	pm := &responses.ItemPartner{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ItemPartner. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	itemPartner := make([]ItemPartner, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		itemPartner = append(itemPartner, ItemPartner{
			BillingDocument:     data.BillingDocument,
			BillingDocumentItem: data.BillingDocumentItem,
			PartnerFunction:     data.PartnerFunction,
			Customer:            data.Customer,
			Supplier:            data.Supplier,
		})
	}

	return itemPartner, nil
}

func ConvertToToHeaderPartner(raw []byte, l *logger.Logger) ([]ToHeaderPartner, error) {
	pm := &responses.ToHeaderPartner{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToHeaderPartner. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toHeaderPartner := make([]ToHeaderPartner, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toHeaderPartner = append(toHeaderPartner, ToHeaderPartner{
			BillingDocument: data.BillingDocument,
			PartnerFunction: data.PartnerFunction,
			Customer:        data.Customer,
			Supplier:        data.Supplier,
		})
	}

	return toHeaderPartner, nil
}

func ConvertToToItem(raw []byte, l *logger.Logger) ([]ToItem, error) {
	pm := &responses.ToItem{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toItem := make([]ToItem, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItem = append(toItem, ToItem{
			BillingDocument:              data.BillingDocument,
			BillingDocumentItem:          data.BillingDocumentItem,
			SalesDocumentItemCategory:    data.SalesDocumentItemCategory,
			ReturnItemProcessingType:     data.ReturnItemProcessingType,
			CreationDate:                 data.CreationDate,
			Division:                     data.Division,
			Material:                     data.Material,
			InternationalArticleNumber:   data.InternationalArticleNumber,
			Batch:                        data.Batch,
			MaterialGroup:                data.MaterialGroup,
			Plant:                        data.Plant,
			StorageLocation:              data.StorageLocation,
			BillingDocumentItemText:      data.BillingDocumentItemText,
			BillingQuantity:              data.BillingQuantity,
			BillingQuantityUnit:          data.BillingQuantityUnit,
			NetAmount:                    data.NetAmount,
			TransactionCurrency:          data.TransactionCurrency,
			GrossAmount:                  data.GrossAmount,
			PricingDate:                  data.PricingDate,
			TaxAmount:                    data.TaxAmount,
			MaterialPricingGroup:         data.MaterialPricingGroup,
			MainItemMaterialPricingGroup: data.MainItemMaterialPricingGroup,
			BusinessArea:                 data.BusinessArea,
			ProfitCenter:                 data.ProfitCenter,
			WBSElement:                   data.WBSElement,
			ControllingArea:              data.ControllingArea,
			ProfitabilitySegment:         data.ProfitabilitySegment,
			OrderID:                      data.OrderID,
			CostCenter:                   data.CostCenter,
			ReferenceSDDocument:          data.ReferenceSDDocument,
			ReferenceSDDocumentItem:      data.ReferenceSDDocumentItem,
			MatlAccountAssignmentGroup:   data.MatlAccountAssignmentGroup,
			SalesDocument:                data.SalesDocument,
			SalesDocumentItem:            data.SalesDocumentItem,
			SDDocumentReason:             data.SDDocumentReason,
			ShippingPoint:                data.ShippingPoint,
			ToItemPartner:                data.ToItemPartner.Deferred.URI,
			ToItemPricingElement:         data.ToItemPricingElement.Deferred.URI,
		})
	}

	return toItem, nil
}

func ConvertToToItemPartner(raw []byte, l *logger.Logger) ([]ToItemPartner, error) {
	pm := &responses.ToItemPartner{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItemPartner. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toItemPartner := make([]ToItemPartner, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItemPartner = append(toItemPartner, ToItemPartner{
			BillingDocument:     data.BillingDocument,
			BillingDocumentItem: data.BillingDocumentItem,
			PartnerFunction:     data.PartnerFunction,
			Customer:            data.Customer,
			Supplier:            data.Supplier,
		})
	}

	return toItemPartner, nil
}

func ConvertToToItemPricingElement(raw []byte, l *logger.Logger) ([]ToItemPricingElement, error) {
	pm := &responses.ToItemPricingElement{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToItemPricingElement. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	toItemPricingElement := make([]ToItemPricingElement, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toItemPricingElement = append(toItemPricingElement, ToItemPricingElement{
			BillingDocument:             data.BillingDocument,
			BillingDocumentItem:         data.BillingDocumentItem,
			PricingProcedureStep:        data.PricingProcedureStep,
			PricingProcedureCounter:     data.PricingProcedureCounter,
			ConditionType:               data.ConditionType,
			PricingDateTime:             data.PricingDateTime,
			ConditionCalculationType:    data.ConditionCalculationType,
			ConditionBaseValue:          data.ConditionBaseValue,
			ConditionRateValue:          data.ConditionRateValue,
			ConditionCurrency:           data.ConditionCurrency,
			ConditionQuantity:           data.ConditionQuantity,
			ConditionQuantityUnit:       data.ConditionQuantityUnit,
			ConditionCategory:           data.ConditionCategory,
			PricingScaleType:            data.PricingScaleType,
			ConditionRecord:             data.ConditionRecord,
			ConditionSequentialNumber:   data.ConditionSequentialNumber,
			TaxCode:                     data.TaxCode,
			ConditionAmount:             data.ConditionAmount,
			TransactionCurrency:         data.TransactionCurrency,
			PricingScaleBasis:           data.PricingScaleBasis,
			ConditionScaleBasisValue:    data.ConditionScaleBasisValue,
			ConditionScaleBasisUnit:     data.ConditionScaleBasisUnit,
			ConditionScaleBasisCurrency: data.ConditionScaleBasisCurrency,
			ConditionIsManuallyChanged:  data.ConditionIsManuallyChanged,
		})
	}

	return toItemPricingElement, nil
}
