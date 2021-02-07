package utils

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// ParseEnv takes a prefix string and exposes environment variables
// for all flags in the default FlagSet (flag.CommandLine) in the
// form of PREFIX_FLAGNAME.
func ParseEnv(p string) {
	updateEnv(p, flag.CommandLine)
}

// updateEnv takes a prefix string p and *flag.FlagSet. Each flag
// in the FlagSet is exposed as an upper case environment variable
// prefixed with p. Any flag that was not explicitly set by a user
// is updated to the environment variable, if set.
func updateEnv(prefix string, fs *flag.FlagSet) {
	// Build a map of explicitly set flags.
	set := map[string]interface{}{}
	fs.Visit(func(f *flag.Flag) {
		set[f.Name] = nil
	})

	sep := "_"
	if prefix == "" {
		sep = ""
	}

	fs.VisitAll(func(f *flag.Flag) {
		// Create an env var name
		// based on the supplied prefix.
		envVar := fmt.Sprintf("%s%s%s", prefix, sep, f.Name)
		envVar = strings.Replace(envVar, "-", "_", -1)
		envVar = strings.ToUpper(envVar)

		// Update the Flag.Value if the
		// env var is non "".
		if val := os.Getenv(envVar); val != "" {
			// Update the value if it hasn't
			// already been set.
			if _, defined := set[f.Name]; !defined {
				fs.Set(f.Name, val)
			}
		}

		// Append the env var to the
		// Flag.Usage field.
		f.Usage = fmt.Sprintf("%s [%s]", f.Usage, envVar)
	})
}
