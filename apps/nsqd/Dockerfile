FROM docker.5th.im:5000/alpine-golang
RUN mkdir -p /home/longbridge/config
RUN apk add curl
COPY ./nsqd /home/longbrige
WORKDIR /home/longbridge

CMD [ "./nsqd"]