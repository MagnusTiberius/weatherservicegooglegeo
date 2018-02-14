FROM golang

ADD ./weatherservicegooglegeo.exe /usr/local/bin/

ENTRYPOINT /usr/local/bin/weatherservicegooglegeo.exe

EXPOSE 8090
