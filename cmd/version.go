/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package cmd

import (
	"fmt"

	"github.com/IBAX-io/go-ibax/packages/consts"

	Short: "Show version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(consts.Version())
	},
}
