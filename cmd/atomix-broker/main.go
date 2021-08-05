// Copyright 2019-present Open Networking Foundation.
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

package main

import (
	"fmt"
	protocolapi "github.com/atomix/atomix-api/go/atomix/protocol"
	"github.com/atomix/atomix-go-framework/pkg/atomix/broker"
	"github.com/atomix/atomix-go-framework/pkg/atomix/cluster"
	"github.com/atomix/atomix-go-framework/pkg/atomix/logging"
	"github.com/spf13/cobra"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logging.SetLevel(logging.InfoLevel)

	cmd := &cobra.Command{
		Use: "atomix-broker",
	}

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Create a new broker node
	broker := broker.NewBroker(cluster.NewCluster(
		cluster.NewNetwork(),
		protocolapi.ProtocolConfig{},
		cluster.WithMemberID("atomix-broker"),
		cluster.WithPort(5678)))

	// Start the node
	if err := broker.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Wait for an interrupt signal
	ch := make(chan os.Signal, 2)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	<-ch

	// Stop the node after an interrupt
	if err := broker.Stop(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
