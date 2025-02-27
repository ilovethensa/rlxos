package ignite

import (
	"fmt"
	"path"
	"rlxos/internal/element"
	"rlxos/internal/updater"
	"strings"
)

func (b *Ignite) Update(e *element.Element) (string, error) {
	if len(e.Sources) == 0 || !strings.HasPrefix(e.Sources[0], "http") {
		return "", nil
	}

	u, _ := updater.New(path.Join(b.projectPath, "updates.yml"))

	version, err := u.GetUpdate(e)
	if err != nil {
		return "", fmt.Errorf("failed to get updates for %s, %v", e.Id, err)
	}
	if version == e.Version {
		return "", nil
	}
	return version, nil
}
