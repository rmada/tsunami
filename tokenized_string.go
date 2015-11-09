package main

type tokenizedString struct {
	base string
}

func (this *tokenizedString) String() string {
	return this.base
}
