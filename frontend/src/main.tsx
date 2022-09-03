import { createRoot } from "react-dom/client";
import { RaycastCMDK } from "./launcher/raycast";
import "./globals.scss";

const container = document.getElementById("app")
const root = createRoot(container!)
root.render(<RaycastCMDK />)
