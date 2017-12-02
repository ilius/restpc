package ripo

import (
	"fmt"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/golang/mock/gomock"
)

func TestFromBody_GetString(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	{
		mockReq.EXPECT().BodyMap().Return(nil, fmt.Errorf("unknown error"))
		value, err := FromBody.GetString(req, "name")
		assert.EqualError(t, err, "unknown error")
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(nil, nil)
		value, err := FromBody.GetString(req, "name")
		assert.NoError(t, err)
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"name": 123,
		}, nil)
		value, err := FromBody.GetString(req, "name")
		assert.EqualError(t, err, "invalid 'name', must be string")
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"name": "John Smith",
		}, nil)
		value, err := FromBody.GetString(req, "name")
		assert.NoError(t, err)
		assert.Equal(t, "John Smith", *value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"name": []byte("John Smith"),
		}, nil)
		value, err := FromBody.GetString(req, "name")
		assert.NoError(t, err)
		assert.Equal(t, "John Smith", *value)
	}
}
func TestFromBody_GetStringList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	{
		mockReq.EXPECT().BodyMap().Return(nil, fmt.Errorf("unknown error"))
		value, err := FromBody.GetStringList(req, "names")
		assert.EqualError(t, err, "unknown error")
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(nil, nil)
		value, err := FromBody.GetStringList(req, "names")
		assert.NoError(t, err)
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"names": 123,
		}, nil)
		value, err := FromBody.GetStringList(req, "names")
		assert.EqualError(t, err, "invalid 'names', must be array of strings")
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"names": "John Smith",
		}, nil)
		value, err := FromBody.GetStringList(req, "names")
		assert.EqualError(t, err, "invalid 'names', must be array of strings")
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"names": []interface{}{"John Smith", 1234},
		}, nil)
		value, err := FromBody.GetStringList(req, "names")
		assert.EqualError(t, err, "invalid 'names', must be array of strings")
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"names": []string{"John Smith", "John Doe"},
		}, nil)
		value, err := FromBody.GetStringList(req, "names")
		assert.NoError(t, err)
		assert.Equal(t, []string{"John Smith", "John Doe"}, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"names": []interface{}{"John Smith", "John Doe"},
		}, nil)
		value, err := FromBody.GetStringList(req, "names")
		assert.NoError(t, err)
		assert.Equal(t, []string{"John Smith", "John Doe"}, value)
	}

}

func TestFromBody_GetInt(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	{
		mockReq.EXPECT().BodyMap().Return(nil, fmt.Errorf("unknown error"))
		value, err := FromBody.GetInt(req, "count")
		assert.EqualError(t, err, "unknown error")
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(nil, nil)
		value, err := FromBody.GetInt(req, "count")
		assert.NoError(t, err)
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"count": "abc",
		}, nil)
		value, err := FromBody.GetInt(req, "count")
		assert.EqualError(t, err, "invalid 'count', must be integer")
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"count": "345",
		}, nil)
		value, err := FromBody.GetInt(req, "count")
		assert.EqualError(t, err, "invalid 'count', must be integer")
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"count": 5001,
		}, nil)
		value, err := FromBody.GetInt(req, "count")
		assert.NoError(t, err)
		assert.Equal(t, 5001, *value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"count": int32(5003),
		}, nil)
		value, err := FromBody.GetInt(req, "count")
		assert.NoError(t, err)
		assert.Equal(t, 5003, *value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"count": int64(6123),
		}, nil)
		value, err := FromBody.GetInt(req, "count")
		assert.NoError(t, err)
		assert.Equal(t, 6123, *value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"count": 14.15,
		}, nil)
		value, err := FromBody.GetInt(req, "count")
		assert.NoError(t, err)
		assert.Equal(t, 14, *value)
	}
}

