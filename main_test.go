package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func writeHTML(content string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, strings.TrimSpace(content))
	})
}

func Test_linkMetaWithzkSync(t *testing.T) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ts := httptest.NewServer(writeHTML(`<!doctype html>
<html>
<body>
  <div id="content">the content</div>
</body>
</html>`))
	defer ts.Close()

	const expr = `(function(d, id, v) {
		var b = d.querySelector('body');
		var el = d.createElement('div');
		el.id = id;
		el.innerText = v;
		b.insertBefore(el, b.childNodes[0]);
	})(document, %q, %q);`

	var nodes []*cdp.Node
	if err := chromedp.Run(ctx,
		chromedp.Navigate(ts.URL),
		chromedp.Nodes(`document`, &nodes, chromedp.ByJSPath),
		chromedp.WaitVisible(`#content`),
		chromedp.ActionFunc(func(ctx context.Context) error {
			s := fmt.Sprintf(expr, "thing", "a new thing!")
			_, exp, err := runtime.Evaluate(s).Do(ctx)
			if err != nil {
				return err
			}
			if exp != nil {
				return exp
			}
			return nil
		}),
		chromedp.WaitVisible(`#thing`),
	); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Document tree:")
	fmt.Print(nodes[0].Dump("  ", "  ", false))

}

func Test2(t *testing.T) {

}
