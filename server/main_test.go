package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(handler)
	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned incorrect status code: got %v, expected %v", status, http.StatusOK)
	}
	expected := "Hello World!"
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned incorrect body: got %v, expected %v", actual, expected)
	}
}

func TestWebsocketHandler(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(websocketHandler))
	defer s.Close()
	u := "ws" + strings.TrimPrefix(s.URL, "http")
	ws, resp, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 3; i++ {
		if err := ws.WriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
			t.Fatal(err)
		}
		_, p, err := ws.ReadMessage()
		if err != nil {
			t.Fatal(err)
		}
		if actual, expected := resp.StatusCode, http.StatusSwitchingProtocols; actual != expected {
			t.Errorf("handler returned incorrect body: got %v, expected %v", actual, expected)
		}
		if actual, expected := string(p), "hello"; actual != expected {
			t.Errorf("handler returned incorrect message: got %v, expected %v", actual, expected)
		}
	}
}

func TestRouter(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)
	resp, err := http.Get(mockServer.URL + "/hello")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("incorrect status code: got %v, expected %v", resp.StatusCode, http.StatusOK)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	actual := string(b)
	expected := "Hello World!"
	if actual != expected {
		t.Errorf("incorrect response: got %v, expected %v", actual, expected)
	}
}

func TestRouterWithNonExistentRoute(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)
	resp, err := http.Post(mockServer.URL+"/hello", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("incorrect status code: got %v, expected %v", resp.StatusCode, http.StatusMethodNotAllowed)
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	actual := string(b)
	expected := ""
	if actual != expected {
		t.Errorf("incorrect response: got %v, expected %v", actual, expected)
	}
}
