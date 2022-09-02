
export interface UpdateResponse {
    Update: SearchResult[]
}

export function isUpdateResponse(response: any): response is UpdateResponse {
    return Object.keys(response).includes("Update")
}

export interface ContextResponse {
    Context: Context
}

export interface SearchResult {
    id: number
    name: string
    description: string
    icon: IconSource
    category_icon: IconSource
    window: [number, number]
}

export interface DesktopEntryResponse {
    path: string    
    gpu_preference: string
}

export interface FillResponse {
    Fill: string
}

type CloseResponse = "Close"

interface Context {
    id: number
    name: string
}

interface IconSource {
    Name: string,
    Mime: string
}


export type Response = UpdateResponse | ContextResponse | DesktopEntryResponse | FillResponse | CloseResponse
