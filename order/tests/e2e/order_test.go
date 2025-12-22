//go:build integration

package integration

import (
	"context"
	"github.com/Muvi7z/boilerplate/order/internal/repository"
	"github.com/Muvi7z/boilerplate/order/internal/usecase/order"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("OrderService", func() {
	var (
		ctx    context.Context
		cancel context.CancelFunc
		repo   order.Repository
	)

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(suiteCtx)

		repo = repository.New(env.Postgres.Client())

	})

	AfterEach(func() {
		err := env.ClearOrderTable(ctx)

		Expect(err).ToNot(HaveOccurred(), "ожидали успешную очистку таблицу orders")

		cancel()
	})

	Describe("Create", func() {
		It("должен успешно создавать новый заказ", func() {
			order := env.GetTestOrder()

			res, err := repo.Create(ctx, order)

			Expect(err).ToNot(HaveOccurred())
			Expect(res).ToNot(BeEmpty())
			Expect(res).To(MatchRegexp(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`))
		})
	})

})
