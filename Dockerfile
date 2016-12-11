FROM scratch
COPY ./config.yaml /config.yaml
COPY ./redirector /redirector
ENTRYPOINT ["/redirector"]
