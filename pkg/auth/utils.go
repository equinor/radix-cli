package auth

import (
	"fmt"
	"slices"
	"strings"
)

func accessTokenCacheKey(scopes []string) string {
	clonedScopes := slices.Clone(scopes)
	slices.Sort(clonedScopes)
	return fmt.Sprintf("access_token-%s", strings.Join(clonedScopes, " "))
}
