package lineinfile

import "testing"
import "github.com/leehambley/go-ansible-modules/test_helper"

func Test_Lineinfile_Unit(t *testing.T) {

	t.Parallel()

	var h = test_helper.H(t)
	
	t.Run("defaults match ansible 2.5", func(t *testing.T) {
		
		// https://docs.ansible.com/ansible/2.5/modules/lineinfile_module.html#parameters
		
		h.StrEql("", StanzaWithDefaults.attributes)
		h.StrEql("", StanzaWithDefaults.attr)

		h.BoolEql(false, StanzaWithDefaults.backrefs)

		h.BoolEql(false, StanzaWithDefaults.backup)

		h.BoolEql(false, StanzaWithDefaults.create)

		h.BoolEql(false, StanzaWithDefaults.firstmatch)

		h.StrEql("", StanzaWithDefaults.group)
		h.StrEql("", StanzaWithDefaults.mode)
		h.StrEql("", StanzaWithDefaults.owner)

		h.StrEql("", StanzaWithDefaults.path)
		h.StrEql("", StanzaWithDefaults.dest)
		h.StrEql("", StanzaWithDefaults.name)

		h.BoolEql(false, StanzaWithDefaults.recurse)

		h.StrEql("p0", StanzaWithDefaults.selevel)
		h.StrEql("", StanzaWithDefaults.serole)
		h.StrEql("", StanzaWithDefaults.setype)
		h.StrEql("", StanzaWithDefaults.seuser)

		h.StrEql("", StanzaWithDefaults.src)

		h.StrEql("present", StanzaWithDefaults.state)

		h.StrEql("no", StanzaWithDefaults.unsafeWrites)

		h.StrEql("", StanzaWithDefaults.validate)


	})

}
