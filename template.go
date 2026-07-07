package azw3

import (
	"text/template"

	"github.com/behringer24/mobi/records"
)

// kf8TemplateString mirrors the mobi skeleton template but adds the EPUB
// namespace (xmlns:epub). That makes epub:type attributes — used for
// footnotes (noteref / footnote) and other semantics — valid in the
// generated KF8 HTML, which is what drives popup footnotes on Kindle.
//
// It is handed to the mobi Book via OverrideTemplate (see toMobiBook) so the
// upstream mobi fork keeps its own default template unchanged; the KF8-only
// namespace concern lives here, in this project.
const kf8TemplateString = `<?xml version="1.0" encoding="UTF-8"?>
<html xmlns="http://www.w3.org/1999/xhtml" xmlns:epub="http://www.idpf.org/2007/ops">
  <head>
    <title>{{ .Mobi.Title }}</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
    {{- range $i, $_ := .Mobi.CSSFlows }}
    <link rel="stylesheet" type="text/css" href="kindle:flow:{{ $i | inc | base32 }}?mime=text/css"/>
    {{- end }}
  </head>
  <body aid="{{ .Chunk.ID | base32 }}">
`

// kf8Template uses the same helper functions the mobi default template
// relies on (base32 via the exported records.To32, and inc), so the
// inventory passed by mobi at render time resolves identically.
var kf8Template = template.Must(template.New("azw3").Funcs(template.FuncMap{
	"inc":    func(i int) int { return i + 1 },
	"base32": records.To32,
}).Parse(kf8TemplateString))