func TestFromBody_GetFloat(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	{
		mockReq.EXPECT().BodyMap().Return(nil, fmt.Errorf("unknown error"))
		value, err := FromBody.GetFloat(req, "weight")
		assert.Nil(t, value)
		assert.EqualError(t, err, "unknown error")
	}
	{
		mockReq.EXPECT().BodyMap().Return(nil, nil)
		value, err := FromBody.GetFloat(req, "weight")
		assert.Nil(t, value)
		assert.NoError(t, err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"weight": "abc",
		}, nil)
		value, err := FromBody.GetFloat(req, "weight")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'weight', must be float")
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"weight": "345",
		}, nil)
		value, err := FromBody.GetFloat(req, "weight")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'weight', must be float")
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"weight": 1231,
		}, nil)
		value, err := FromBody.GetFloat(req, "weight")
		assert.Equal(t, 1231.0, *value)
		assert.NoError(t, err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"weight": int32(2345),
		}, nil)
		value, err := FromBody.GetFloat(req, "weight")
		assert.Equal(t, 2345.0, *value)
		assert.NoError(t, err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"weight": int64(7123),
		}, nil)
		value, err := FromBody.GetFloat(req, "weight")
		assert.Equal(t, 7123.0, *value)
		assert.NoError(t, err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"weight": 104.15,
		}, nil)
		value, err := FromBody.GetFloat(req, "weight")
		assert.Equal(t, 104.15, *value)
		assert.NoError(t, err)
	}
	{
		weight := float32(104.15)
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"weight": weight,
		}, nil)
		value, err := FromBody.GetFloat(req, "weight")
		assert.Equal(t, float64(weight), *value)
		assert.NoError(t, err)
	}
}

func TestFromBody_GetBool(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	{
		mockReq.EXPECT().BodyMap().Return(nil, fmt.Errorf("unknown error"))
		value, err := FromBody.GetBool(req, "agree")
		assert.Nil(t, value)
		assert.EqualError(t, err, "unknown error")
	}
	{
		mockReq.EXPECT().BodyMap().Return(nil, nil)
		value, err := FromBody.GetBool(req, "agree")
		assert.Nil(t, value)
		assert.NoError(t, err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"agree": "abcd",
		}, nil)
		value, err := FromBody.GetBool(req, "agree")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'agree', must be true or false")
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"agree": "3465",
		}, nil)
		value, err := FromBody.GetBool(req, "agree")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'agree', must be true or false")
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"agree": 1231,
		}, nil)
		value, err := FromBody.GetBool(req, "agree")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'agree', must be true or false")
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"agree": true,
		}, nil)
		value, err := FromBody.GetBool(req, "agree")
		assert.Equal(t, true, *value)
		assert.NoError(t, err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"agree": false,
		}, nil)
		value, err := FromBody.GetBool(req, "agree")
		assert.Equal(t, false, *value)
		assert.NoError(t, err)
	}
}

func TestFromBody_GetTime(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	{
		mockReq.EXPECT().BodyMap().Return(nil, fmt.Errorf("unknown error"))
		value, err := FromBody.GetTime(req, "since")
		assert.Nil(t, value)
		assert.EqualError(t, err, "unknown error")
	}
	{
		mockReq.EXPECT().BodyMap().Return(nil, nil)
		value, err := FromBody.GetTime(req, "since")
		assert.Nil(t, value)
		assert.NoError(t, err)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"since": "abcd",
		}, nil)
		value, err := FromBody.GetTime(req, "since")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'since', must be RFC3339 time string")
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"since": 3465,
		}, nil)
		value, err := FromBody.GetTime(req, "since")
		assert.Nil(t, value)
		assert.EqualError(t, err, "invalid 'since', must be RFC3339 time string")
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"since": "2017-12-20T17:30:00Z",
		}, nil)
		value, err := FromBody.GetTime(req, "since")
		assert.NoError(t, err)
		assert.Equal(t, time.Date(2017, time.Month(12), 20, 17, 30, 0, 0, time.UTC), *value)
	}
}

