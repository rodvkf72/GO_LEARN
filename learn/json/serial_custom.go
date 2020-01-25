package json

import "errors"

type status int

const (
	UNKNOWN status = iota
	TODO
	DONE

func (s status) MarshalJSON() ([]byte, error) {
	switch s {
	case UNKNOWN:
		return []byte(`"UNKNOWN"`), nil
	case TODO:
		return []byte(`"TODO"`), nil
	case DONE:
		return []byte(`"DONE"`), nil
	default:
		return nil, errors.New("status.MarshalJSON: unknown value")
	}
}
func (s *status) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"UNKNOWN"`:
		*s = UNKNOWN
	case `"TODO"`:
		*s = TODO
	case `"DONE"`:
		*s = DONE
	default:
		return errors.New("status.UnmarshalJSON: unknown value")
	}
	return nil
}
