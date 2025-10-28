package eval

import (
	"roach/version"
)

const (
	LANG_OBJ  = "LANG_OBJ"
	lang_name = "lang"
)

type LangObj struct{}

func NewLangObj() Object {
	ret := &LangObj{}
	SetGlobalObj(lang_name, ret)

	SetGlobalObj(lang_name+".Version", NewString(version.Version))
	SetGlobalObj(lang_name+".Major", NewString(version.Major))
	SetGlobalObj(lang_name+".Minor", NewString(version.Minor))
	SetGlobalObj(lang_name+".Release", NewString(version.Release))
	SetGlobalObj(lang_name+".BuildNumber", NewString(version.BuildNumber))

	return ret
}

func (t *LangObj) Inspect() string {
	return "<" + lang_name + ">"
}

func (t *LangObj) Type() ObjectType { return LANG_OBJ }

func (t *LangObj) CallMethod(line string, scope *Scope, method string, args ...Object) Object {
	switch method {
	}
	return NewError(line, NOMETHODERROR, method, t.Type())
}
