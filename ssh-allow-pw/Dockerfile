FROM alpine:3.9.2

LABEL maintainer="aimof(aimof.aimof@gmail.com)"

RUN apk --no-cache update && \
    apk --no-cache add openssh openrc && \
    ssh-keygen -A && \
    echo root:poor | chpasswd && \
    echo "PermitRootLogin yes" >> /etc/ssh/sshd_config && \
    sed s/session\srequired\spam/session\soptional\spam_loginuid.so/g >> /etc/ssh/sshd_config

ADD public_key/sample_key.pub /root/.ssh/authorized_keys

RUN chmod -R 600 /root/.ssh

CMD /usr/sbin/sshd -D