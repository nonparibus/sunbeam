
export interface UpdateResponse {
    Update: SearchResult[]
}

export function isUpdateResponse(response: any): response is UpdateResponse {
    return Object.keys(response).includes("Update")
}

export interface ContextResponse {
    Context: Context
}

export function isContextResponse(response: any): response is UpdateResponse {
    return Object.keys(response).includes("Context")
}

export interface SearchResult {
    id: number
    name: string
    description: string
    icon?: IconSource
    category_icon?: IconSource
    window: [number, number]
}

export interface DesktopEntryResponse {
    DesktopEntry: {
    path: string    
    gpu_preference: string
    }
}

export function isDesktopEntryResponse(response: any): response is DesktopEntryResponse {
    return Object.keys(response).includes("DesktopEntry")
}

export interface FillResponse {
    Fill: string
}

export function isFillResponse(response: any): response is FillResponse {
    return Object.keys(response).includes("Fill")
}

type CloseResponse = "Close"
export function isCloseResponse(response: any): response is CloseResponse {
    return response === "Close"
}

export interface Context {
    id: number
    name: string
}

export interface IconSource {
    Name?: string,
    Mime?: string
}

export interface MimeIcon {
    Mime: string
}

export interface NamedIcon {
    Name: string
}

export function isMimeIcon(icon: IconSource): icon is MimeIcon {
    return Object.keys(icon).includes("Mime")
}

export function isNamedIcon(icon: IconSource): icon is NamedIcon {
    return Object.keys(icon).includes("Name")
}

export type Response = UpdateResponse | ContextResponse | DesktopEntryResponse | FillResponse | CloseResponse
