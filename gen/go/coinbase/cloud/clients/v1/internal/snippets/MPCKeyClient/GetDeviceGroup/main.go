// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

// [START api_v1_generated_MPCKeyService_GetDeviceGroup_sync]

package main

import (
	"context"

	v1 "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/clients/v1"
	mpc_keyspb "github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1"
)

func main() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := v1.NewMPCKeyClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &mpc_keyspb.GetDeviceGroupRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/github.com/coinbase/waas-client-library-go/gen/go/coinbase/cloud/mpc_keys/v1#GetDeviceGroupRequest.
	}
	resp, err := c.GetDeviceGroup(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END api_v1_generated_MPCKeyService_GetDeviceGroup_sync]
