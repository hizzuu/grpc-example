#!/bin/sh

docker run --rm -v ${PWD}:/proto hizzuu/protoc \
  -I/proto \
  -I/go/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.2 \
  --doc_opt=markdown,README.md \
  user.proto
