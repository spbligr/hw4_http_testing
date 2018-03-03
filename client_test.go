package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"encoding/xml"
	"encoding/json"
)

type xmlRow struct {
	Id int `xml:"id"`
	Guid string `xml:"guid"`
	IsActive bool `xml:"isActive"`
	Balance string `xml:"balance"`
	Picture string `xml:"picture"`
	Age int `xml:"age"`
	EyeColor string `xml:"eyeColor"`
	FirstName string `xml:"first_name"`
	LastName string `xml:"last_name"`
	Gender string `xml:"gender"`
	Company string `xml:"company"`
	Email string `xml:"email"`
	Phone string `xml:"phone"`
	Address string `xml:"address"`
	About string `xml:"about"`
}

type xmlStructure struct {
	Version string `xml:"version"`
	Row []xmlRow `xml:"row"`

}


func SearchServer(w http.ResponseWriter, r *http.Request)  {
	dataFile, err := ioutil.ReadFile("dataset.xml")
	checkError(err)

	usersXml := &xmlStructure{}
	xml.Unmarshal(dataFile, &usersXml)

	var users []User

	for _, user := range usersXml.Row {
		users = append(users, User{
			Id: user.Id,
			Name: user.FirstName,
			Age: user.Age,
			About: user.About,
			Gender: user.Gender,
		})
	}


	response := &SearchResponse{
		Users: users,
		NextPage:false,
	}

	jsonResponse, err := json.Marshal(response)
	checkError(err)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}


func TestServer(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(SearchServer))

	searchClient := &SearchClient{
		URL: ts.URL,
	}

	_, err := searchClient.FindUsers(SearchRequest{
		Limit: 5,
		Offset: 0,
		Query: "Name",
		OrderField: "Name",
		OrderBy: 1,

	})

	if err != nil {
		panic(err)
	}

	ts.Close()
}

func checkError(err error)  {
	if err != nil {
		panic(err)
	}
}

func temp()  {
	//response, err := http.Get(ts.URL)
	//checkError(err)
	//
	//body, _ := ioutil.ReadAll(response.Body)
	//
	//fmt.Println(string(body))
	//
	//defer response.Body.Close()

}