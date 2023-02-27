package gover

import "runtime"

func Version() string {
	return runtime.Version()
}
