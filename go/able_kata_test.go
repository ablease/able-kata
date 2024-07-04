package able_kata_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	able_kata "github.com/ablease/able-kata"
)

var _ = Describe("AbleKata", func() {
	var foxInSocks, lesMis *able_kata.Book

	BeforeEach(func() {
		lesMis = &able_kata.Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  2783,
		}

		foxInSocks = &able_kata.Book{
			Title:  "Fox In Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}
	})

	Describe("Categorizing able_kata", func() {
		Context("with more than 300 pages", func() {
			It("should be a novel", func() {
				Expect(lesMis.Category()).To(Equal(lesMis.CategoryNovel()))
			})
		})

		Context("with fewer than 300 pages", func() {
			It("should be a short story", func() {
				Expect(foxInSocks.Category()).To(Equal(foxInSocks.CategoryShortStory()))
			})
		})
	})
})
