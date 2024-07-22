// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/hyperledger/fabric-sdk-go/test/performance

replace github.com/hyperledger/fabric-sdk-go => ../../

require (
	github.com/golang/protobuf v1.5.4
	github.com/hyperledger/fabric-protos-go v0.3.0
	github.com/hyperledger/fabric-sdk-go v0.0.0-00010101000000-000000000000
	github.com/hyperledger/fabric-sdk-go/test/integration v0.0.0-20200909154308-842c4b3ea51e // indirect
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.8.3
	golang.org/x/net v0.23.0
	google.golang.org/grpc v1.57.2
)

go 1.14
