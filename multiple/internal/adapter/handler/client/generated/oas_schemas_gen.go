// Code generated by ogen, DO NOT EDIT.

package generated

// Ref: #/components/schemas/Address
type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

// GetStreet returns the value of Street.
func (s *Address) GetStreet() string {
	return s.Street
}

// GetCity returns the value of City.
func (s *Address) GetCity() string {
	return s.City
}

// SetStreet sets the value of Street.
func (s *Address) SetStreet(val string) {
	s.Street = val
}

// SetCity sets the value of City.
func (s *Address) SetCity(val string) {
	s.City = val
}

// Ref: #/components/schemas/Store
type Store struct {
	Name    string  `json:"name"`
	Address Address `json:"address"`
}

// GetName returns the value of Name.
func (s *Store) GetName() string {
	return s.Name
}

// GetAddress returns the value of Address.
func (s *Store) GetAddress() Address {
	return s.Address
}

// SetName sets the value of Name.
func (s *Store) SetName(val string) {
	s.Name = val
}

// SetAddress sets the value of Address.
func (s *Store) SetAddress(val Address) {
	s.Address = val
}
