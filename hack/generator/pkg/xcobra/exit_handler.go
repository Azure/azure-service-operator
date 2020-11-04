//+build !noexit

/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package xcobra

import (
	"errors"
	"os"
)

func exitWithCode(err error) {
	if err == nil {
		return
	}

	var e ErrorWithCode
	if errors.As(err, &e) {
		os.Exit(e.Code)
	}

	os.Exit(1)
}
