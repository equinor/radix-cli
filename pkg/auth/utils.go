package auth

import (
	"fmt"
	"strings"
)

func accessTokenCacheKey(scopes []string) string {
	return fmt.Sprintf("access_token-%s", strings.Join(scopes, " "))
}
