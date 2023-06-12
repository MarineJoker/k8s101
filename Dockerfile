FROM ubuntu
ADD bin/main /httpserver
EXPOSE 80
ENTRYPOINT ./httpserver
