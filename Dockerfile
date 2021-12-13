FROM alpine

#HEALTHCHECK --interval=20s --timeout=3s --start-period=5s --retries=3 \
#  CMD /bin/service-starter-kit -health

COPY ./bin/service-starter-kit /bin/service-starter-kit

EXPOSE 80

ENTRYPOINT [ "/bin/service-starter-kit" ]
