package api_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Teams", func() {
	var (
		serverUrl   string
		payloadMap  map[string]interface{}
		payloadJson []byte
		request     *http.Request
		client      *http.Client
	)

	BeforeEach(func() {
		serverUrl = fmt.Sprintf("http://%s/v1/teams", getHTTPURI())
		name := "Donatello"
		externalEntityId := "5"
		imageUrl := "/image/url"
		payloadMap = map[string]interface{}{"name": name, "external_entity_id": externalEntityId, "image_url": imageUrl}
		payloadJson = toJSON(payloadMap)
		request, _ = http.NewRequest("POST", serverUrl, bytes.NewBuffer(payloadJson))
		request.Close = true
		request.Header.Set("Content-Type", "application/json")
		client = &http.Client{}
	})

	It("Should return OK status", func() {
		resp, err := client.Do(request)
		if err != nil {
			log.Println("request failed: " + err.Error())
			panic(err)
		}
		defer resp.Body.Close()

		Expect(resp.StatusCode).To(Equal(http.StatusOK))
		Expect(err).To(Succeed())
	})

	It("Should return the id of the created team", func() {
		resp, err := client.Do(request)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		jsonResponse := fromJSON(bodyBytes)
		Expect(jsonResponse["id"]).NotTo(BeZero())
		Expect(err).To(Succeed())
	})
})
