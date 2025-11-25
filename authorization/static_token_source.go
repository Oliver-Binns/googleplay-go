package authorization

func StaticTokenSource(
	token string,
) TokenSource {
	return &staticTokenSource{
		token: token,
	}
}

type staticTokenSource struct {
	token string
}

func (s *staticTokenSource) Token() (string, error) {
	return s.token, nil
}
