package filetree

import (
	"sync"

	"github.com/lobes/lazytask/pkg/commands/models"
	"github.com/lobes/lazytask/pkg/gui/context/traits"
	"github.com/lobes/lazytask/pkg/gui/types"
	"github.com/lobes/lazytask/pkg/utils"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
)

type ITaskTreeViewModel interface {
	ITaskTree
	types.IListCursor
}

// This combines our TaskTree struct with a cursor that retains information about
// which item is selected. It also contains logic for repositioning that cursor
// after the files are refreshed
type TaskTreeViewModel struct {
	sync.RWMutex
	types.IListCursor
	ITaskTree
}

var _ ITaskTreeViewModel = &TaskTreeViewModel{}

func NewTaskTreeViewModel(getTasks func() []*models.Task, log *logrus.Entry, showTree bool) *TaskTreeViewModel {
	fileTree := NewTaskTree(getTasks, log, showTree)
	listCursor := traits.NewListCursor(fileTree.Len)
	return &TaskTreeViewModel{
		ITaskTree:   fileTree,
		IListCursor: listCursor,
	}
}

func (self *TaskTreeViewModel) GetSelected() *TaskNode {
	if self.Len() == 0 {
		return nil
	}

	return self.Get(self.GetSelectedLineIdx())
}

func (self *TaskTreeViewModel) GetSelectedItemId() string {
	item := self.GetSelected()
	if item == nil {
		return ""
	}

	return item.ID()
}

func (self *TaskTreeViewModel) GetSelectedItems() ([]*TaskNode, int, int) {
	if self.Len() == 0 {
		return nil, 0, 0
	}

	startIdx, endIdx := self.GetSelectionRange()

	nodes := []*TaskNode{}
	for i := startIdx; i <= endIdx; i++ {
		nodes = append(nodes, self.Get(i))
	}

	return nodes, startIdx, endIdx
}

func (self *TaskTreeViewModel) GetSelectedItemIds() ([]string, int, int) {
	selectedItems, startIdx, endIdx := self.GetSelectedItems()

	ids := lo.Map(selectedItems, func(item *TaskNode, _ int) string {
		return item.ID()
	})

	return ids, startIdx, endIdx
}

func (self *TaskTreeViewModel) GetSelectedTask() *models.Task {
	node := self.GetSelected()
	if node == nil {
		return nil
	}

	return node.Task
}

func (self *TaskTreeViewModel) GetSelectedPath() string {
	node := self.GetSelected()
	if node == nil {
		return ""
	}

	return node.GetPath()
}

func (self *TaskTreeViewModel) SetTree() {
	newTasks := self.GetAllTasks()
	selectedNode := self.GetSelected()

	// for when you stage the old file of a rename and the new file is in a collapsed dir
	for _, file := range newTasks {
		if selectedNode != nil && selectedNode.Path != "" && file.PreviousName == selectedNode.Path {
			self.ExpandToPath(file.Name)
		}
	}

	prevNodes := self.GetAllItems()
	prevSelectedLineIdx := self.GetSelectedLineIdx()

	self.ITaskTree.SetTree()

	if selectedNode != nil {
		newNodes := self.GetAllItems()
		newIdx := self.findNewSelectedIdx(prevNodes[prevSelectedLineIdx:], newNodes)
		if newIdx != -1 && newIdx != prevSelectedLineIdx {
			self.SetSelection(newIdx)
		}
	}

	self.ClampSelection()
}

// Let's try to find our file again and move the cursor to that.
// If we can't find our file, it was probably just removed by the user. In that
// case, we go looking for where the next file has been moved to. Given that the
// user could have removed a whole directory, we continue iterating through the old
// nodes until we find one that exists in the new set of nodes, then move the cursor
// to that.
// prevNodes starts from our previously selected node because we don't need to consider anything above that
func (self *TaskTreeViewModel) findNewSelectedIdx(prevNodes []*TaskNode, currNodes []*TaskNode) int {
	getPaths := func(node *TaskNode) []string {
		if node == nil {
			return nil
		}
		if node.Task != nil && node.Task.IsRename() {
			return node.Task.Names()
		} else {
			return []string{node.Path}
		}
	}

	for _, prevNode := range prevNodes {
		selectedPaths := getPaths(prevNode)

		for idx, node := range currNodes {
			paths := getPaths(node)

			// If you started off with a rename selected, and now it's broken in two, we want you to jump to the new file, not the old file.
			// This is because the new should be in the same position as the rename was meaning less cursor jumping
			foundOldTaskInRename := prevNode.Task != nil && prevNode.Task.IsRename() && node.Path == prevNode.Task.PreviousName
			foundNode := utils.StringArraysOverlap(paths, selectedPaths) && !foundOldTaskInRename
			if foundNode {
				return idx
			}
		}
	}

	return -1
}

func (self *TaskTreeViewModel) SetStatusFilter(filter TaskTreeDisplayFilter) {
	self.ITaskTree.SetStatusFilter(filter)
	self.IListCursor.SetSelection(0)
}

// If we're going from flat to tree we want to select the same file.
// If we're going from tree to flat and we have a file selected we want to select that.
// If instead we've selected a directory we need to select the first file in that directory.
func (self *TaskTreeViewModel) ToggleShowTree() {
	selectedNode := self.GetSelected()

	self.ITaskTree.ToggleShowTree()

	if selectedNode == nil {
		return
	}
	path := selectedNode.Path

	if self.InTreeMode() {
		self.ExpandToPath(path)
	} else if len(selectedNode.Children) > 0 {
		path = selectedNode.GetLeaves()[0].Path
	}

	index, found := self.GetIndexForPath(path)
	if found {
		self.SetSelectedLineIdx(index)
	}
}
