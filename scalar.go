package scalar

import (
	"fmt"
	"net/http"
	"os"

	s "github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

func initOptions(options *Options) *s.Options {
	op := &Options{}

	if options != nil {
		op = options
	}

	if op.Layout == "" {
		op.Layout = LayoutModern
	}

	if op.Theme == "" {
		op.Theme = ThemeDefault
	}

	if op.SpecFile != "" {
		data, err := os.ReadFile(op.SpecFile)
		if err == nil {
			op.SpecContent = string(data)
		}
	}

	sop := &s.Options{
		SpecURL:            op.SpecURL,
		SpecContent:        options.SpecContent,
		Theme:              s.ThemeId(op.Theme),
		Layout:             s.ReferenceLayoutType(options.Layout),
		CDN:                op.CDN,
		Proxy:              op.Proxy,
		IsEditable:         op.IsEditable,
		ShowSidebar:        op.ShowSidebar,
		HideModels:         op.HideModels,
		HideDownloadButton: op.HideDownloadButton,
		DarkMode:           op.DarkMode,
		SearchHotKey:       op.SearchHotKey,
		MetaData:           op.MetaData,
		HiddenClients:      op.HiddenClients,
		CustomCss:          op.CustomCss,
		Authentication:     op.Authentication,
		PathRouting:        op.PageTitle,
		BaseServerURL:      op.BaseServerURL,
		WithDefaultFonts:   op.WithDefaultFonts,
	}

	if op.PageTitle == "" {
		sop.CustomOptions = s.CustomOptions{
			PageTitle: op.PageTitle,
		}
	}

	return sop
}

// Handler returns an http.HandlerFunc that can be used with both
// the Fiber V2
func FiberHandler(options *Options) fiber.Handler {
	sop := initOptions(options)
	return adaptor.HTTPHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		htmlContent, err := s.ApiReferenceHTML(sop)

		if err != nil {
			fmt.Printf("%v", err)
		}

		fmt.Fprintln(w, htmlContent)
	}))
}

// Handler returns an http.HandlerFunc that can be used with both
// the Fiber v3 framework and the standard Go http package.
func Handler(options *Options) func(w http.ResponseWriter, r *http.Request) {
	sop := initOptions(options)
	return func(w http.ResponseWriter, r *http.Request) {
		htmlContent, err := s.ApiReferenceHTML(sop)

		if err != nil {
			fmt.Printf("%v", err)
		}

		fmt.Fprintln(w, htmlContent)
	}
}
