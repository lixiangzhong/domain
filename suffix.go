package domain

import (
	"strings"
	"sync"
)

type Suffixs struct {
	suffix *sync.Map
}

func NewSuffixs() *Suffixs {
	return &Suffixs{suffix: new(sync.Map)}
}

//Load  加载后缀 ,如 []string{".com",".cn}
func (s *Suffixs) Load(suffixs []string) {
	for _, v := range suffixs {
		v = strings.TrimSpace(v)
		v = "." + strings.Trim(v, ".")
		s.suffix.Store(v, true)
	}
}

//MatchDomain 根据已有后缀匹配出主域名
func (s *Suffixs) MatchDomain(host string) (string, bool) {
	domain := host
	var ok bool
	for {
		suffix := s.cutHead(domain)
		if suffix == domain {
			break
		}
		_, ok = s.suffix.Load(suffix)
		if ok {
			break
		} else {
			domain = suffix
		}
	}
	return strings.Trim(domain, "."), ok
}

func (*Suffixs) cutHead(host string) string {
	s := strings.Trim(host, ".")
	n := strings.Index(s, ".")
	if n < 0 {
		return s
	}
	return s[n:]
}
