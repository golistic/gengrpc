package gengrpc

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
)

// Generate will generate Go code by compiling .proto files found within
// the sourcePackages in the directory sourcePath. The resulting files will be
// source in directory destPath.
func Generate(module string, source sourcer, destPath string) error {
	protoc, err := exec.LookPath("protoc")
	if err != nil {
		return fmt.Errorf("protoc executable not available")
	}

	contracts, err := source.Contracts()
	if err != nil {
		return err
	}

	args := []string{protoc, "--proto_path=" + destPath,
		"--go_out=.",
		"--go_opt=paths=import",
		"--go_opt=module=" + module,
		"--proto_path=" + source.ContractPath(),
	}

	for _, f := range contracts {
		m := filepath.Base(filepath.Dir(f))
		args = append(args, fmt.Sprintf("--go_opt=M%s=%s/%s/%s", f, module, destPath, m))
	}

	args = append(args, contracts...)

	output := bytes.NewBuffer(nil)

	cmd := exec.Cmd{
		Path:   protoc,
		Args:   args,
		Stdout: output,
		Stderr: output,
	}

	err = cmd.Run()
	switch err.(type) {
	case *exec.ExitError:
		return fmt.Errorf("failed executing protoc (%w)", fmt.Errorf(output.String()))
	case error:
		return fmt.Errorf("failed executing protoc (%w)", err)
	}

	return nil
}
