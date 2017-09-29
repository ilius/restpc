package restpc

import (
	"fmt"
	"time"
)

var FromBody FromX

func init() {
	FromBody = &fromBody{}
}

type fromBody struct{}

func (f *fromBody) GetString(req Request, key string) (*string, error) {
	data, err := req.BodyMap()
	if err != nil {
		return nil, err
	}
	valueIn := data[key]
	if valueIn != nil {
		switch value := valueIn.(type) {
		case string:
			valueStr := value // to copy
			return &valueStr, nil
		case []byte:
			valueStr := string(value)
			return &valueStr, nil
		default:
			return nil, NewError(
				InvalidArgument,
				fmt.Sprintf("invalid '%v', must be string", key),
				nil,
				"value", value,
			)
		}
	}
	return nil, nil
}
func (f *fromBody) GetStringList(req Request, key string) ([]string, error) {
	data, err := req.BodyMap()
	if err != nil {
		return nil, err
	}
	valueIn := data[key]
	if valueIn != nil {
		switch value := valueIn.(type) {
		case []string:
			valueSlice := append([]string(nil), value...) // to copy
			return valueSlice, nil
		default:
			return nil, NewError(
				InvalidArgument,
				fmt.Sprintf("invalid '%v', must be array of strings", key),
				nil,
				"value", value,
			)
		}
	}
	return nil, nil
}
func (f *fromBody) GetInt(req Request, key string) (*int, error) {
	data, err := req.BodyMap()
	if err != nil {
		return nil, err
	}
	valueIn := data[key]
	if valueIn != nil {
		switch value := valueIn.(type) {
		case float64:
			valueInt := int(value)
			return &valueInt, nil
		case int:
			valueInt := value // to copy
			return &valueInt, nil
		case int32:
			valueInt := int(value)
			return &valueInt, nil
		case int64:
			valueInt := int(value)
			return &valueInt, nil
		default:
			return nil, NewError(
				InvalidArgument,
				fmt.Sprintf("invalid '%v', must be integer", key),
				nil,
				"value", value,
			)
		}
	}
	return nil, nil
}
func (f *fromBody) GetFloat(req Request, key string) (*float64, error) {
	data, err := req.BodyMap()
	if err != nil {
		return nil, err
	}
	valueIn := data[key]
	if valueIn != nil {
		switch value := valueIn.(type) {
		case float64:
			valueF := value // to copy
			return &valueF, nil
		case float32:
			valueF := float64(value)
			return &valueF, nil
		case int:
			valueF := float64(value)
			return &valueF, nil
		case int64:
			valueF := float64(value)
			return &valueF, nil
		case int32:
			valueF := float64(value)
			return &valueF, nil
		default:
			return nil, NewError(
				InvalidArgument,
				fmt.Sprintf("invalid '%v', must be float", key),
				nil,
				"value", value,
			)
		}
	}
	return nil, nil
}
func (f *fromBody) GetBool(req Request, key string) (*bool, error) {
	data, err := req.BodyMap()
	if err != nil {
		return nil, err
	}
	valueIn := data[key]
	if valueIn != nil {
		switch value := valueIn.(type) {
		case bool:
			valueBool := value // to copy
			return &valueBool, nil
		default:
			return nil, NewError(
				InvalidArgument,
				fmt.Sprintf("invalid '%v', must be true or false", key),
				nil,
				"value", value,
			)
		}
	}
	return nil, nil
}
func (f *fromBody) GetTime(req Request, key string) (*time.Time, error) {
	data, err := req.BodyMap()
	if err != nil {
		return nil, err
	}
	valueIn := data[key]
	if valueIn != nil {
		switch value := valueIn.(type) {
		case string:
			valueTm, err := time.Parse(time.RFC3339, value)
			if err != nil {
				return nil, NewError(
					InvalidArgument,
					fmt.Sprintf("invalid '%v', must be RFC3339 time string", key),
					err,
					"value", value,
				)
			}
			return &valueTm, nil
		default:
			return nil, NewError(
				InvalidArgument,
				fmt.Sprintf("invalid '%v', must be RFC3339 time string", key),
				nil,
				"value", value,
			)
		}
	}
	return nil, nil
}