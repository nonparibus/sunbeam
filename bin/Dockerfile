FROM rust:latest

RUN mkdir /build
WORKDIR /build


ENV PATH /root/.local/bin:${PATH}
RUN cargo install just

RUN git clone https://github.com/pop-os/launcher
RUN cd launcher && just && just install

RUN mkdir /root/.cache
ENTRYPOINT [ "pop-launcher" ]
