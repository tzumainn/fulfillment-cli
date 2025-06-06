/*
Copyright (c) 2025 Red Hat Inc.

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the
License. You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific
language governing permissions and limitations under the License.
*/

package logout

import (
	"fmt"

	"github.com/innabox/fulfillment-cli/internal/config"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	runner := &runnerContext{}
	result := &cobra.Command{
		Use:   "logout [flags]",
		Short: "Discard connection and authentication details",
		RunE:  runner.run,
	}
	return result
}

type runnerContext struct {
}

func (c *runnerContext) run(cmd *cobra.Command, args []string) error {
	// Load the configuration:
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}
	if cfg == nil {
		cfg = &config.Config{}
	}

	// Clear all the details:
	cfg.Token = ""
	cfg.Plaintext = false
	cfg.Insecure = false
	cfg.Address = ""

	// Save the configuration:
	err = config.Save(cfg)
	if err != nil {
		return fmt.Errorf("failed to save configuration: %w", err)
	}

	return nil
}
