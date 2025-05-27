package scalar

type ThemeId string

const (
	ThemeDefault    ThemeId = "default"
	ThemeAlternate  ThemeId = "alternate"
	ThemeMoon       ThemeId = "moon"
	ThemePurple     ThemeId = "purple"
	ThemeSolarized  ThemeId = "solarized"
	ThemeBluePlanet ThemeId = "bluePlanet"
	ThemeDeepSpace  ThemeId = "deepSpace"
	ThemeSaturn     ThemeId = "saturn"
	ThemeKepler     ThemeId = "kepler"
	ThemeMars       ThemeId = "mars"
	ThemeNone       ThemeId = "none"
	ThemeNil        ThemeId = ""
)

type Layout string

const (
	LayoutClassic Layout = "classic"
	LayoutModern  Layout = "modern"
)

type Options struct {
	CDN                string      `json:"cdn,omitempty"`
	PageTitle          string      `json:"pageTitle"`
	SpecFile           string      `json:"specFile,omitempty"`
	SpecURL            string      `json:"specUrl,omitempty"` // allow external URL ou local path file
	SpecContent        interface{} `json:"specContent,omitempty"`
	Layout             Layout      `json:"layout"`
	DarkMode           bool        `json:"darkMode"`
	Theme              ThemeId     `json:"theme,omitempty"`
	IsEditable         bool        `json:"isEditable,omitempty"`
	SearchHotKey       string      `json:"searchHotKey,omitempty"`
	MetaData           string      `json:"metaData,omitempty"`
	HiddenClients      []string    `json:"hiddenClients,omitempty"`
	HideModels         bool        `json:"hideModels,omitempty"`
	PathRouting        string      `json:"pathRouting,omitempty"`
	Authentication     string      `json:"authentication,omitempty"`
	CustomCss          string      `json:"customCss,omitempty"`
	BaseServerURL      string      `json:"baseServerUrl,omitempty"`
	WithDefaultFonts   bool        `json:"withDefaultFonts,omitempty"`
	HideDownloadButton bool        `json:"hideDownloadButton,omitempty"`
	Proxy              string      `json:"proxy,omitempty"`
	ShowSidebar        bool        `json:"showSidebar,omitempty"`
}