func TestFromBody_GetObject(t *testing.T) {
	type Person struct {
		Name      string  `json:"name"`
		BirthDate []int   `json:"birthDate"` // mapstructure does not support [3]int
		Age       float64 `json:"age"`
	}
	PersonType := reflect.TypeOf(Person{})
	PersonTypePtr := reflect.TypeOf(&Person{})
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockReq := NewMockExtendedRequest(ctrl)
	var req ExtendedRequest = mockReq
	{
		mockReq.EXPECT().BodyMap().Return(nil, fmt.Errorf("unknown error"))
		value, err := FromBody.GetObject(req, "info", PersonTypePtr)
		assert.EqualError(t, err, "unknown error")
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(nil, nil)
		value, err := FromBody.GetObject(req, "info", PersonTypePtr)
		assert.NoError(t, err)
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"info": 123,
		}, nil)
		value, err := FromBody.GetObject(req, "info", PersonTypePtr)
		assert.EqualError(t, err, "invalid 'info', must be a compatible object")
		assert.Nil(t, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"info": map[string]interface{}{},
		}, nil)
		value, err := FromBody.GetObject(req, "info", PersonTypePtr)
		assert.NoError(t, err)
		if err != nil {
			log.Println("Private:", err.(RPCError).Private())
		}
		assert.Equal(t, &Person{}, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"info": map[string]interface{}{},
		}, nil)
		value, err := FromBody.GetObject(req, "info", PersonType)
		assert.NoError(t, err)
		if err != nil {
			log.Println("Private:", err.(RPCError).Private())
		}
		assert.Equal(t, Person{}, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"info": map[string]interface{}{
				"name": "John Smith",
			},
		}, nil)
		value, err := FromBody.GetObject(req, "info", PersonTypePtr)
		assert.NoError(t, err)
		if err != nil {
			log.Println("Private:", err.(RPCError).Private())
		}
		assert.Equal(t, &Person{
			Name: "John Smith",
		}, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"info": map[string]interface{}{
				"name":      "John Smith",
				"birthDate": []int{1987, 12, 30},
			},
		}, nil)
		value, err := FromBody.GetObject(req, "info", PersonTypePtr)
		assert.NoError(t, err)
		if err != nil {
			log.Println("Private:", err.(RPCError).Private())
		}
		assert.Equal(t, &Person{
			Name:      "John Smith",
			BirthDate: []int{1987, 12, 30},
		}, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"info": map[string]interface{}{
				"name":      "John Smith",
				"birthDate": []int{1987, 12, 30},
				"age":       30.8,
			},
		}, nil)
		value, err := FromBody.GetObject(req, "info", PersonTypePtr)
		assert.NoError(t, err)
		if err != nil {
			log.Println("Private:", err.(RPCError).Private())
		}
		assert.Equal(t, &Person{
			Name:      "John Smith",
			BirthDate: []int{1987, 12, 30},
			Age:       30.8,
		}, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"info": map[string]interface{}{
				"name":      "John Smith",
				"birthDate": []int{1987, 12, 30},
				"age":       30.8,
			},
		}, nil)
		value, err := FromBody.GetObject(req, "info", PersonType)
		assert.NoError(t, err)
		if err != nil {
			log.Println("Private:", err.(RPCError).Private())
		}
		assert.Equal(t, Person{
			Name:      "John Smith",
			BirthDate: []int{1987, 12, 30},
			Age:       30.8,
		}, value)
	}
	{
		mockReq.EXPECT().BodyMap().Return(map[string]interface{}{
			"guestList": []interface{}{
				map[string]interface{}{
					"name":      "John Smith",
					"birthDate": []int{1987, 12, 30},
					"age":       30.8,
				},
				map[string]interface{}{
					"name":      "Jane Smith",
					"birthDate": []int{1991, 6, 30},
					"age":       27.3,
				},
			},
		}, nil)
		value, err := FromBody.GetObject(req, "guestList", reflect.SliceOf(PersonType))
		assert.NoError(t, err)
		if err != nil {
			log.Println("Private:", err.(RPCError).Private())
		}
		assert.Equal(t, []Person{
			{
				Name:      "John Smith",
				BirthDate: []int{1987, 12, 30},
				Age:       30.8,
			},
			{
				Name:      "Jane Smith",
				BirthDate: []int{1991, 6, 30},
				Age:       27.3,
			},
		}, value)
	}
}
