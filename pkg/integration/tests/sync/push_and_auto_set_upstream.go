package sync

import (
	"github.com/lobes/lazytask/pkg/config"
	. "github.com/lobes/lazytask/pkg/integration/components"
)

var PushAndAutoSetUpstream = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Push a commit and set the upstream automatically as configured by git",
	ExtraCmdArgs: []string{},
	Skip:         false,
	SetupConfig: func(config *config.AppConfig) {
	},
	SetupRepo: func(shell *Shell) {
		shell.EmptyCommit("one")

		shell.CloneIntoRemote("origin")

		shell.EmptyCommit("two")

		shell.SetConfig("push.default", "current")
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		// assert no mention of upstream/downstream changes
		t.Views().Vitals().Content(MatchesRegexp(`^\s+repo → master`))

		t.Views().Files().
			IsFocused().
			Press(keys.Universal.Push)

		assertSuccessfullyPushed(t)
	},
})
