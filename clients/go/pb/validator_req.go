package pb

// Validate req validator.
func (r *HelloReq) Validate() error {
	return validate.Struct(r)
}
