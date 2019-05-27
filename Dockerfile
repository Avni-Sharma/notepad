FROM registry.access.redhat.com/ubi7/ubi-minimal:latest

COPY ./notepad /usr/local/bin/notepad
USER 1001

ENTRYPOINT [ "/usr/local/bin/notepad" ]
