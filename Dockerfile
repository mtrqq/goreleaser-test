FROM scratch
ENTRYPOINT ["/usr/bin/goexec"]
COPY goexec /usr/bin/goexec
