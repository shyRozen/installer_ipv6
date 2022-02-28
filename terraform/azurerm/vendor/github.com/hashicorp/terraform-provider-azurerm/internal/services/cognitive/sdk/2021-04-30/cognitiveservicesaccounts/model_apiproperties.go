package cognitiveservicesaccounts

type ApiProperties struct {
	AadClientId                    *string `json:"aadClientId,omitempty"`
	AadTenantId                    *string `json:"aadTenantId,omitempty"`
	EventHubConnectionString       *string `json:"eventHubConnectionString,omitempty"`
	QnaAzureSearchEndpointId       *string `json:"qnaAzureSearchEndpointId,omitempty"`
	QnaAzureSearchEndpointKey      *string `json:"qnaAzureSearchEndpointKey,omitempty"`
	QnaRuntimeEndpoint             *string `json:"qnaRuntimeEndpoint,omitempty"`
	StatisticsEnabled              *bool   `json:"statisticsEnabled,omitempty"`
	StorageAccountConnectionString *string `json:"storageAccountConnectionString,omitempty"`
	SuperUser                      *string `json:"superUser,omitempty"`
	WebsiteName                    *string `json:"websiteName,omitempty"`
}
