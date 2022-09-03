FROM gitpod/workspace-full-vnc

# Add just repository
RUN curl -q 'https://proget.makedeb.org/debian-feeds/prebuilt-mpr.pub' \
        | gpg --dearmor \
        | sudo tee /usr/share/keyrings/prebuilt-mpr-archive-keyring.gpg 1> /dev/null \
    && echo "deb [signed-by=/usr/share/keyrings/prebuilt-mpr-archive-keyring.gpg] https://proget.makedeb.org prebuilt-mpr $(lsb_release -cs)" \
        | sudo tee /etc/apt/sources.list.d/prebuilt-mpr.list

RUN sudo apt-get update \
    && sudo apt-get -y install \
        libgtk-3-dev \
        libwebkit2gtk-4.0-dev \
        appstream \
        fuse \
        just

RUN sudo wget https://github.com/AppImage/AppImageKit/releases/download/continuous/appimagetool-x86_64.AppImage -O /usr/local/bin/appimagetool \
    && sudo chmod +x /usr/local/bin/appimagetool
    
RUN go install github.com/wailsapp/wails/v2/cmd/wails@latest

ENV DESKTOP_SESSION xfce
