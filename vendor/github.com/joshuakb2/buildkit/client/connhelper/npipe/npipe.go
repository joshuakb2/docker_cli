// Package npipe provides connhelper for npipe://<address>
package npipe

import "github.com/joshuakb2/buildkit/client/connhelper"

func init() {
	connhelper.Register("npipe", Helper)
}
