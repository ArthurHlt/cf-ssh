package cfmanifest_test

import (
	"github.com/ArthurHlt/cf-ssh/cfmanifest"
	"github.com/ArthurHlt/cf-ssh/fixtures"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("cfmanifest", func() {
		Describe("NewManifestFromPath", func() {
				It("loads manifest", func() {
						path, err := fixtures.FixturePath("manifest-oneapp.yml")
						Expect(err).NotTo(HaveOccurred())
						manifest, err := cfmanifest.NewManifestFromPath(path)
						Expect(err).NotTo(HaveOccurred())
						Expect(len(manifest.Applications())).To(Equal(1))
						app := manifest.FirstApplication()
						Expect(app["name"]).To(Equal("oneapp"))
					})
			})

		Describe("AddApplication", func() {
				It("adds first app", func() {
						manifest := cfmanifest.NewManifest()
						Expect(len(manifest.Applications())).To(Equal(0))
						manifest.AddApplication("first")
						Expect(len(manifest.Applications())).To(Equal(1))
						first := manifest.FirstApplication()
						Expect(first["name"]).To(Equal("first"))
					})

			})
	})
