FROM alpine

COPY ./bin/service-starter-kit /bin/service-starter-kit

EXPOSE 80

ENTRYPOINT [ "/bin/service-starter-kit" ]
