FROM gitpod/workspace-full-vnc

RUN sudo apt-get update \
    && sudo apt-get -y install libgtk-3-dev libwebkit2gtk-4.0-dev
RUN go install github.com/wailsapp/wails/v2/cmd/wails@latest
