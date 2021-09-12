package sitemap

import (
	"encoding/xml"
	"io"
)

type SitemapIndex struct {
	XMLName xml.Name  `xml:"sitemapindex"`
	Ns      string    `xml:"xmlns,attr"`
	Writer  io.Writer `xml:"-"`
	URLs    []URL     `xml:"sitemap"`
	Indent  bool      `xml:"-"`
}

func NewSitemapIndex(writer io.Writer, indent bool) *SitemapIndex {
	return &SitemapIndex{
		Writer: writer,
		URLs:   make([]URL, 0),
		Ns:     "http://www.sitemaps.org/schemas/sitemap/0.9",
		Indent: indent,
	}
}

func (s *SitemapIndex) Add(url string) {
	s.URLs = append(
		s.URLs,
		URL{
			Loc: url,
		},
	)
}

func (s *SitemapIndex) Write() error {
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
