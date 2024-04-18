package shortener

type RedirectSerializer interface {
	Decode(intput []byte) (*Redirect, error)
	Encode(input *Redirect) ([]byte, error)
}
