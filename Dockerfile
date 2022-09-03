FROM gitpod/workspace-full-vnc

RUN sudo apt-get update \
    && sudo apt-get -y install libgtk-3-dev libwebkit2gtk-4.0-dev

# Install pop launcher system wide
RUN git clone https://github.com/pop-os/launcher \
    && cd launcher \
    && brew install just \
    && just \
    && sudo $(which just) rootDir=/ install

RUN go install github.com/wailsapp/wails/v2/cmd/wails@latest

ENV DESKTOP_SESSION xfce
