package sync

import (
	"github.com/lobes/lazytask/pkg/config"
	. "github.com/lobes/lazytask/pkg/integration/components"
)

var Pull = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Pull a commit from the remote",
	ExtraCmdArgs: []string{},
	Skip:         false,
	SetupConfig:  func(config *config.AppConfig) {},
	SetupRepo: func(shell *Shell) {
		shell.EmptyCommit("one")
		shell.EmptyCommit("two")

		shell.CloneIntoRemote("origin")
		shell.SetBranchUpstream("master", "origin/master")

		// remove the 'two' commit so that we have something to pull from the remote
		shell.HardReset("HEAD^")
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Commits().
			Lines(
				Contains("one"),
			)

		t.Views().Vitals().Content(Contains("↓1 repo → master"))

		t.Views().Files().IsFocused().Press(keys.Universal.Pull)

		t.Views().Commits().
			Lines(
				Contains("two"),
				Contains("one"),
			)

		t.Views().Vitals().Content(Contains("✓ repo → master"))
	},
})
