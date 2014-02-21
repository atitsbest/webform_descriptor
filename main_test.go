package main

import (
    . "github.com/smartystreets/goconvey/convey"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

func Test_Server(t *testing.T) {
    Convey("Subject: Server", t, func() {
        Convey("A running server", func() {
            sut := initMartini()
            go http.ListenAndServe(":3001", sut)

            Convey("should GET /", func() { shouldServe(t, "http://localhost:3001/") })
        })
    })
}

func Test_Project_Api(t *testing.T) {
    Convey("Subject: Project API", t, func() {
        Convey("A running server", func() {
            sut := initMartini()
            go http.ListenAndServe(":3001", sut)

            Convey("should GET /api/projects", func() { shouldServe(t, "http://localhost:3001/api/projects") })
            Convey("should GET /api/projects/12345", func() { shouldServe(t, "http://localhost:3001/api/projects/12345") })
            Convey("should POST /api/projects", func() {
                response := httptest.NewRecorder()
                json := `{
                      "name":"Eriksson",
                      "leader":"RatAn",
                      "risk":"A",
                      "accountingMode":"Fixpreis",
                      "state":"beauftragt",
                      "orderDate":"2014-02-19T07:48:33.833Z",
                      "techs":["C#","F#","JavaScript","SharePoint","Progress"],
                      "customer":"sldslkdh",
                      "orderAmount":33.99,
                      "orderAmountDays":44
                    }`
                b := strings.NewReader(json)
                req, err := http.NewRequest("POST", "http://localhost:3000/api/projects", b)
                if err != nil {
                    t.Error(err)
                }

                // Act
                sut.ServeHTTP(response, req)

                // Assert
                So(response.Code, ShouldEqual, http.StatusCreated)
                So(response.Body.Len(), ShouldBeGreaterThan, 0)
            })
        })
    })
}

// ========================================================
func shouldServe(t *testing.T, url string) {
    // Arrange
    sut := initMartini()
    response := httptest.NewRecorder()
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        t.Error(err)
    }

    // Act
    sut.ServeHTTP(response, req)

    // Assert
    So(response.Code, ShouldEqual, http.StatusOK)
    So(response.Body.Len(), ShouldBeGreaterThan, 0)
}
