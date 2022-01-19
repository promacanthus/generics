# !/usr/bin/env zsh

set -e

# 模板源文件

SRC_FILE=${1}

# 包名

PACKAGE=${2}

# 实际需要具体化的类型

TYPE=${3}

# 用于构造目标文件的后缀

DES=${4}

# uppcase the first char

PREFIX="$(tr '[:lower:]' '[:upper:]' <<< ${TYPE:0:1})${TYPE:1}"

DES_FILE=$(echo ${TYPE}| tr '[:upper:]' '[:lower:]')_${DES}.go

sed 's/PACKAGE_NAME/'"${PACKAGE}"'/g' ${SRC_FILE} | \

sed 's/GENERIC_TYPE/'"${TYPE}"'/g' | \

sed 's/GENERIC_NAME/'"${PREFIX}"'/g' > ${DES_FILE}