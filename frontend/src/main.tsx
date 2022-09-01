import { render } from "react-dom";
import { RaycastCMDK } from "./components/raycast";
import "./globals.scss";
import { WindowMinimise } from "../wailsjs/runtime";

render(<RaycastCMDK />, document.getElementById("app")!);

// Minimise the window when unfocused
window.onblur = () => {
  WindowMinimise();
};
