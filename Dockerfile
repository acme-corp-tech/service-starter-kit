FROM alpine

#HEALTHCHECK --interval=20s --timeout=3s --start-period=5s --retries=3 \
#  CMD /bin/brick-starter-kit -health

COPY ./bin/brick-starter-kit /bin/brick-starter-kit

EXPOSE 80

ENTRYPOINT [ "/bin/brick-starter-kit" ]
