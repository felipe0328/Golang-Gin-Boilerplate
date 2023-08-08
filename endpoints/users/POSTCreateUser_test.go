package users

import (
	"bytes"
	"encoding/json"
	"errors"
	"gin-microservice/controller/usersController/mocks"
	usersmodels "gin-microservice/models/usersModels"
	"gin-microservice/utils"
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
	endpoint  UsersEndpoints
}

func TestPOSTCreateUser(t *testing.T) {
	suite.Run(t, new(POSTCreateUserTestSuite))
}

func (t *POSTCreateUserTestSuite) SetupTest() {
	t.contoller = new(mocks.IUsersController)
	t.endpoint = UsersEndpoints{
		Controller: t.contoller,
	}
}

func (t *POSTCreateUserTestSuite) Test_CreateUser_Success(){
	userData := usersmodels.User{
		Username: "Felipe",
		Password: "12345",
		Email: "aw@gmail.com",
		FirstName: "Felipe",
		LastName: "Puentes",
	}

	writer, c := createWriterWithUserPostBody(userData)

	var timeNow time.Time
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

func (t *POSTCreateUserTestSuite) Test_CreateUser_NoValidEntry_Fail(){
	inputObject := make(map[string]string)
	inputObject["Data"] = "Value"

	writer, c := createWriterWithUserPostBody(inputObject)

	t.endpoint.CreateUser(c)

	t.Equal(http.StatusBadRequest, writer.Code)
	t.Equal(utils.ErrInvalidUserObject.Error(), writer.Body.String())
	t.contoller.AssertNumberOfCalls(t.T(), "CreateUser", 0)
}

func (t *POSTCreateUserTestSuite) Test_CreateUser_ControllerError_Fail(){
	userData := usersmodels.User{
		Username: "Felipe",
		Password: "12345",
		Email: "aw@gmail.com",
		FirstName: "Felipe",
		LastName: "Puentes",
	}

	writer, c := createWriterWithUserPostBody(userData)

	expectedError := errors.New("expected error")
	t.contoller.On("CreateUser", userData).Return(usersmodels.UserObject{}, expectedError)

	t.endpoint.CreateUser(c)

	wrappedError := utils.StringWrapError(utils.ErrUnableToCreateUser, expectedError)

	t.Equal(http.StatusBadRequest, writer.Code)
	t.Equal(wrappedError, writer.Body.String())
	t.contoller.AssertNumberOfCalls(t.T(), "CreateUser", 1)
}

func createWriterWithUserPostBody(input any)(*httptest.ResponseRecorder, *gin.Context){
	writer := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(writer)

	bodyData, _ := json.Marshal(input)
	bytesBuffer := bytes.NewBuffer(bodyData)
	request := httptest.NewRequest(http.MethodPost, "/users", bytesBuffer)
	c.Request = request

	return writer, c
}