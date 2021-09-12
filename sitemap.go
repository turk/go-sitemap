package sitemap

import (
	"encoding/xml"
	"io"
)

type URL struct {
	Loc     string `xml:"loc,omitempty"`
	Lastmod string `xml:"lastmod,omitempty"`
}

type Sitemap struct {
	XMLName   xml.Name  `xml:"urlset"`
	Ns     string    `xml:"xmlns,attr"`
	Writer io.Writer `xml:"-"`
	URLs   []URL     `xml:"url"`
	Indent bool      `xml:"-"`
}

func NewSitemap(writer io.Writer, indent bool) *Sitemap {
	return &Sitemap{
		Writer: writer,
		URLs:   make([]URL, 0),
		Ns:     "http://www.sitemaps.org/schemas/sitemap/0.9",
		Indent: indent,
	}
}

func (s *Sitemap) Add(url string, lastmod string) {
	s.URLs = append(
		s.URLs,
		URL{
			Loc:     url,
			Lastmod: lastmod,
		},
	)
}

func (s *Sitemap) Write() error {
	xmlEncoder := xml.NewEncoder(s.Writer)

	if s.Indent {
		xmlEncoder.Indent("", "  ")
	}

	_, err := s.Writer.Write([]byte(xml.Header))
	if err != nil {
		return err
	}

	err = xmlEncoder.Encode(s)
	if err != nil {
		return err
	}

	return nil
}
