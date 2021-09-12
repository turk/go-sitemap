# Install
`go get github.com/turk/go-sitemap`

## Example for sitemapindex

```go
func () main(c *gin.Context) {
    s := sitemap.NewSitemapIndex(c.Writer, true)
    s.Add("https://example.org/sitemaps/1.xml")
    s.Add("https://example.org/sitemaps/2.xml")
    s.Add("https://example.org/sitemaps/3.xml")
    s.Write()
}
```

## Example for sitemap

```go
func () main(c *gin.Context) {
    s := sitemap.NewSitemap(c.Writer, true)
    s.Add(
        "https://example.org/who-is-john-doe",
        "2021-09-01 15:04:05",
    )
    s.Write()
}
```
## Support
Please feel free to ask for new features, create pr or report issue.
