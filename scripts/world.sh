#!/bin/bash

BASEDIR="$(
    cd -- "$(dirname "$0")" >/dev/null 2>&1
    pwd -P
)/../"

STORAGE_DIR=${STORAGE_DIR:-${BASEDIR}/build}

function RunInContainer() {
    ${BASEDIR}/scripts/exec.sh ${@}
}

ALL_PKGS=$(find ${STORAGE_DIR}/${VERSION}/recipes/ -type f -name "*.yml" -exec basename {} \; | sed 's|.yml||g')

echo ":: generating dependency tree ::"
DEPENDENCY_TREE=$(RunInContainer pkgupd depends --force ${ALL_PKGS})
if [[ $? != 0 ]] ; then
    echo "failed to build dependency tree ${DEPENDENCY_TREE}"
    exit 1
fi

for pkg in ${DEPENDENCY_TREE} ; do
    pkg=$(echo ${pkg} | tr -cd '[:print:]')
    echo ":: compiling ${pkg} ::"
    RunInContainer pkgupd co ${pkg}
done
