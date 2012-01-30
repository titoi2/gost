package main

type TokenScanner struct {
	buf []byte
}

func (t *TokenScanner) Init(s []byte) {
	t.buf = s
}

func (t *TokenScanner) GetToken() string {
	var i int
	for i = 0; i < len(t.buf); i++ {
		if t.buf[i] > ' ' {
			break
		}
	}
	t.buf = t.buf[i:]
	for i = 0; i < len(t.buf); i++ {
		if t.buf[i] <= ' ' {
			break
		}
	}
	s := string(t.buf[:i])
	t.buf = t.buf[i:]
	return s
}
