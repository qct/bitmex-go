#!/bin/bash

alias cp='cp'

pwd=$(cd `dirname $0`; pwd);cd $pwd

official_home="/Users/quchentao/playground/btc/api-connectors/clients/go"

for file in `ls $pwd/swagger`; do
  echo "synchronizing $file"
  cp $official_home/$file $pwd/swagger
done