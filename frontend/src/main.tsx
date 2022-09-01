import { createRoot } from "react-dom/client";
import { RaycastCMDK } from "./components/raycast";
import "./globals.scss";
import { WindowMinimise } from "../wailsjs/runtime";

const container = document.getElementById("app")
const root = createRoot(container!)
root.render(<RaycastCMDK/>)

// Minimise the window when unfocused
// window.onblur = () => {
//   WindowMinimise();
// };
