FROM gitpod/workspace-full-vnc

RUN sudo apt-get update \
    && sudo apt-get -y install libgtk-3-dev libwebkit2gtk-4.0-dev

RUN curl -q 'https://proget.makedeb.org/debian-feeds/prebuilt-mpr.pub' \
        | gpg --dearmor \
        | sudo tee /usr/share/keyrings/prebuilt-mpr-archive-keyring.gpg 1> /dev/null \
    && echo "deb [signed-by=/usr/share/keyrings/prebuilt-mpr-archive-keyring.gpg] https://proget.makedeb.org prebuilt-mpr $(lsb_release -cs)" \
        | sudo tee /etc/apt/sources.list.d/prebuilt-mpr.list \
    && sudo apt update \
    && sudo apt install just

# Install pop launcher system wide
RUN git clone https://github.com/pop-os/launcher \
    && cd launcher \
    && just \
    && sudo just rootdir=/ install

ENV DESKTOP_SESSION xfce
