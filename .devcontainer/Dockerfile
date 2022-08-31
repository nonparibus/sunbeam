FROM kasmweb/core-ubuntu-focal:1.11.0
USER root

ENV HOME /home/kasm-default-profile
ENV STARTUPDIR /dockerstartup
ENV INST_SCRIPTS $STARTUPDIR/install
WORKDIR $HOME

RUN apt-get update
RUN apt-get install -y sudo \
    && echo 'kasm-user ALL=(ALL) NOPASSWD: ALL' >> /etc/sudoers \
    && rm -rf /var/lib/apt/list/*

# nodejs
ENV PATH $HOME/go/bin:$PATH
RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash - \
    && apt-get install -y nodejs

# golang
RUN add-apt-repository ppa:longsleep/golang-backports \
    && apt-get update \
    && apt-get install -y golang-go

# gtk dependencies
RUN apt-get update \
    && apt install -y libgtk-3-dev libwebkit2gtk-4.0-dev


# pop-launcher
RUN add-apt-repository ppa:system76/pop \
    && apt-get update \
    && apt-get install -y pop-launcher

RUN chown 1000:0 $HOME
RUN $STARTUPDIR/set_user_permission.sh $HOME

RUN mkdir -p $HOME && chown -R 1000:0 $HOME

USER 1000
