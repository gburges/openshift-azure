#!/bin/bash -ex

cleanup() {
    kill $(jobs -p) &>/dev/null || true
    wait
}

trap cleanup EXIT

if [[ $# -eq 0 && ! -e _data/containerservice.yaml ]]; then
    echo error: _data/containerservice.yaml must exist
    exit 1
fi

if [[ $# -eq 1 ]]; then
    export RESOURCEGROUP=$1
else
    export RESOURCEGROUP=$(awk '/^    resourceGroup:/ { print $2 }' <_data/containerservice.yaml)
fi

if [[ -n "$TEST_IN_PRODUCTION" ]]; then
    TEST_IN_PRODUCTION="-use-prod=true"
else
    go generate ./...
    go run cmd/fakerp/main.go &
fi
if [[ -n "$ADMIN_MANIFEST" ]]; then
    ADMIN_MANIFEST="-admin-manifest=$ADMIN_MANIFEST"
fi

go run cmd/createorupdate/createorupdate.go -timeout 1h ${TEST_IN_PRODUCTION:-} ${ADMIN_MANIFEST:-}
