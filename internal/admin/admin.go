/*
Package admin provides admin related functions
*/
package admin

import (
	"encoding/json"
	"io"
)

// Admin contains name and the public key for the admin to encrypt with
type Admin struct {
	Name string `json:"name,omitempty"`
	Key  string `json:"key"`
	Type int    `json:"type"`
}

// Admins is a collection of admins
type Admins struct {
	Adms []Admin `json:"admins"`
}

// Save serializes and writes the admins to a file
func (a *Admins) Save(w io.Writer) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	return enc.Encode(a)
}

// Read reads the admins file
func (a *Admins) Read(r io.Reader) error {
	return json.NewDecoder(r).Decode(&a)
}

// Add adds a new admin
func (a *Admins) Add(admin Admin) {
	a.Adms = append(a.Adms, admin)
}

// NewAdmins returns an empty slice of admin
func NewAdmins() Admins {
	return Admins{
		Adms: make([]Admin, 0),
	}
}
