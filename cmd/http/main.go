package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/adityalstkp/flatbuffz/flatbuffer_gen/users"
	flatbuffers "github.com/google/flatbuffers/go"
)

func getParamUsersCount(r *http.Request) (uint64, error) {
	cParam := r.URL.Query().Get("users_count")
	if cParam == "" {
		return 10, nil
	}

	s, err := strconv.ParseUint(cParam, 10, 64)
	if err != nil {
		return 0, err
	}

	return uint64(s), nil
}

type GeoObject struct {
	Lat  string `json:"lat"`
	Long string `json:"long"`
}

type AddressObject struct {
	Street string    `json:"street"`
	City   string    `json:"city"`
	Geo    GeoObject `json:"geo"`
}

type UserObject struct {
	Name    string          `json:"name"`
	Email   string          `json:"email"`
	Phone   string          `json:"phone"`
	Address []AddressObject `json:"address"`
}

func createUserBuff(u UserObject, addr flatbuffers.UOffsetT, b *flatbuffers.Builder) flatbuffers.UOffsetT {
	// build user
	name := b.CreateString(u.Name)
	email := b.CreateString(u.Email)
	phone := b.CreateString(u.Phone)
	users.UserStart(b)
	users.UserAddName(b, name)
	users.UserAddEmail(b, email)
	users.UserAddPhone(b, phone)
	users.UserAddAddress(b, addr)
	return users.UserEnd(b)
}

func createAddressBuff(a AddressObject, b *flatbuffers.Builder) flatbuffers.UOffsetT {
	// build geo
	lat := b.CreateString(a.Geo.Lat)
	long := b.CreateString(a.Geo.Long)
	users.GeoStart(b)
	users.GeoAddLat(b, lat)
	users.GeoAddLong(b, long)
	geo := users.GeoEnd(b)

	// build address
	street := b.CreateString(a.Street)
	city := b.CreateString(a.City)
	users.AddressStart(b)
	users.AddressAddStreet(b, street)
	users.AddressAddCity(b, city)
	users.AddressAddGeo(b, geo)
	address := users.AddressEnd(b)
	return address
}

type httpHandler struct{}

func (h httpHandler) flatbuffHandler(w http.ResponseWriter, r *http.Request) {
	userCount, err := getParamUsersCount(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	log.Println("populating flatbuff users", userCount)

	b := flatbuffers.NewBuilder(0)

	var addressBuffOffset []flatbuffers.UOffsetT
	for i := 0; i < 3; i++ {
		addressBuffOffset = append(addressBuffOffset, createAddressBuff(AddressObject{
			Street: "street",
			City:   "city",
			Geo: GeoObject{
				Lat:  "lat",
				Long: "long",
			},
		}, b))
	}
	users.UserStartAddressVector(b, 3)
	for i := len(addressBuffOffset) - 1; i >= 0; i-- {
		b.PrependUOffsetT(addressBuffOffset[i])
	}
	addressBuff := b.EndVector(3)

	var usersBuffOffset []flatbuffers.UOffsetT
	for i := 0; i < int(userCount); i++ {
		usersBuffOffset = append(usersBuffOffset, createUserBuff(UserObject{
			Name:  fmt.Sprintf("name-%d", i+1),
			Email: fmt.Sprintf("email-%d", i+1),
			Phone: fmt.Sprintf("phone-%d", i+1),
		}, addressBuff, b))
	}

	users.UsersStartListVector(b, int(userCount))
	for i := len(usersBuffOffset) - 1; i >= 0; i-- {
		b.PrependUOffsetT(usersBuffOffset[i])
	}
	usersBuff := b.EndVector(int(userCount))

	users.UsersStart(b)
	users.UsersAddList(b, usersBuff)
	usersList := users.UsersEnd(b)
	b.Finish(usersList)

	buf := b.FinishedBytes()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")

	w.Write(buf)
}

func (h httpHandler) jsonHandler(w http.ResponseWriter, r *http.Request) {
	userCount, err := getParamUsersCount(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	log.Println("populating json users", userCount)

	var addrs []AddressObject
	for i := 0; i < 3; i++ {
		addrs = append(addrs, AddressObject{
			Street: "street",
			City:   "city",
			Geo: GeoObject{
				Lat:  "lat",
				Long: "long",
			},
		})
	}
	var users []UserObject
	for i := 0; i < int(userCount); i++ {
		users = append(users, UserObject{
			Name:    fmt.Sprintf("name-%d", i+1),
			Email:   fmt.Sprintf("email-%d", i+1),
			Phone:   fmt.Sprintf("phone-%d", i+1),
			Address: addrs,
		})
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h httpHandler) msgPackHandler(w http.ResponseWriter, r *http.Request) {
	_, err := getParamUsersCount(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("not yet implemented"))
}

func (h httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/flatbuff" {
		h.flatbuffHandler(w, r)
		return
	}

	if r.URL.Path == "/json" {
		h.jsonHandler(w, r)
		return
	}

	if r.URL.Path == "/msgpack" {
		h.msgPackHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 not found"))
}

func main() {
	handler := httpHandler{}
	s := &http.Server{
		Addr:    ":8000",
		Handler: handler,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Panicln(err)
	}
}
