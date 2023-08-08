package tests

import (
	"bytes"
	"encoding/json"
	"gin-microservice/controller/usersController/mocks"
	"gin-microservice/endpoints/users"
	usersmodels "gin-microservice/models/usersModels"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type POSTCreateUserTestSuite struct {
	suite.Suite
	contoller *mocks.IUsersController
	endpoint  users.UsersEndpoints
}

func TestPOSTCreateUser(t *testing.T) {
	suite.Run(t, new(POSTCreateUserTestSuite))
}

func (t *POSTCreateUserTestSuite) SetupTest() {
	t.contoller = new(mocks.IUsersController)
	t.endpoint = users.UsersEndpoints{
		Controller: t.contoller,
	}
}

func (t *POSTCreateUserTestSuite) CreateUser_Success(){
	userData := usersmodels.User{
		Username: "Felipe",
		Password: "12345",
		Email: "aw@gmail.com",
		FirstName: "Felipe",
		LastName: "Puentes",
	}

	writer := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(writer)

	bodyData, _ := json.Marshal(userData)
	bytesBuffer := bytes.NewBuffer(bodyData)
	request := httptest.NewRequest(http.MethodPost, "/users", bytesBuffer)
	c.Request = request

	timeNow := time.Now()
	expectedResponse := usersmodels.UserObject{
		ID: "1",
		Username: "Felipe",
		Email: "a@gma.com",
		FirstName: "Felipe",
		LastName: "Puentes",
		CreatedOn: &timeNow,
		IsActive: true,
	}

	t.contoller.On("CreateUser", userData).Return(expectedResponse, nil)

	t.endpoint.CreateUser(c)

	var response usersmodels.UserObject
	err := json.Unmarshal(writer.Body.Bytes(), &response)

	t.contoller.AssertNumberOfCalls(t.T(), "CreateUser", 1)
	t.Nil(err)
	t.Equal(http.StatusCreated, writer.Code)
	t.Equal(expectedResponse, response)
}
