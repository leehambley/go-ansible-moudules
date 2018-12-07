package file

// Stanza represents the Ansible "file" module entries
// see https://docs.ansible.com/ansible/2.5/modules/file_module.html
type Stanza struct {
	attributes string `yaml:"attributes,omitempty"`
	attr       string `yaml:"attr,omitempty"`

	follow bool `yaml:"follow,omitempty"`

	force bool `yaml:"force,omitempty"`

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
}

// StanzaWithDefaults as described in Ansible module documentation
// falsy defaults are not listed because of Go's sane handling of
// zero values.
//
// selevel default should be `s0` but it is omitted as selinux
// support was not yet needed and I wanted to avoid always setting
// an selinux property by interacting with the filesystem
var StanzaWithDefaults = Stanza{
	follow: true,
	state:  "file",
}
