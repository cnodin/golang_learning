package pipeline_test

import (
	"bytes"
	"os/exec"
	"testing"
)

func Test_Pipeline(t *testing.T)  {
	t.Log("testing")
	cmd1 := exec.Command("ps", "au")
	cmd2 := exec.Command("grep", "apipe")

	var outputBuf1 bytes.Buffer
	cmd1.Stdout = &outputBuf1
	if err := cmd1.Start(); err != nil {
		t.Logf("Error: The first command can not be startup %s\n", err)
		return
	}

	if err := cmd1.Wait(); err != nil {
		t.Logf("Error: Couldn't wait for the first command: %s\n", err)
		return
	}

	cmd2.Stdin = &outputBuf1
	var outputBuf2 bytes.Buffer
	cmd2.Stdout = &outputBuf2

	if err := cmd2.Start(); err != nil {
		t.Logf("Error: The second command can not be startup: %s\n", err)
		return
	}
	if err := cmd2.Wait(); err != nil {
		t.Logf("Error: Couldn't wait for the second command: %s\n", err)
		return
	}

	t.Logf("%s\n", outputBuf2.Bytes())
}
