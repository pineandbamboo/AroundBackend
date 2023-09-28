// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by aliasgen. DO NOT EDIT.

// Package clientgateways aliases all exported identifiers in package
// "cloud.google.com/go/beyondcorp/clientgateways/apiv1/clientgatewayspb".
//
// Deprecated: Please use types in: cloud.google.com/go/beyondcorp/clientgateways/apiv1/clientgatewayspb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package clientgateways

import (
	src "cloud.google.com/go/beyondcorp/clientgateways/apiv1/clientgatewayspb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/beyondcorp/clientgateways/apiv1/clientgatewayspb
const (
	ClientGateway_CREATING          = src.ClientGateway_CREATING
	ClientGateway_DELETING          = src.ClientGateway_DELETING
	ClientGateway_DOWN              = src.ClientGateway_DOWN
	ClientGateway_ERROR             = src.ClientGateway_ERROR
	ClientGateway_RUNNING           = src.ClientGateway_RUNNING
	ClientGateway_STATE_UNSPECIFIED = src.ClientGateway_STATE_UNSPECIFIED
	ClientGateway_UPDATING          = src.ClientGateway_UPDATING
)

// Deprecated: Please use vars in: cloud.google.com/go/beyondcorp/clientgateways/apiv1/clientgatewayspb
var (
	ClientGateway_State_name                                                     = src.ClientGateway_State_name
	ClientGateway_State_value                                                    = src.ClientGateway_State_value
	File_google_cloud_beyondcorp_clientgateways_v1_client_gateways_service_proto = src.File_google_cloud_beyondcorp_clientgateways_v1_client_gateways_service_proto
)

// Message describing ClientGateway object.
//
// Deprecated: Please use types in: cloud.google.com/go/beyondcorp/clientgateways/apiv1/clientgatewayspb
type ClientGateway = src.ClientGateway

// Represents the metadata of the long-running operation.
//
// Deprecated: Please use types in: cloud.google.com/go/beyondcorp/clientgateways/apiv1/clientgatewayspb
type ClientGatewayOperationMetadata = src.ClientGatewayOperationMetadata

// Represents the different states of a gateway.
//
// Deprecated: Please use types in: cloud.google.com/go/beyondcorp/clientgateways/apiv1/clientgatewayspb
type ClientGateway_State = src.ClientGateway_State

// ClientGatewaysServiceClient is the client API for ClientGatewaysService
// service. For semantics around ctx use and closing/ending streaming RPCs,
// please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/beyondcorp/clientgateways/apiv1/clientgatewayspb
type ClientGatewaysServiceClient = src.ClientGatewaysServiceClient

// ClientGatewaysServiceServer is the server API for ClientGatewaysService
// service.
//
// Deprecated: Please use types in: cloud.google.com/go/beyondcorp/clientgateways/apiv1/clientgatewayspb
type ClientGatewaysServiceServer = src.ClientGatewaysServiceServer

// Message for creating a ClientGateway.
//
// Deprecated: Please use types in: cloud.google.com/go/beyondcorp/clientgateways/apiv1/clientgatewayspb
type CreateClientGatewayRequest = src.CreateClientGatewayRequest

// Message for deleting a ClientGateway
//
// Deprecated: Please use types in: cloud.google.com/go/beyondcorp/clientgateways/apiv1/clientgatewayspb
type DeleteClientGatewayRequest = src.DeleteClientGatewayRequest

// Message for getting a ClientGateway.
//
// Deprecated: Please use types in: cloud.google.com/go/beyondcorp/clientgateways/apiv1/clientgatewayspb
type GetClientGatewayRequest = src.GetClientGatewayRequest

// Message for requesting list of ClientGateways.
//
// Deprecated: Please use types in: cloud.google.com/go/beyondcorp/clientgateways/apiv1/clientgatewayspb
type ListClientGatewaysRequest = src.ListClientGatewaysRequest

// Message for response to listing ClientGateways.
//
// Deprecated: Please use types in: cloud.google.com/go/beyondcorp/clientgateways/apiv1/clientgatewayspb
type ListClientGatewaysResponse = src.ListClientGatewaysResponse

// UnimplementedClientGatewaysServiceServer can be embedded to have forward
// compatible implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/beyondcorp/clientgateways/apiv1/clientgatewayspb
type UnimplementedClientGatewaysServiceServer = src.UnimplementedClientGatewaysServiceServer

// Deprecated: Please use funcs in: cloud.google.com/go/beyondcorp/clientgateways/apiv1/clientgatewayspb
func NewClientGatewaysServiceClient(cc grpc.ClientConnInterface) ClientGatewaysServiceClient {
	return src.NewClientGatewaysServiceClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/beyondcorp/clientgateways/apiv1/clientgatewayspb
func RegisterClientGatewaysServiceServer(s *grpc.Server, srv ClientGatewaysServiceServer) {
	src.RegisterClientGatewaysServiceServer(s, srv)
}