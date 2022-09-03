package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/go-ini/ini"
)

func currentTheme() (string, error) {
	cmd := exec.Command("gsettings", "get", "org.gnome.desktop.interface", "gtk-theme")
	theme, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("An error occured while fetching the theme: %s", err)
	}

	return string(theme), nil
}

type IconFinder struct {
	iconMap map[string]string
}

func NewIconFinder() *IconFinder {
	return &IconFinder{
		iconMap: make(map[string]string),
	}
}

func (i *IconFinder) loadThemeIcons(themeDirectory string) error {
	config, err := ini.Load(path.Join(themeDirectory, "index.theme"))
	if err != nil {
		return NewConfigError(err)
	}
	mainSection, err := config.GetSection("Icon Theme")
	if err != nil {
		return NewConfigError(err)
	}
	directoriesKey, err := mainSection.GetKey("Directories")
	if err != nil {
		return NewConfigError(err)
	}
	directories := strings.Split(directoriesKey.Value(), ",")

	for _, directory := range directories {
		if directory == "" {
			continue
		}
		// directorySection, err := config.GetSection(directory)
		// if err != nil {
		// 	return NewConfigError(err)
		// }

		// ContextKey, err := directorySection.GetKey("Context")
		// if err != nil {
		// 	return NewConfigError(err)
		// }
		// if ContextKey.Value() != "Applications" {
		// 	continue
		// }

		// SizeKey, err := directorySection.GetKey("Size")
		// if err != nil {
		// 	return NewConfigError(err)
		// }

		directoryRelPath := strings.ReplaceAll(directory, "@2", "")
		directoryPath := path.Join(themeDirectory, directoryRelPath)
		icons, err := os.ReadDir(directoryPath)
		if err != nil {
			continue
		}

		for _, icon := range icons {
			basename := icon.Name()
			i.iconMap[basename] = path.Join(directoryPath, basename)
		}

	}

	return nil
}

type ConfigError struct {
	err error
}

func NewConfigError(err error) *ConfigError {
	return &ConfigError{err}
}

func (c *ConfigError) Error() string {
	return fmt.Sprintf("Unable to load config: %s", c.err)
}

func (i *IconFinder) getIconPath(iconName string, iconType string, acceptedExtensions []string) (string, bool) {
	if iconType == "mime" {
		iconName = strings.Replace(iconName, "/", "-", -1)
	}
	for _, extension := range acceptedExtensions {
		iconKey := fmt.Sprintf("%s.%s", iconName, extension)
		iconValue, ok := i.iconMap[iconKey]
		if ok {
			return iconValue, true
		}
	}
	return "", false
}

type FileLoader struct {
	http.Handler
	iconFinder *IconFinder
	extensions []string
}

func NewFileLoader(iconFinder *IconFinder) *FileLoader {
	return &FileLoader{
		iconFinder: iconFinder,
	}
}

var iconThemes = []string{
	"Adwaita",
	"hicolor",
}

func getContentType(icon string) (string, error) {
	if strings.HasSuffix(icon, ".png") {
		return "image/png", nil
	}
	if strings.HasSuffix(icon, ".svg") {
		return "image/svg+xml", nil
	}

	return "", fmt.Errorf("Could not find content type of icon: %s", icon)
}

func (h *FileLoader) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	iconName := strings.TrimPrefix(req.URL.Path, "/")
	iconType := req.URL.Query().Get("type")
	iconPath, ok := h.iconFinder.getIconPath(iconName, iconType, []string{"png", "svg"})

	if !ok {
		println("not found: %s", iconName)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf("Could not find icon %s", iconName)))
	}

	fileData, err := os.ReadFile(iconPath)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf("Could not load icon %s", iconName)))
	}

	contentType, err := getContentType(iconPath)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(fmt.Sprintf("Could not infer content type")))
	}

	res.Header().Set("content-type", contentType)
	res.Write(fileData)
}
