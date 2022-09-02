package main

type ActivateRequest struct {
	Activate int
}

type ActivateContextRequest struct {
	ActivateContext ActivateContext
}

type CompleteRequest struct {
	Complete int
}

type ContextRequest struct {
	Context int
}

type QuitRequest struct {
	Quit int
}

type SearchRequest struct {
	Search string
}

type ActivateContext struct {
	Id      int
	Context int
}

// type Response struct {
// 	Context      Context
// 	Update       []SearchResult
// 	DesktopEntry DesktopEntry
// 	Fill         string
// }

// type Context struct {
// 	Id   int    `json:"id"`
// 	Name string `json:"name"`
// }

// type DesktopEntry struct {
// 	Path          string `json:"path"`
// 	GpuPreference string `json:"gpu_preference"`
// }

// type SearchResult struct {
// 	Id           int        `json:"id"`
// 	Name         string     `json:"name"`
// 	Description  string     `json:"description"`
// 	Icon         IconSource `json:"icon"`
// 	CategoryIcon IconSource `json:"category_icon"`
// 	Window       []int      `json:"window"`
// }

// type IconSource struct {
// 	Name string
// 	Mime string
// }
