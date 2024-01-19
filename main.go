package main

import (
	"context"
)

type Grype struct{}

// example usage: "dagger call test-container-image --image alpine:latest"
func (m *Grype) TestContainerImage(ctx context.Context, image string) (string, error) {
	containerUnderTest := dag.Container().From(image)
	return m.TestContainer(ctx, containerUnderTest)
}

// example usage: from dagger code where you have a Container object: TestContainer(container)
func (m *Grype) TestContainer(ctx context.Context, container *Container) (string, error) {
	return dag.Container().
		From("anchore/grype:latest").
		WithFile("/image.tar", container.AsTarball()).
		WithExec([]string{"/image.tar"}).
		Stdout(ctx)
}
