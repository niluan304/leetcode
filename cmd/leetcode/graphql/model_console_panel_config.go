package graphql

type (
	ConsolePanelConfigRes struct {
		Question ConsolePanelConfigQuestion `json:"question"`
	}
	ConsolePanelConfigQuestion struct {
		QuestionId           string `json:"questionId"`
		QuestionFrontendId   string `json:"questionFrontendId"`
		QuestionTitle        string `json:"questionTitle"`
		EnableRunCode        bool   `json:"enableRunCode"`
		EnableSubmit         bool   `json:"enableSubmit"`
		EnableTestMode       bool   `json:"enableTestMode"`
		JsonExampleTestcases string `json:"jsonExampleTestcases"`
		ExampleTestcases     string `json:"exampleTestcases"`
		MetaData             string `json:"metaData"`
		SampleTestCase       string `json:"sampleTestCase"`
	}
)
