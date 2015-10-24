package sqlengine_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/cf-platform-eng/rds-broker/sqlengine"

	"github.com/pivotal-golang/lager"
)

var _ = Describe("Provider Service", func() {
	var (
		sqlProvider *ProviderService
		logger      lager.Logger
	)

	BeforeEach(func() {
		logger = lager.NewLogger("sql_engine_test")
		sqlProvider = NewProviderService(logger)
	})

	Describe("GetDatabase", func() {
		It("returns error if engine is not supported", func() {
			_, err := sqlProvider.GetSQLEngine("unknown")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("SQL Engine 'unknown' not supported"))
		})

		Context("when engine is aurora", func() {
			It("return the proper Database", func() {
				sqlEngine, err := sqlProvider.GetSQLEngine("aurora")
				Expect(err).ToNot(HaveOccurred())
				Expect(sqlEngine).To(BeAssignableToTypeOf(&MySQLEngine{}))
			})
		})

		Context("when engine is mariadb", func() {
			It("return the proper Database", func() {
				sqlEngine, err := sqlProvider.GetSQLEngine("mariadb")
				Expect(err).ToNot(HaveOccurred())
				Expect(sqlEngine).To(BeAssignableToTypeOf(&MySQLEngine{}))
			})
		})

		Context("when engine is mysql", func() {
			It("return the proper Database", func() {
				sqlEngine, err := sqlProvider.GetSQLEngine("mysql")
				Expect(err).ToNot(HaveOccurred())
				Expect(sqlEngine).To(BeAssignableToTypeOf(&MySQLEngine{}))
			})
		})

		Context("when engine is postgres", func() {
			It("return the proper Database", func() {
				sqlEngine, err := sqlProvider.GetSQLEngine("postgres")
				Expect(err).ToNot(HaveOccurred())
				Expect(sqlEngine).To(BeAssignableToTypeOf(&PostgresEngine{}))
			})
		})

		Context("when engine is postgresql", func() {
			It("return the proper Database", func() {
				sqlEngine, err := sqlProvider.GetSQLEngine("postgresql")
				Expect(err).ToNot(HaveOccurred())
				Expect(sqlEngine).To(BeAssignableToTypeOf(&PostgresEngine{}))
			})
		})
	})
})