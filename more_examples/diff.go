package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"io/ioutil"
	"os/exec"
	"runtime"

	"github.com/r3labs/diff"
	"github.com/sergi/go-diff/diffmatchpatch"
)

// Returns diff of two arrays of bytes in diff tool format.
func Diff(prefix string, b1, b2 []byte) ([]byte, error) {
	f1, err := writeTempFile(prefix, b1)
	if err != nil {
		return nil, err
	}
	defer os.Remove(f1)

	f2, err := writeTempFile(prefix, b2)
	if err != nil {
		return nil, err
	}
	defer os.Remove(f2)

	cmd := "diff"
	if runtime.GOOS == "plan9" {
		cmd = "/bin/ape/diff"
	}

	data, err := exec.Command(cmd, "-u", f1, f2).CombinedOutput()
	if len(data) > 0 {
		// diff exits with a non-zero status when the files don't match.
		// Ignore that failure as long as we get output.
		err = nil
	}
	return data, err
}

func writeTempFile(prefix string, data []byte) (string, error) {
	file, err := ioutil.TempFile("", prefix)
	if err != nil {
		return "", err
	}
	_, err = file.Write(data)
	if err1 := file.Close(); err == nil {
		err = err1
	}
	if err != nil {
		os.Remove(file.Name())
		return "", err
	}
	return file.Name(), nil
}

type UpdateOpts struct {
	MaintenanceWindowStart        string `json:"maintenance_window_start,omitempty"`
	EnableAutorepair              *bool  `json:"enable_autorepair,omitempty"`
	EnablePatchVersionAutoUpgrade *bool  `json:"enable_patch_version_auto_upgrade,omitempty"`

	KubernetesOptions *KubernetesOptions `json:"kubernetes_options,omitempty"`
}

type KubernetesOptions struct {
	EnablePodSecurityPolicy bool `json:"enable_pod_security_policy"`
}

func main() {
	a := UpdateOpts{
		MaintenanceWindowStart: "04:00:00",
		KubernetesOptions: &KubernetesOptions{
			EnablePodSecurityPolicy: false,
		},
	}
	b := UpdateOpts{
		MaintenanceWindowStart: "07:00:00",
		KubernetesOptions: &KubernetesOptions{
			EnablePodSecurityPolicy: true,
		},
	}

	changelog, _ := diff.Diff(a, b)
	fmt.Println(changelog)

	buf1 := bytes.Buffer{}
	enc := gob.NewEncoder(&buf1)
	if err := enc.Encode(a); err != nil {
		log.Fatal(err)
	}

	buf2 := bytes.Buffer{}
	enc2 := gob.NewEncoder(&buf2)
	if err := enc2.Encode(b); err != nil {
		log.Fatal(err)
	}

	changelog2, _ := Diff("diff_test", buf1.Bytes(), buf2.Bytes())
	fmt.Println(string(changelog2))

	a1, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(a1))

	b1, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b1))

	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(string(a1), string(b1), false)

	fmt.Println(dmp.DiffPrettyText(diffs))
}
