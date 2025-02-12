#!/bin/bash

_K8S_POD_NAME_=${PodName}
if [ -z "$_K8S_POD_NAME_" ]; then
  echo "got empty [PodName] env value"
  exit 255
fi

_K8S_POD_IP_=${PodIP}
if [ -z "$_K8S_POD_IP_" ]; then
  echo "got empty [PodIP] env value"
  exit 255
fi

declare -l _LISTEN_ADDRESS_=${_K8S_POD_NAME_}.tars-tarsregistry
echo "${_K8S_POD_IP_}" "${_LISTEN_ADDRESS_}" >>/etc/hosts

REGISTRY_EXECUTION_FILE=/usr/local/app/tars/tarsregistry/bin/tarsregistry

REGISTRY_CONFIG_FILE=/usr/local/app/tars/tarsregistry/conf/tarsregistry.conf

declare -a ReplaceKeyList=(
  _LISTEN_ADDRESS_
)

declare -a ReplaceFileList=(
  "${REGISTRY_CONFIG_FILE}"
)

for KEY in "${ReplaceKeyList[@]}"; do
  for FILE in "${ReplaceFileList[@]}"; do
    sed -i "s#${KEY}#${!KEY}#g" "${FILE}"
    if [[ 0 -ne $? ]]; then
      exit 255
    fi
  done
done

chmod +x ${REGISTRY_EXECUTION_FILE}
exec ${REGISTRY_EXECUTION_FILE} --config=${REGISTRY_CONFIG_FILE}
# exec /bin/guard
