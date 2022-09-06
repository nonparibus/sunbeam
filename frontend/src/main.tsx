import { createRoot } from "react-dom/client";
import { Raycast } from "./launcher/raycast";
import "./globals.scss";

const container = document.getElementById("app")
const root = createRoot(container!)
root.render(< Raycast/>)
