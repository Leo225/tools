// Copyright © 2023 OpenIM. All rights reserved.
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

package discovery

import (
	"context"

	"google.golang.org/grpc"
)

//
//type Conn interface {
//	GetConns(ctx context.Context, serviceName string, opts ...grpc.DialOption) ([]*grpc.ClientConn, error) //1
//	GetConn(ctx context.Context, serviceName string, opts ...grpc.DialOption) (*grpc.ClientConn, error)    //2
//	GetSelfConnTarget() string                                                                             //3
//	AddOption(opts ...grpc.DialOption)                                                                     //4
//	CloseConn(conn *grpc.ClientConn)                                                                       //5
//	// do not use this method for call rpc
//
//	GetClientLocalConns() map[string][]*grpc.ClientConn //del
//
//	GetUserIdHashGatewayHost(ctx context.Context, userId string) (string, error) //del
//}
//
//type SvcDiscoveryRegistry interface {
//	Conn
//	Register(serviceName, host string, port int, opts ...grpc.DialOption) error //6
//	UnRegister() error                                                          //7
//	RegisterConf2Registry(key string, conf []byte) error                        //del
//	GetConfFromRegistry(key string) ([]byte, error)                             //del
//	Close()                                                                     //
//}

type Conn interface {
	GetConn(ctx context.Context, serviceName string, opts ...grpc.DialOption) (grpc.ClientConnInterface, error)
	GetConns(ctx context.Context, serviceName string, opts ...grpc.DialOption) ([]grpc.ClientConnInterface, error)
	IsSelfNode(cc grpc.ClientConnInterface) bool
}

type SvcDiscoveryRegistry interface {
	Conn
	AddOption(opts ...grpc.DialOption)
	Register(serviceName, host string, port int, opts ...grpc.DialOption) error //6
	UnRegister() error                                                          //7
	Close()
	GetUserIdHashGatewayHost(ctx context.Context, userId string) (string, error) //
}
