#!/bin/bash -e
STORAGE_DIR=${STORAGE_DIR:-"$(dirname ${PWD})/storage.rlxos.dev"}

docker run                                                          \
    --privileged                                                    \
    -v ${PWD}:/rlxos                                                \
    -v ${PWD}/pkgupd.yml:/etc/pkgupd.yml                            \
    -v /var/run/docker.sock:/var/run/docker.sock                    \
    -v ${STORAGE_DIR}:/storage                                      \
    -w /rlxos                                                       \
    -it itsmanjeet/devel-docker:2200-3                              \
    bash