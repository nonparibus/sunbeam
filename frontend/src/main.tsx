import { createRoot } from "react-dom/client";
import { RaycastCMDK } from "./launcher/raycast";
import "./globals.scss";
import runtime from "../wailsjs/runtime";

const container = document.getElementById("app")
const root = createRoot(container!)
root.render(<RaycastCMDK/>)

// Minimise the window when unfocused
// window.onblur = () => {
//   WindowMinimise();
// };
