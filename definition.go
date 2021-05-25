package goapidoc

// Definition represents an api definition information.
type Definition struct {
	name string
	desc string

	generics   []string
	properties []*Property
}

// NewDefinition creates a default Definition with given arguments.
func NewDefinition(title string, desc string) *Definition {
	return &Definition{name: title, desc: desc}
}

// Generics sets the generics in Definition.
func (d *Definition) Generics(generics ...string) *Definition {
	d.generics = generics
	return d
}

// Properties sets the properties in Definition.
func (d *Definition) Properties(properties ...*Property) *Definition {
	d.properties = properties
	return d
}

// Property represents a definition property information.
type Property struct {
	name     string
	typ      string
	required bool
	desc     string

	allowEmpty bool
	dft        interface{}
	example    interface{}
	enums      []interface{}
	minLength  int
	maxLength  int
	minimum    int
	maximum    int
}

// NewProperty creates a default Property with given arguments.
func NewProperty(name string, typ string, req bool, desc string) *Property {
	return &Property{name: name, typ: typ, required: req, desc: desc}
}

// AllowEmpty sets the allowEmpty in Property.
func (p *Property) AllowEmpty(allow bool) *Property {
	p.allowEmpty = allow
	return p
}

// Default sets the dft in Property.
func (p *Property) Default(dft interface{}) *Property {
	p.dft = dft
	return p
}

// Example sets the example in Property.
func (p *Property) Example(example interface{}) *Property {
	p.example = example
	return p
}

// Enum sets the enums in Property.
func (p *Property) Enum(enums ...interface{}) *Property {
	p.enums = enums
	return p
}

// MinLength sets the minLength in Property.
func (p *Property) MinLength(min int) *Property {
	p.minLength = min
	return p
}

// MaxLength sets the maxLength in Property.
func (p *Property) MaxLength(max int) *Property {
	p.maxLength = max
	return p
}

// Length sets the minLength and maxLength in Property.
func (p *Property) Length(min, max int) *Property {
	p.minLength = min
	p.maxLength = max
	return p
}

// Minimum sets the minimum in Property.
func (p *Property) Minimum(min int) *Property {
	p.minimum = min
	return p
}

// Maximum sets the maximum in Property.
func (p *Property) Maximum(max int) *Property {
	p.maximum = max
	return p
}

// Minimum sets the minimum and maximum in Property.
func (p *Property) MinMaximum(min, max int) *Property {
	p.minimum = min
	p.maximum = max
	return p
}
