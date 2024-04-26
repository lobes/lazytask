package context

import (
	"github.com/lobes/lazytask/pkg/commands/models"
	"github.com/lobes/lazytask/pkg/gui/filetree"
	"github.com/lobes/lazytask/pkg/gui/presentation"
	"github.com/lobes/lazytask/pkg/gui/presentation/icons"
	"github.com/lobes/lazytask/pkg/gui/types"
	"github.com/samber/lo"
)

type TaskTreeContext struct {
	*tasktree.FileTreeViewModel
	*ListContextTrait
	*SearchTrait
}

var _ types.IListContext = (*TaskTreeContext)(nil)

func NewTaskTreeContext(c *ContextCommon) *TaskTreeContext {
	viewModel := tasktree.NewFileTreeViewModel(
		func() []*models.File { return c.Model().Files },
		c.Log,
		c.UserConfig.Gui.ShowFileTree,
	)

	getDisplayStrings := func(_ int, _ int) [][]string {
		showFileIcons := icons.IsIconEnabled() && c.UserConfig.Gui.ShowFileIcons
		lines := presentation.RenderFileTree(viewModel, c.Model().Submodules, showFileIcons)
		return lo.Map(lines, func(line string, _ int) []string {
			return []string{line}
		})
	}

	ctx := &TaskTreeContext{
		SearchTrait:       NewSearchTrait(c),
		FileTreeViewModel: viewModel,
		ListContextTrait: &ListContextTrait{
			Context: NewSimpleContext(NewBaseContext(NewBaseContextOpts{
				View:       c.Views().Files,
				WindowName: "files",
				Key:        FILES_CONTEXT_KEY,
				Kind:       types.SIDE_CONTEXT,
				Focusable:  true,
			})),
			ListRenderer: ListRenderer{
				list:              viewModel,
				getDisplayStrings: getDisplayStrings,
			},
			c: c,
		},
	}

	ctx.GetView().SetOnSelectItem(ctx.SearchTrait.onSelectItemWrapper(func(selectedLineIdx int) error {
		ctx.GetList().SetSelection(selectedLineIdx)
		return ctx.HandleFocus(types.OnFocusOpts{})
	}))

	return ctx
}
