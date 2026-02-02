package builder

import (
	"strings"

	"github.com/ishaqcherry9/sqlx/pkg/frag"
)

// ForUpdate write locker
func ForUpdate(modifiers ...string) Addition {
	return AsAddition(
		addition_FOR_UPDATE,
		frag.Lit(strings.Join(append([]string{"FOR UPDATE"}, modifiers...), " ")),
	)
}

func ForUpdateSkipLocked() Addition {
	return ForUpdate("SKIP LOCKED")
}

func ForUpdateNoWait() Addition {
	return ForUpdate("NOWAIT")
}
