#!/usr/bin/env bash


ROOT_DIR=$(dirname "${BASH_SOURCE[0]}")/..
source "${ROOT_DIR}/hack/common.sh"

if [ $# -ne 4 ];then
    onex::log::error "Usage: gen-service-config.sh SERVICE_NAME ENV_FILE TEMPLATE_FILE OUTPUT_DIR"
    exit 1
fi

export SERVICE_NAME=$1
ENV_FILE=$2
TEMPLATE_FILE=$3
OUTPUT_DIR=$4

echo TEMPLATE_FILE=${TEMPLATE_FILE}
echo OUTPUT_DIR=${OUTPUT_DIR}
echo ENV_FILE=${ENV_FILE}
echo SERVICE_NAME=${SERVICE_NAME}

if [ ! -d ${OUTPUT_DIR} ];then
  mkdir -p ${OUTPUT_DIR}
fi

#  判断SERVICE_NAME是否在ALL_COMPONENTS中
if [[ " ${ALL_COMPONENTS[*]} " != *" ${SERVICE_NAME} "* ]]; then
  echo "Invalid SERVICE_NAME:${SERVICE_NAME}"
  exit 0
fi

function get_apiserver_systemd_file()
{
  cat << EOF
# onex-apiserver.service generated by scripts/gen-systemd.sh. DO NOT EDIT.
[Unit]
Description=Systemd unit file for onex-apiserver
Documentation=https://github.com/superproj/onex/blob/master/manifests/installation/onex/systemd/README.md

[Service]
WorkingDirectory=${INSTALL_DIR}
ExecStartPre=/usr/bin/mkdir -p ${DATA_DIR}/onex-apiserver
ExecStartPre=/usr/bin/mkdir -p ${LOG_DIR}
ExecStart=/opt/onex/bin/onex-apiserver --bind-address=${ONEX_APISERVER_BIND_ADDRESS} --secure-port ${ONEX_APISERVER_SECURE_PORT} --etcd-servers ${ONEX_APISERVER_ETCD_SERVERS} --client-ca-file=${ONEX_APISERVER_CLIENT_CA_FILE} --tls-cert-file=${ONEX_APISERVER_TLS_CERT_FILE} --tls-private-key-file=${ONEX_APISERVER_TLS_PRIVATE_KEY_FILE} --v=${ONEX_APISERVER_V_LEVEL}
Restart=always
RestartSec=5
StartLimitInterval=0

[Install]
WantedBy=multi-user.target
EOF
}

echo "ENV_FILE=${ENV_FILE}"

source ${ENV_FILE}

echo TEMPLATE_FILE=${TEMPLATE_FILE}

# Some customized processing
case ${TEMPLATE_FILE} in
  *systemd.tmpl.service)
    echo "Generating systemd file for ${SERVICE_NAME}"
    if [ "${SERVICE_NAME}" == "pay-server" ]; then
      get_apiserver_systemd_file > ${OUTPUT_DIR}/${SERVICE_NAME}.service
    fi
    if [ "${SERVICE_NAME}" == "ctl" ];then
      exit 0
    fi
    ;;
  *onex-apiserver.config.tmpl.yaml)
    exit 0
    ;;
  *)
    ;;
esac


#  生成配置文件
suffix=$(echo $TEMPLATE_FILE | awk -F'.' '{print $NF}')
${ROOT_DIR}/hack/gen-config.sh ${ENV_FILE} ${TEMPLATE_FILE} > ${OUTPUT_DIR}/${SERVICE_NAME}.${suffix}

#  为onex-apiserver.service添加kubeconfig
if [[ "${TEMPLATE_FILE}" =~ .*systemd.tmpl.service ]] && [[ "${SERVICE_NAME}" =~ *server ]];then
  escaped_config_dir="$(sed -e 's/[\/&]/\\&/g' <<< "${CONFIG_DIR}")"
  echo "Adding kubeconfig to ${escaped_config_dir} service"
  sed -i "/ExecStart=/s/$/ --kubeconfig=${escaped_config_dir}\/config/" ${OUTPUT_DIR}/${SERVICE_NAME}.${suffix}
fi