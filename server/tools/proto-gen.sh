#!/bin/bash
set -e

# Root directories
REPO_ROOT=$(cd "$(dirname "${BASH_SOURCE[0]}")"/../.. && pwd)
PROTO_DIR="${REPO_ROOT}/server/proto"
PB_OUT="${REPO_ROOT}/server/pb"
SWAGGER_OUT="${REPO_ROOT}/server/swagger-ui"

# grpc-gateway + googleapis
GATEWAY_DIR="${REPO_ROOT}/server/tools/contrib/grpc-gateway"
GOOGLEAPIS_DIR="${REPO_ROOT}/server/tools/contrib/googleapis"

# Ensure output directory exists
mkdir -p "${PB_OUT}"

echo "Generating gRPC + gateway + swagger..."
protoc \
  -I "${PROTO_DIR}" \
  -I "${GOOGLEAPIS_DIR}" \
  -I "${GATEWAY_DIR}" \
  --go_out="${PB_OUT}" --go_opt=paths=source_relative \
  --go-grpc_out="${PB_OUT}" --go-grpc_opt=paths=source_relative \
  --grpc-gateway_out="${PB_OUT}" --grpc-gateway_opt=paths=source_relative \
  --openapiv2_out=allow_merge=true,merge_file_name=api:"${SWAGGER_OUT}" \
  "${PROTO_DIR}"/*.proto

echo "✅ Done. Output in:"
echo "  - Go files:       ${PB_OUT}"
echo "  - OpenAPI:        ${SWAGGER_OUT}/api.swagger.json"