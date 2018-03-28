package nats_streaming_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fmt"

	"github.com/solo-io/gloo-api/pkg/api/types/v1"
	. "github.com/solo-io/gloo-function-discovery/internal/nats-streaming"
	"github.com/solo-io/gloo-plugins/nats-streaming"
)

var _ = Describe("DiscoverNats", func() {
	Describe("happy path", func() {
		Context("upstream for a nats-streaming server", func() {
			It("returns service info for nats-streaming", func() {
				err = natsStreamingInstance.Run()
				Expect(err).NotTo(HaveOccurred())
				detector := NewNatsDetector(natsStreamingInstance.ClusterId())
				Expect(detector.DetectsFor(&v1.Upstream{})).To(BeTrue())
				svcInfo, annotations, err := detector.DetectFunctionalService(fmt.Sprintf("localhost:%v", natsStreamingInstance.NatsPort()))
				Expect(err).To(BeNil())
				Expect(annotations).To(BeNil())
				Expect(svcInfo).To(Equal(&v1.ServiceInfo{
					Type: natsstreaming.ServiceTypeNatsStreaming,
					Properties: natsstreaming.EncodeServiceProperties(natsstreaming.ServiceProperties{
						ClusterID: natsStreamingInstance.ClusterId(),
					}),
				}))
			})
		})
	})
})
