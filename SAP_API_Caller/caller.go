package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-billing-document-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetBillingDocument(billingDocument, partnerFunction, billingDocumentItem, itemPartnerFunction string, accepter []string) {
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
				c.HeaderPartner(billingDocument, partnerFunction)
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

func (c *SAPAPICaller) Header(billingDocument string) {
	headerData, err := c.callBillingDocumentSrvAPIRequirementHeader("A_BillingDocument", billingDocument)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(headerData)

	headerPartnerData, err := c.callToHeaderPartner(headerData[0].ToHeaderPartner)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(headerPartnerData)

	itemData, err := c.callToItem(headerData[0].ToItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemData)

	itemPartnerData, err := c.callToItemPartner(itemData[0].ToItemPartner)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemPartnerData)

	itemPricingElementData, err := c.callToItemPricingElement(itemData[0].ToItemPricingElement)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemPricingElementData)
}

func (c *SAPAPICaller) callBillingDocumentSrvAPIRequirementHeader(api, billingDocument string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_BILLING_DOCUMENT_SRV", api}, "/")
	param := c.getQueryWithHeader(map[string]string{}, billingDocument)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToHeaderPartner(url string) ([]sap_api_output_formatter.ToHeaderPartner, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToHeaderPartner(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItem(url string) ([]sap_api_output_formatter.ToItem, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItem(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemPartner(url string) ([]sap_api_output_formatter.ToItemPartner, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemPartner(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemPricingElement(url string) ([]sap_api_output_formatter.ToItemPricingElement, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemPricingElement(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) HeaderPartner(billingDocument, partnerFunction string) {
	data, err := c.callBillingDocumentSrvAPIRequirementHeaderPartner("A_BillingDocumentPartner", billingDocument, partnerFunction)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBillingDocumentSrvAPIRequirementHeaderPartner(api, billingDocument, partnerFunction string) ([]sap_api_output_formatter.HeaderPartner, error) {
	url := strings.Join([]string{c.baseURL, "API_BILLING_DOCUMENT_SRV", api}, "/")

	param := c.getQueryWithHeaderPartner(map[string]string{}, billingDocument, partnerFunction)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeaderPartner(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Item(billingDocument, billingDocumentItem string) {
	itemData, err := c.callBillingDocumentSrvAPIRequirementItem("A_BillingDocumentItem", billingDocument, billingDocumentItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemData)

	itemPartnerData, err := c.callToItemPartner(itemData[0].ToItemPartner)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemPartnerData)

	itemPricingElementData, err := c.callToItemPricingElement(itemData[0].ToItemPricingElement)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemPricingElementData)
}

func (c *SAPAPICaller) callBillingDocumentSrvAPIRequirementItem(api, billingDocument, billingDocumentItem string) ([]sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "API_BILLING_DOCUMENT_SRV", api}, "/")

	param := c.getQueryWithItem(map[string]string{}, billingDocument, billingDocumentItem)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) ItemPartner(billingDocument, billingDocumentItem, partnerFunction string) {
	data, err := c.callBillingDocumentSrvAPIRequirementItemPartner("A_BillingDocumentItemPartner", billingDocument, billingDocumentItem, partnerFunction)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callBillingDocumentSrvAPIRequirementItemPartner(api, billingDocument, billingDocumentItem, partnerFunction string) ([]sap_api_output_formatter.ItemPartner, error) {
	url := strings.Join([]string{c.baseURL, "API_BILLING_DOCUMENT_SRV", api}, "/")

	param := c.getQueryWithItemPartner(map[string]string{}, billingDocument, billingDocumentItem, partnerFunction)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItemPartner(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithHeader(params map[string]string, billingDocument string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("BillingDocument eq '%s'", billingDocument)
	return params
}

func (c *SAPAPICaller) getQueryWithHeaderPartner(params map[string]string, billingDocument, partnerFunction string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("BillingDocument eq '%s' and PartnerFunction eq '%s'", billingDocument, partnerFunction)
	return params
}

func (c *SAPAPICaller) getQueryWithItem(params map[string]string, billingDocument, billingDocumentItem string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("BillingDocument eq '%s' and BillingDocumentItem eq '%s'", billingDocument, billingDocumentItem)
	return params
}

func (c *SAPAPICaller) getQueryWithItemPartner(params map[string]string, billingDocument, billingDocumentItem, partnerFunction string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("BillingDocument eq '%s' and BillingDocumentItem eq '%s' and PartnerFunction eq '%s'", billingDocument, billingDocumentItem, partnerFunction)
	return params
}
