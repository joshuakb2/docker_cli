package container

import "github.com/joshuakb2/moby/api/types/common"

// CommitResponse response for the commit API call, containing the ID of the
// image that was produced.
type CommitResponse = common.IDResponse
