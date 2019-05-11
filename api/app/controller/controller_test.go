package controller_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/eyalkenig/meizam-api/api/app/controller"
	"github.com/eyalkenig/meizam-api/api/app/service/mock_service"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Controller", func() {
	var (
		mockCtrl      *gomock.Controller
		service       *mock_service.MockService
		appController *controller.Controller
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		service = mock_service.NewMockService(mockCtrl)
		appController = controller.NewController(service)
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("ListTeams", func() {
		It("Should have default page and per page", func() {
			request, _ := http.NewRequest("GET", "/teams", nil)
			responseRecorder := httptest.NewRecorder()

			service.EXPECT().ListTeams(10, 0)

			appController.ListTeams(responseRecorder, request)
		})
		It("Should use page and per page from url params", func() {
			request, _ := http.NewRequest("GET", "/teams", nil)
			values := url.Values{}
			values.Add("page", "5")
			values.Add("per_page", "20")
			request.URL.RawQuery = values.Encode()
			responseRecorder := httptest.NewRecorder()

			service.EXPECT().ListTeams(20, 100)

			appController.ListTeams(responseRecorder, request)
		})
		It("Should return error for non-positive page", func() {
			request, _ := http.NewRequest("GET", "/teams", nil)
			values := url.Values{}
			values.Add("page", "-1")
			request.URL.RawQuery = values.Encode()
			responseRecorder := httptest.NewRecorder()

			appController.ListTeams(responseRecorder, request)

			Expect(responseRecorder.Code).To(Equal(http.StatusInternalServerError))
		})
		It("Should return error for non-positive per_page", func() {
			request, _ := http.NewRequest("GET", "/teams", nil)
			values := url.Values{}
			values.Add("per_page", "-1")
			request.URL.RawQuery = values.Encode()
			responseRecorder := httptest.NewRecorder()

			appController.ListTeams(responseRecorder, request)

			Expect(responseRecorder.Code).To(Equal(http.StatusInternalServerError))
		})
	})
})
