FROM ubuntu:18.04

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && \
    apt-get -y install wget gcc make libdb5.3-dev libsasl2-dev groff-base

WORKDIR /root

RUN wget ftp://ftp.openldap.org/pub/OpenLDAP/openldap-release/openldap-2.4.49.tgz && \
    tar xvzf openldap-2.4.49.tgz

WORKDIR openldap-2.4.49

RUN ./configure --with-tls -enable-crypt --with-cyrus-sasl -enable-rwm --enable-monitor=no && make depend && make && make install

ADD ./slapd.conf /usr/local/etc/openldap/slapd.conf
ADD ./rsyslog.conf /etc/rsyslog.conf

WORKDIR /root

RUN rm -rf ./openldap-2.4.49 ./openldap-2.4.49.tgz 

CMD ["bash"]
