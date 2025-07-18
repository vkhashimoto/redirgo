package links

import (
	"testing"
)

func TestFindRedirectionExistingHost(t *testing.T) {
	expected := "https://github.com/vkhashimoto/redirgo"
	LoadLinks("../config/links.toml")
	link, _ := FindRedirection("localhost:8080", "source")
	if link != expected {
		t.Errorf("Wanted: %s\nGot: %s", expected, link)
		t.FailNow()
	}
}

func TestFindRedirectionFallbackHost(t *testing.T) {
	expected := "https://example.org"
	LoadLinks("../config/links.toml")
	link, _ := FindRedirection("local-host:8080", "example")
	if link != expected {
		t.Errorf("Wanted: %s\nGot: %s", expected, link)
		t.FailNow()
	}
}

func TestFindRedirectionLinkNotFound(t *testing.T) {
	expected := "Link not found"
	LoadLinks("../config/links.toml")
	_, err := FindRedirection("localhost:8080", "git")
	if err == nil {
		t.Error("Error expected")
		t.FailNow()
	}
	if err.Error() != expected {
		t.Errorf("Wanted: `%s`\nGot: `%s`", expected, err.Error())
		t.FailNow()
	}
}

func TestFindRedirectionRootPath(t *testing.T) {
	expected := "https://github.com/vkhashimoto"
	LoadLinks("../config/links.toml")
	link, err := FindRedirection("localhost:8080", "")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		t.FailNow()
	}
	if link != expected {
		t.Errorf("Wanted: %s\nGot: %s", expected, link)
		t.FailNow()
	}
}

func TestFindRedirectionLinkNotFoundRootPath(t *testing.T) {
	expected := "Link not found"
	LoadLinks("../config/links.toml")
	_, err := FindRedirection("localhost:8081", "")
	if err == nil {
		t.Error("Expected error for unconfigured root path")
		t.FailNow()
	}
	if err.Error() != expected {
		t.Errorf("Wanted: `%s`\nGot: `%s`", expected, err.Error())
		t.FailNow()
	}
}
