package pipelines_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("A job with nested volume mounts", func() {
	BeforeEach(func() {
		flyHelper.ConfigurePipeline(
			pipelineName,
			"-c", "fixtures/volume-mounting.yml",
		)
	})

	It("procceds through the plan with input under output mounts", func() {
		watch := flyHelper.TriggerJob(pipelineName, "input-under-output")
		<-watch.Exited
		Expect(watch).To(gexec.Exit(0))
		Expect(watch).To(gbytes.Say("some-resource"))
	})

	It("procceds through the plan with input under input mounts", func() {
		watch := flyHelper.TriggerJob(pipelineName, "input-under-input")
		<-watch.Exited
		Expect(watch).To(gexec.Exit(0))
		Expect(watch).To(gbytes.Say("helloworld"))
	})

	It("procceds through the plan with output under input mounts", func() {
		watch := flyHelper.TriggerJob(pipelineName, "output-under-input")
		<-watch.Exited
		Expect(watch).To(gexec.Exit(0))
		Expect(watch).To(gbytes.Say("hello"))
	})

	It("procceds through the plan with input same as output mounts", func() {
		watch := flyHelper.TriggerJob(pipelineName, "input-same-output")
		<-watch.Exited
		Expect(watch).To(gexec.Exit(0))
		Expect(watch).To(gbytes.Say("hello"))
	})
})
