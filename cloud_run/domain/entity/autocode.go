package entity

type AutocodeURL string

func (a AutocodeURL) String() string {
	return string(a)
}
