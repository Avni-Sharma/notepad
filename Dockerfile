FROM registry.access.redhat.com/ubi7/ubi:latest

COPY ./notepad /usr/local/bin/notepad
USER 1001

ENTRYPOINT [ "/usr/local/bin/notepad" ]
