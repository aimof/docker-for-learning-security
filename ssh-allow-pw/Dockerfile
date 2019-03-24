FROM alpine:3.9.2

LABEL maintainer="aimof(aimof.aimof@gmail.com)"

RUN apk --no-cache update && \
    apk --no-cache add openssh openrc syslog-ng && \
    ssh-keygen -A && \
    rc-update add sshd && \
    rc-update add syslog-ng && \
    echo root:poor | chpasswd && \
    touch /var/log/auth.log

ADD public_key/sample_key.pub /root/.ssh/authorized_keys

CMD tail -f /var/log/auth.log