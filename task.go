package utils

type TaskConfig map[string]string

func (tc TaskConfig) GetString(key string) string {
	value, ok := tc[key]
	if !ok {
		return ""
	}

	return value
}
