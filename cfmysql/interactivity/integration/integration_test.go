package main_test

import (
	"io/ioutil"
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Interactivity Detection Feature", func() {
	BeforeEach(func() {
		err := exec.Command("go", "build", ".").Run()
		Expect(err).To(BeNil())
	})

	It("detects Pipe mode when STDOUT is not interactive", func() {
		nonInteractiveOut, err := os.Create("stdout.txt")
		Expect(err).To(BeNil())

		cwd, err := os.Getwd()
		Expect(err).To(BeNil())
		processAttributes := os.ProcAttr{
			Files: []*os.File{os.Stdin, nonInteractiveOut, os.Stderr},
			Dir:   cwd,
		}
		process, err := os.StartProcess("./integration", nil, &processAttributes)
		Expect(err).To(BeNil())
		state, err := process.Wait()
		Expect(err).To(BeNil())
		Expect(state.Success()).To(BeTrue())
		buffer, err := ioutil.ReadFile("./output.txt")
		Expect(err).To(BeNil())
		output := string(buffer)

		Expect(output).To(ContainSubstring("pipe mode"))
		Expect(output).NotTo(ContainSubstring("interactive mode"))
	})

	It("detects Pipe mode when STDIN is not interactive", func() {
		nonInteractiveIn, err := os.Create("stdin.txt")
		Expect(err).To(BeNil())

		cwd, err := os.Getwd()
		Expect(err).To(BeNil())
		processAttributes := os.ProcAttr{
			Files: []*os.File{nonInteractiveIn, os.Stdout, os.Stderr},
			Dir:   cwd,
		}
		process, err := os.StartProcess("./integration", nil, &processAttributes)
		Expect(err).To(BeNil())
		state, err := process.Wait()
		Expect(err).To(BeNil())
		Expect(state.Success()).To(BeTrue())
		buffer, err := ioutil.ReadFile("./output.txt")
		Expect(err).To(BeNil())
		output := string(buffer)

		Expect(output).To(ContainSubstring("pipe mode"))
		Expect(output).NotTo(ContainSubstring("interactive mode"))
	})

	It("detects Interactive mode when both STDIN and STDOUT are interactive", func() {
		cwd, err := os.Getwd()
		Expect(err).To(BeNil())
		processAttributes := os.ProcAttr{
			Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
			Dir:   cwd,
		}
		process, err := os.StartProcess("./integration", nil, &processAttributes)
		Expect(err).To(BeNil())
		state, err := process.Wait()
		Expect(err).To(BeNil())
		Expect(state.Success()).To(BeTrue())
		buffer, err := ioutil.ReadFile("./output.txt")
		Expect(err).To(BeNil())
		output := string(buffer)

		Expect(output).To(ContainSubstring("interactive mode"))
		Expect(output).NotTo(ContainSubstring("pipe mode"))
	})
})
