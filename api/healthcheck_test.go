// podium
// https://github.com/topfreegames/podium
// Licensed under the MIT license:
// http://www.opensource.org/licenses/mit-license
// Copyright © 2016 Top Free Games <backend@tfgco.com>
// Forked from
// https://github.com/dayvson/go-leaderboard
// Copyright © 2013 Maxwell Dayvson da Silva

package api_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/topfreegames/podium/api"
)

var _ = Describe("Healthcheck Handler", func() {
	It("Should respond with default WORKING string", func() {
		a := api.GetDefaultTestApp()
		res := api.Get(a, "/healthcheck")

		Expect(res.Raw().StatusCode).To(Equal(http.StatusOK))
		Expect(res.Body().Raw()).To(Equal("WORKING"))
	})

	It("Should respond with customized WORKING string", func() {
		a := api.GetDefaultTestApp()

		a.Config.Set("healthcheck.workingText", "OTHERWORKING")
		res := api.Get(a, "/healthcheck")

		Expect(res.Raw().StatusCode).To(Equal(http.StatusOK))
		Expect(res.Body().Raw()).To(Equal("OTHERWORKING"))
	})

	It("Should fail if redis failing", func() {
		a := api.GetDefaultTestApp()
		a.RedisClient = api.GetFaultyRedis(a.Logger)

		res := api.Get(a, "/healthcheck")

		Expect(res.Raw().StatusCode).To(Equal(500))
		Expect(res.Body().Raw()).To(ContainSubstring("connection refused"))
	})
})
