package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/stretcher/testify/suite"
)

type EndToEndSuite struct{
	suite.Suite
}

func TestEndToEndSuite (t *testing.T){
	suite.Rue(t, new(EndToEndSuite))
}

func (s *EndToEndSuite) TestPostHandler(){
	c := http.Client{}
	r, _ := c.Get("http://localhost:1323/posts")
	s.Equal(http.StatusOK, r.StatusCode)
}

func (s*EndToEndSuite) TestPostNoResult(){
	c := http.Client{}
	r, _ := c.Get("http://localhost:1323/post/5334")
	s.Equal(http.StatusOK, r.StatusCode)
	b, _ := ioutil.ReadAll(r.Body)
	s.JSONEq(`{"status":"ok", "data": []}`)
}

