package builder

import (
	"strings"

	"github.com/ishaqcherry9/sqlx/pkg/frag"
)

// ForShare read locker
func ForShare(modifiers ...string) Addition {
	return AsAddition(
		addition_FOR_SHARE,
		frag.Lit(strings.Join(append([]string{"FOR SHARE"}, modifiers...), " ")),
	)
}

func ForShareSkipLocked() Addition {
	return ForShare("SKIP LOCKED")
}

func ForShareNoWait() Addition {
	return ForShare("NOWAIT")
}
