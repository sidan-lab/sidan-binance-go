package utils

import "fmt"

// CheckRequiredParameter validates that a required parameter is not empty
func CheckRequiredParameter(value interface{}, name string) error {
	switch v := value.(type) {
	case string:
		if v == "" {
			return fmt.Errorf("required parameter %s is empty", name)
		}
	case int:
		// For integers, we don't check for zero as it might be valid
		return nil
	case float64:
		// For floats, we don't check for zero as it might be valid
		return nil
	default:
		if v == nil {
			return fmt.Errorf("required parameter %s is nil", name)
		}
	}
	return nil
}

// CheckRequiredParameters validates multiple required parameters
func CheckRequiredParameters(params map[string]interface{}) error {
	for name, value := range params {
		if err := CheckRequiredParameter(value, name); err != nil {
			return err
		}
	}
	return nil
}
