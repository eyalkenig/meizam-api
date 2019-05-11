package command

import (
	"github.com/eyalkenig/meizam-api/api/app/repository"
	"github.com/eyalkenig/meizam-api/api/app/repository/mock_repository"
	"github.com/eyalkenig/meizam-api/api/app/service"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Command", func() {
	var (
		mockCtrl   *gomock.Controller
		repository repository.Repository
		command    service.Command
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		repository = mock_repository.NewMockRepository(mockCtrl)
		command = NewMeizamCommand(repository)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("CreateTeam", func() {
		It("Should return error on empty team name", func() {
			_, err := command.CreateTeam("", nil, nil)

			Expect(err).To(Not(Equal(nil)))
		})
	})
})
