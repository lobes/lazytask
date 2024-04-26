package presentation

import (
	// "strings"

	// "github.com/gookit/color"
	"github.com/lobes/lazytask/pkg/commands/models"
	// "github.com/lobes/lazytask/pkg/commands/patch"
	"github.com/lobes/lazytask/pkg/gui/filetree"
	// "github.com/lobes/lazytask/pkg/gui/presentation/icons"
	// "github.com/lobes/lazytask/pkg/gui/style"
	// "github.com/lobes/lazytask/pkg/theme"
	// "github.com/lobes/lazytask/pkg/utils"
)

// const (
// 	EXPANDED_ARROW  = "▼"
// 	COLLAPSED_ARROW = "►"
// 	TASK_OPEN       = "☐"
// 	TASK_CLOSED     = "☒"
// )

func RenderTaskTree(
	tree filetree.IFileTree,
	submoduleConfigs []*models.SubmoduleConfig,
	showFileIcons bool,
) []string {
	// read all task toml
	// make a []string of the descriptions
	return models.ReadTask()
}

// TODO repurpose for task status
// func getColorForChangeStatus(changeStatus string) style.TextStyle {
// 	switch changeStatus {
// 	case "A":
// 		return style.FgGreen
// 	case "M", "R":
// 		return style.FgYellow
// 	case "D":
// 		return theme.UnstagedChangesColor
// 	case "C":
// 		return style.FgCyan
// 	case "T":
// 		return style.FgMagenta
// 	default:
// 		return theme.DefaultTextColor
// 	}
// }

// func split(str string) []string {
// 	return strings.Split(str, "/")
// }

// func join(strs []string) string {
// 	return strings.Join(strs, "/")
// }
