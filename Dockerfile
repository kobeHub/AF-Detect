FROM scratch 
MAINTAINER Inno Jia <jdgets111@gmail.com>

#WORKDIR /backend
ENV TZ=Asia/Shanghai
COPY ./Upload /
COPY ./model /
ADD app /
ADD predict /
ADD predict.py /

# private and public 
EXPOSE 8081

# private only
# EXPOSE 8080

ENV GIN_MODE=release
CMD ["/app"] 
