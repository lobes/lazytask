package types

import (
	"github.com/jesseduffield/gocui"
	"github.com/lobes/lazytask/pkg/commands/models"
	"github.com/lobes/lazytask/pkg/config"
	"github.com/lobes/lazytask/pkg/gui/types"
)

// these interfaces are used by the gui package so that it knows what it needs
// to provide to a test in order for the test to run.

type IntegrationTest interface {
	Run(GuiDriver)
	SetupConfig(config *config.AppConfig)
	RequiresHeadless() bool
	// width and height when running headless
	HeadlessDimensions() (int, int)
	// If true, we are recording/replaying a demo
	IsDemo() bool
}

// this is the interface through which our integration tests interact with the lazygit gui
type GuiDriver interface {
	PressKey(string)
	Click(int, int)
	Keys() config.KeybindingConfig
	CurrentContext() types.Context
	ContextForView(viewName string) types.Context
	Fail(message string)
	// These two log methods are for the sake of debugging while testing. There's no need to actually
	// commit any logging.
	// logs to the normal place that you log to i.e. viewable with `lazygit --logs`
	Log(message string)
	// logs in the actual UI (in the commands panel)
	LogUI(message string)
	CheckedOutRef() *models.Branch
	// the view that appears to the right of the side panel
	MainView() *gocui.View
	// the other view that sometimes appears to the right of the side panel
	// e.g. when we're showing both staged and unstaged changes
	SecondaryView() *gocui.View
	View(viewName string) *gocui.View
	SetCaption(caption string)
	SetCaptionPrefix(prefix string)
	// Pop the next toast that was displayed; returns nil if there was none
	NextToast() *string
	CheckAllToastsAcknowledged()
}
