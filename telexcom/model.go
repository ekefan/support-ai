package telexcom

// Date structure
type Date struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Descriptions structure
type Descriptions struct {
	AppDescription  string `json:"app_description"`
	AppLogo         string `json:"app_logo"`
	AppName         string `json:"app_name"`
	AppURL          string `json:"app_url"`
	BackgroundColor string `json:"background_color"`
}

// MonitoringUser structure
type MonitoringUser struct {
	AlwaysOnline bool   `json:"always_online"`
	DisplayName  string `json:"display_name"`
}

// Permissions structure
type Permissions struct {
	MonitoringUser MonitoringUser `json:"monitoring_user"`
}

// Setting structure
type IntegrationSetting struct {
	Label    string `json:"label"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
	Default  string `json:"default"`
}

type Data struct {
	Date                Date                 `json:"date"`
	Descriptions        Descriptions         `json:"descriptions"`
	IntegrationCategory string               `json:"integration_category"`
	IntegrationType     string               `json:"integration_type"`
	IsActive            bool                 `json:"is_active"`
	KeyFeatures         []string             `json:"key_features"`
	Author              string               `json:"author"`
	Settings            []IntegrationSetting `json:"settings"`
	TickURL             string               `json:"tick_url"`
	TargetURL           string               `json:"target_url"`
}

// IntegrationJson main structure
type Integration struct {
	Data Data `json:"data"`
}

// TODO: edit integration.json accordingly
var IntegrationJson = Integration{
	Data: Data{
		Date: Date{
			CreatedAt: "2025-03-05",
			UpdatedAt: "2025-03-05",
		},
		Descriptions: Descriptions{
			AppDescription:  "The support assistant for your organisation. Reply FAQs, check documentation and thrive",
			AppLogo:         "https://raw.githubusercontent.com/samuelIkoli/Chatbot-AI-Agent/refs/heads/dev/public/home/chatbot.png",
			AppName:         "Telex Support AI",
			AppURL:          "https://telex.im",
			BackgroundColor: "#ffff",
		},
		IntegrationCategory: "AI & Machine Learning",
		IntegrationType:     "interval",
		IsActive:            true,
		KeyFeatures:         []string{"Gives apt responses to FAQs.", "Sends users response based on data stored in database."},
		Author:              "TSA team hng12",
		Settings: []IntegrationSetting{
			{
				Label:    "support-channel-id",
				Type:     "text",
				Required: true,
				Default:  "",
			},
		},
		TargetURL: "https://support-ai-uuyb.onrender.com/target",
		TickURL:   "",
	},
}

var NgrokIntegrationJson = Integration{
	Data: Data{
		Date: Date{
			CreatedAt: "2025-03-05",
			UpdatedAt: "2025-03-05",
		},
		Descriptions: Descriptions{
			AppDescription:  "The support assistant for your organisation. Reply FAQs, check documentation and thrive",
			AppLogo:         "https://raw.githubusercontent.com/samuelIkoli/Chatbot-AI-Agent/refs/heads/dev/public/home/chatbot.png",
			AppName:         "Ngrok Support AI",
			AppURL:          "https://telex.im",
			BackgroundColor: "#ffff",
		},
		IntegrationCategory: "AI & Machine Learning",
		IntegrationType:     "interval",
		IsActive:            true,
		KeyFeatures:         []string{"Gives apt responses to FAQs.", "Sends users response based on data stored in database."},
		Author:              "TSA team hng12",
		Settings: []IntegrationSetting{
			{
				Label:    "support-channel-id",
				Type:     "text",
				Required: true,
				Default:  "",
			},
		},
		TargetURL: "https://c24b-102-88-55-142.ngrok-free.app/target",
		TickURL:   "",
	},
}

type MonitorPayload struct {
	ChannelID string        `json:"channel_id,omitempty"`
	ReturnURL string        `json:"return_url,omitempty"`
	Settings  []interface{} `json:"settings,omitempty"`
}

//TODO: clear this from here
// type PromptPayload struct {
// 	PathURL  string `json:"path_url,omitempty"`
// 	Question string `json:"question,omitempty"`
// }

type TelexChatPayload struct {
	ChannelID string               `json:"channel_id,omitempty"`
	Message   string               `json:"message,omitempty"`
	Settings  []IntegrationSetting `json:"settings,omitempty"`
}

type TelexResponsePayload struct {
	Message   string `json:"message"`
	Username  string `json:"username"` // the name of your integration
	EventName string `json:"event_name"`
	Status    string `json:"status"`
}


