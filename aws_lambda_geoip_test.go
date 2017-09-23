package main_test

import (
	geoip "github.com/futurenda/aws-lambda-geoip"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AwsLambdaGeoip", func() {
	It("Should be able to query", func() {
		ips, err := geoip.Query([]string{
            "1.2.3.4",
        })
        Expect(err).To(BeNil())
        Expect(ips[0].IP).To(Equal("1.2.3.4"))
	})
	It("Should ignore invalid ip", func() {
		ips, err := geoip.Query([]string{
            "1.2.3.",
            "",
        })
        Expect(err).To(BeNil())
        Expect(len(ips)).To(Equal(0))
	})
})
