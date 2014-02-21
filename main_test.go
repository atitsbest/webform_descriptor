package main

import (
  "net/http"
  "net/http/httptest"
  "reflect"
  "strings"
  "testing"
)


func Test_Server_Run(t *testing.T) {
  sut := initMartini()
  go http.ListenAndServe(":3001", sut)
}

func Test_Server_GET_Projects_Index(t *testing.T) {
  shouldServe(t, "http://localhost:3001/")
}

func Test_Server_GET_Api_Projects(t *testing.T) {
  shouldServe(t, "http://localhost:3001/api/projects")
}

func Test_Server_GET_Api_Project(t *testing.T) {
  shouldServe(t, "http://localhost:3001/api/projects/12345")
}

func Test_Server_POST_Api_Projects(t *testing.T) {
  // Arrange
  sut := initMartini()
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
  if (err != nil) { t.Error(err) }

  // Act
  sut.ServeHTTP(response, req)

  // Assert
  expect(t, response.Code, http.StatusCreated)
  if response.Body.Len() == 0 { t.Errorf("Leerer Body beim POST Request!") }
}

// ========================================================
func shouldServe(t *testing.T, url string) {
  // Arrange
  sut := initMartini()
  response := httptest.NewRecorder()
  req, err := http.NewRequest("GET", url, nil)
  if (err != nil) { t.Error(err) }

  // Act
  sut.ServeHTTP(response, req)

  // Assert
  expect(t, response.Code, http.StatusOK)
  if response.Body.Len() == 0 { t.Errorf("Leerer Body beim GET Request!") }
}

// Test-Helper
func expect(t *testing.T, a interface{}, b interface{}) {
  if a != b {
    t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
  }
}
