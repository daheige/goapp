package setting

type Option func(s *Setting)

func WithWatchFile(b bool) Option {
	return func(s *Setting) {
		s.watchFile = true
	}
}
