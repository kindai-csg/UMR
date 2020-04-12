#!/bin/bash

set -e

docker run -it --rm -v $PWD/client:/src -w /src node:13.8 yarn install
docker run -it --rm -v $PWD/client:/src -w /src node:13.8 yarn build 

docker-compose up -d

docker-compose exec ldap ldapadd -x -D "cn=Manager,dc=kindai-csg,dc=dev" -W -f /usr/local/etc/openldap/schema/init.ldif