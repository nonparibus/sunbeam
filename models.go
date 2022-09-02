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
