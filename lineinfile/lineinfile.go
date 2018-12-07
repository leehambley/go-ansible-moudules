package lineinfile

// Stanza represents the Ansible "lineinfile" module entries
// see https://docs.ansible.com/ansible/2.5/modules/lineinfile_module.html
type Stanza struct {
	attributes string `yaml:"attributes,omitempty"`
	attr       string `yaml:"attr,omitempty"`

	backrefs bool `yaml:"backrefs,omitempty"`

	backup bool `yaml:"backup,omitempty"`

	create bool `yaml:"create,omitempty"`

	firstmatch bool `yaml:"firstmatch,omitempty"`

	group string `yaml:"group,omitempty"`

	mode string `yaml:"mode,omitempty"`

	owner string `yaml:"owner,omitempty"`

	path string `yaml:"path,omitempty"`
	dest string `yaml:"dest,omitempty"`
	name string `yaml:"name,omitempty"`

	recurse bool `yaml:"recurse,omitempty"`

	selevel string `yaml:"selevel,omitempty"`
	serole  string `yaml:"serole,omitempty"`
	setype  string `yaml:"setype,omitempty"`
	seuser  string `yaml:"seuser,omitempty"`

	src string `yaml:"src,omitempty"`

	state string `yaml:"state,omitempty"`

	unsafeWrites string `yaml:"unsafe_writes,omitempty"`

	validate string `yaml:"validate,omitempty"`
}

// StanzaWithDefaults as described in Ansible module documentation
// falsy defaults are not listed because of Go's sane handling of
// zero values.
//
// selevel default should be `s0` but it is omitted as selinux
// support was not yet needed and I wanted to avoid always setting
// an selinux property by interacting with the filesystem
var StanzaWithDefaults = Stanza{
	selevel:      "p0",
	state:        "present",
	unsafeWrites: "no",
}

// t0 booking/:uuid:
// 	+ requested {hotel: here, bed_size: double, offer_id: abc}

// t5 transaction/:uuid:
// 	+ associate_booking {booking_urn: booking/:uuid:}
// 	+ set_amount {amount: 15600, currency: eur, unit: cent}
// 	+ add_payment_method {type: card, subtype: visa, number: ...}
// 	+ add_payment_method {type: paypal, bearer_token: ...}

// <some "side effect from a worker">

// t5 transaction/:uuid:
// 	associate_booking {booking_urn: booking/:uuid:}
// 	set_amount {amount: 15600, currency: eur, unit: cent}
// 	+ confirmed_by_issuer {card_type: visa, mandant_ref: :uuid:}

// <some "side effect from a worker">

// t10 booking/:uuid
// 	+ requested {hotel: here, bed_size: double, offer_id: abc}
// 	+ confirmed {room_number: 123, booking_ref: :uuid:}
