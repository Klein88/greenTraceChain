package Model

type PDFData struct {
	Reportid       string `json:"report_id"`
	Txid           string `json:"txid"`
	CompanyName    string `json:"company_name"`
	CompanyType    string `json:"company_type"`
	CarbonModel    string `json:"carbon_model"`
	TotalEmissions string `json:"total_emissions"`
	TimeStamp      string `json:"time_stamp"`
}
