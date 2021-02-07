package mysqlutils

import (
	"fmt"
	"io"
	"os/exec"
	"strconv"
)

func ImportSQL(opts ConnOptions, data io.Reader) error {
	cmd := exec.Command("mysql", "-v",
		"-h", opts.Host, "-P", strconv.Itoa(opts.Port),
		"-u", opts.User, "-p"+opts.Password,
		opts.DBName)
	cmd.Stdin = data

	if out, err := cmd.Output(); err != nil {
		return fmt.Errorf("error importing SQL '%s': %w", string(out), err)
	}

	return nil
}
