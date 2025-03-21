package validator

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}
func (v *Validator) IsValid() bool {
	return len(v.Errors) == 0
}

func In[T comparable](value T, values []T) bool {
	for _, val := range values {
		if val == value {
			return true
		}
	}
	return false
}
