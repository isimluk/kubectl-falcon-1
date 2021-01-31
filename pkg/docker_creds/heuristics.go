package docker_creds

import (
	"fmt"

	"github.com/containers/image/v5/types"
)

const AWS_EKR_RE = `123456.dkr.ecr.region.amazonaws.com/`

// BestGuess returns docker registry crendentials based on the imageUri. Credentials are fetched
// from various sources to mimic docker credStore behavior. Docker credStore is not readily
// available outside of docker command yet: See https://github.com/moby/moby/issues/39377
// https://github.com/containers/image/pull/656
func BestGuess(imageRef types.ImageReference) (*types.DockerAuthConfig, error) {
	fmt.Println("TODO")
	fmt.Println(imageRef.DockerReference())
	return nil, nil
}
