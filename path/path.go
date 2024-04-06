package path

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// ApplicationDir returns best base directory for specific OS.
func ApplicationDir(subDir ...string) string {
	cas := cases.Title(language.Dutch, cases.NoLower)
	for i := range subDir {
		if runtime.GOOS == "windows" || runtime.GOOS == "darwin" {
			//lint:ignore SA1019 we cannot change the behavior and it doesn't make a big difference.
			//nolint:staticcheck // see above
			subDir[i] = cas.String(subDir[i])
		} else {
			subDir[i] = strings.ToLower(subDir[i])
		}
	}
	var appDir string
	home := os.Getenv("HOME")
	switch runtime.GOOS {
	case "windows":
		// Windows standards: https://msdn.microsoft.com/en-us/library/windows/apps/hh465094.aspx?f=255&MSPPError=-2147217396
		for _, env := range []string{"AppData", "AppDataLocal", "UserProfile", "Home"} {
			val := os.Getenv(env)
			if val != "" {
				appDir = val
				break
			}
		}
	case "darwin":
		// Mac standards: https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/FileSystemProgrammingGuide/MacOSXDirectories/MacOSXDirectories.html
		appDir = filepath.Join(home, "Library", "Application Support")
	case "linux":
		fallthrough
	default:
		// Linux standards: https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html
		appDir = os.Getenv("XDG_DATA_HOME")
		if appDir == "" && home != "" {
			appDir = filepath.Join(home, ".local", "share")
		}
	}
	return filepath.Join(append([]string{appDir}, subDir...)...)
}
