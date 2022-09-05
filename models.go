package main

type SearchItem struct {
	IconSource     string   `json:"icon_src"`
	Title          string   `json:"title" validate:"required"`
	Subtitle       string   `json:"subtitle"`
	Fill           string   `json:"fill"`
	AccessoryTitle string   `json:"accessory_title"`
	Keywords       []string `json:"keywords"`
	Actions        []Action `json:"actions" validate:"required,gte=1,dive"`
}

type Action struct {
	Icon  string `json:"icon"`
	Title string `json:"title" validate:"required"`
	Command
	Shortcut Shortcut `json:"shortcut"`
}

type Command struct {
	Type   string                 `json:"type" validate:"required"`
	Params map[string]interface{} `json:"params" validate:"required"`
}

type Shortcut struct {
	Ctrl  bool   `json:"ctrl"`
	Shift bool   `json:"shift"`
	Alt   bool   `json:"alt"`
	Super bool   `json:"super"`
	Key   string `json:"key" validate:"required"`
}

func (c Command) Icon() string {
	switch c.Type {
	case "open-file":
		return "/raycast/icon-blank-document-16.svg"
	case "open-url":
		return "/raycast/icon-globe-01-16.svg"
	case "copy-to-clipboard":
		return "/raycast/icon-copy-clipboard-16.svg"
	case "run-script":
		return "/raycast/icon-globe-01-16.svg"
	case "push-list":
		return "/raycast/app-window-list-16.svg"
	default:
		return ""
	}
}

func NewOpenCommand(filepath string) Command {
	return Command{Type: "open-file", Params: map[string]interface{}{
		"filepath": filepath,
	}}
}

func NewOpenInBrowserCommand(url string) Command {
	return Command{Type: "open-url", Params: map[string]interface{}{
		"url": url,
	}}
}

func NewCopyToClipboardCommand(content string) Command {
	return Command{Type: "copy-to-clipboard", Params: map[string]interface{}{
		"content": content,
	}}
}

func NewRunScriptCommand(scriptPath string, args []string, mode string) Command {
	return Command{Type: "run-script", Params: map[string]interface{}{
		"scriptpath": scriptPath,
		"args":       args,
	}}
}

func NewPushListCommand(scriptPath string, args []string, mode string) Command {
	return Command{Type: "push-list", Params: map[string]interface{}{
		"scriptpath": scriptPath, "mode": mode,
	}}
}
