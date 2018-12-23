package simpletracker_test

import (
	. "github.com/dgruber/drmaa2os/pkg/jobtracker/simpletracker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/dgruber/drmaa2interface"
)

var _ = Describe("Jobstore", func() {

	Context("Basic JobStore operations", func() {

		It("should be possible to create a JobStore, save a job, and get the PID", func() {
			store := NewJobStore()
			Ω(store).ShouldNot(BeNil())
			store.SaveJob("13", drmaa2interface.JobTemplate{RemoteCommand: "rc"}, 77)
			store.SaveJob("1", drmaa2interface.JobTemplate{RemoteCommand: "rc2"}, 13)
			store.SaveJob("12", drmaa2interface.JobTemplate{RemoteCommand: "rc3"}, 10)
			pid, err := store.GetPID("12")
			Ω(err).Should(BeNil())
			Ω(pid).Should(BeNumerically("==", 10))
			pid, err = store.GetPID("1")
			Ω(err).Should(BeNil())
			Ω(pid).Should(BeNumerically("==", 13))
			pid, err = store.GetPID("13")
			Ω(err).Should(BeNil())
			Ω(pid).Should(BeNumerically("==", 77))
		})

		It("should find PID of array job task", func() {
			store := NewJobStore()
			Ω(store).ShouldNot(BeNil())
			store.SaveArrayJob("13",
				[]int{77, 78, 79},
				drmaa2interface.JobTemplate{RemoteCommand: "rc"},
				1, 3, 1)
			store.SaveJob("13.1", drmaa2interface.JobTemplate{RemoteCommand: "rc"}, 77)
			store.SaveJob("13.2", drmaa2interface.JobTemplate{RemoteCommand: "rc"}, 78)
			store.SaveJob("13.3", drmaa2interface.JobTemplate{RemoteCommand: "rc"}, 79)
			pid, err := store.GetPID("13.1")
			Ω(err).Should(BeNil())
			Ω(pid).Should(BeNumerically("==", 77))
			pid, err = store.GetPID("13.2")
			Ω(err).Should(BeNil())
			Ω(pid).Should(BeNumerically("==", 78))
			pid, err = store.GetPID("13.3")
			Ω(err).Should(BeNil())
			Ω(pid).Should(BeNumerically("==", 79))
		})

		It("should error when job is not found", func() {
			store := NewJobStore()
			Ω(store).ShouldNot(BeNil())
			pid, err := store.GetPID("12")
			Ω(err).ShouldNot(BeNil())
			Ω(pid).Should(BeNumerically("==", -1))
		})

		It("should error when job id is wrong", func() {
			store := NewJobStore()
			Ω(store).ShouldNot(BeNil())
			pid, err := store.GetPID("12.asdf")
			Ω(err).ShouldNot(BeNil())
			Ω(pid).Should(BeNumerically("==", -1))
			pid, err = store.GetPID("..")
			Ω(err).ShouldNot(BeNil())
			Ω(pid).Should(BeNumerically("==", -1))
			store.SaveJob("13.2", drmaa2interface.JobTemplate{RemoteCommand: "rc"}, 77)
			pid, err = store.GetPID("13.asdf")
			Ω(err).ShouldNot(BeNil())
			Ω(pid).Should(BeNumerically("==", -1))
		})

		It("should error when task is not found", func() {
			store := NewJobStore()
			Ω(store).ShouldNot(BeNil())
			store.SaveJob("13.1", drmaa2interface.JobTemplate{RemoteCommand: "rc"}, 77)
			pid, err := store.GetPID("13.77")
			Ω(err).ShouldNot(BeNil())
			Ω(pid).Should(BeNumerically("==", -1))
			pid, err = store.GetPID("13.abc")
			Ω(err).ShouldNot(BeNil())
			Ω(pid).Should(BeNumerically("==", -1))
		})

		It("should error when task is not found", func() {
			store := NewJobStore()
			Ω(store).ShouldNot(BeNil())
			store.SaveArrayJob("13",
				[]int{77, 78, 79},
				drmaa2interface.JobTemplate{RemoteCommand: "rc"},
				1, 3, 1)
			pid, err := store.GetPID("13.10")
			Ω(err).ShouldNot(BeNil())
			Ω(pid).Should(BeNumerically("==", -1))
			pid, err = store.GetPID("13.abc")
			Ω(err).ShouldNot(BeNil())
			Ω(pid).Should(BeNumerically("==", -1))
		})
	})

})