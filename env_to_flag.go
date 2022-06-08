package flagx

import (
	"flag"
	"os"
	"strings"
)

// EnvToFlag is func for set to flag value from enviroments
var EnvToFlag = EnvToFlagWithPrefix("")

// EnvToFlagWithPrefix return the func for set to flag value from enviroments value with prefix
func EnvToFlagWithPrefix(prefix string) func(f *flag.Flag) {
	return func(f *flag.Flag) {
		names := []string{
			strings.ToUpper(prefix + strings.ReplaceAll(f.Name, "-", "_")),
			strings.ToLower(prefix + strings.ReplaceAll(f.Name, "-", "_")),
		}
		for _, name := range names {
			if s, ok := os.LookupEnv(name); ok {
				f.Value.Set(s)
				break
			}
		}
	}
}
