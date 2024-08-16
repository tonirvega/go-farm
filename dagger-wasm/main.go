// A generated module for Dagger functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/dagger/internal/dagger"
)

type DaggerWasm struct{}

func (m *DaggerWasm) ModoWasm(
	ctx context.Context,
	projectDir *dagger.Directory,
) *dagger.Container {
	containerGolang := m.
		CompileWasm(ctx, projectDir)

	return m.
		ServeWasm(ctx,
			containerGolang.
				Directory("/app").
				File("json.wasm"))
}

func (m *DaggerWasm) ModoEscritorio(
	ctx context.Context,

	projectDir *dagger.Directory,

	// +optional
	// +default="true"
	flagDebug string,

	// +optional
	// +default="ENABLED"
	flagDesktop string,

) *dagger.Container {
	return dag.Container().
		From("golang").
		WithWorkdir("/app").
		WithMountedDirectory("/app", projectDir).
		WithEnvVariable("DEBUG_MODE", flagDebug).
		WithEnvVariable("DESKTOP", flagDesktop).
		WithExec([]string{"go", "build"}).
		Terminal(
			dagger.ContainerTerminalOpts{
				Cmd: []string{"go", "run", "."},
			},
		)
}

func (m *DaggerWasm) ServeWasm(ctx context.Context, wasmBinary *dagger.File) *dagger.Container {

	indexHtml := dag.CurrentModule().
		Source().
		File("index.html")

	jsExecFile := dag.CurrentModule().
		Source().
		File("wasm_exec.js")

	return dag.Container().
		From("nginx").
		WithMountedFile("/usr/share/nginx/html/json.wasm", wasmBinary).
		WithMountedFile("/usr/share/nginx/html/wasm_exec.js", jsExecFile).
		WithMountedFile("/usr/share/nginx/html/index.html", indexHtml).
		WithExposedPort(80)
}

// Returns lines that match a pattern in the files of the provided Directory
func (m *DaggerWasm) CompileWasm(ctx context.Context, projectDir *dagger.Directory) *dagger.Container {
	return dag.Container().
		From("golang").
		WithWorkdir("/app").
		WithMountedDirectory("/app", projectDir).
		WithEnvVariable("GOOS", "js").
		WithEnvVariable("GOARCH", "wasm").
		WithExec([]string{"go", "build", "-o", "json.wasm"})
}
