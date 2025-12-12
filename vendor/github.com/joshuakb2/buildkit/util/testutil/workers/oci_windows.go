package workers

import "github.com/joshuakb2/buildkit/util/bklog"

func initOCIWorker() {
	bklog.L.Info("OCI Worker not supported on Windows.")
}
