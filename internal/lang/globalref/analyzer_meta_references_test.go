package globalref

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/terraform/internal/addrs"
)

func TestAnalyzerMetaReferences(t *testing.T) {
	azr := testAnalyzer(t, "assorted")

	ref, diags := addrs.ParseRefStr("test_thing.single.string")
	if diags.HasErrors() {
		t.Fatalf("errors %s", diags.Err())
	}
	t.Logf("test %s", ref.DisplayString())
	_, refs := azr.MetaReferences(addrs.RootModuleInstance, ref)

	for _, ref := range refs {
		t.Logf("test %s", ref.DisplayString())
	}

	spew.Dump(refs)
}
