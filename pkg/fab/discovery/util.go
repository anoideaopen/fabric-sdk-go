/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package discovery

import (
	"reflect"

	discclient "github.com/hyperledger/fabric-sdk-go/internal/github.com/hyperledger/fabric/discovery/client"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
)

// GetProperties extracts the properties from the discovered peer.
func GetProperties(endpoint *discclient.Peer) fab.Properties {
	if endpoint.StateInfoMessage == nil {
		return nil
	}

	stateInfo := endpoint.StateInfoMessage.GetStateInfo()
	if stateInfo == nil || stateInfo.Properties == nil {
		return nil
	}

	properties := make(fab.Properties)

	val := reflect.Indirect(reflect.ValueOf(stateInfo.Properties))
	elem := reflect.ValueOf(stateInfo.Properties)

	for i := 0; i < val.Type().NumField(); i++ {
		// Exclude protobuf fields
		if val.Type().Field(i).IsExported() {
			properties[val.Type().Field(i).Name] = elem.Elem().Field(i).Interface()
		}
	}

	return properties
}
