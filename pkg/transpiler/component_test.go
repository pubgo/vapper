package components

import (
	"bytes"
	"github.com/pubgo/errors"
	"strings"
	"testing"
)

var goodTpl = `<nav>
<ul>
	<li><a class='{path === "/"  ? "selected" : ""}' href='.'>home</a></li>
	<li><a class='{path === "/about"  ? "selected" : ""}' href='about'>about</a></li>

	<!-- for the blog link, we're using rel=prefetch so that Factor prefetches
		 the blog data when we hover over the link or tap it on a touchscreen -->
	<li><a rel=prefetch class='{path.startsWith("/blog")  ? "selected" : ""}' href='blog'>blog</a></li>
</ul>
</nav>`
var goodStyle = `<style>
	nav {
		border-bottom: 1px solid rgba(170,30,30,0.1);
		font-weight: 300;
		padding: 0 1em;
	}

	ul {
		margin: 0;
		padding: 0;
	}

	/* clearfix */
	ul::after {
		content: '';
		display: block;
		clear: both;
	}

	li {
		display: block;
		float: left;
	}

	.selected {
		position: relative;
		display: inline-block;
	}

	.selected::after {
		position: absolute;
		content: '';
		width: calc(100% - 1em);
		height: 2px;
		background-color: rgb(170,30,30);
		display: block;
		bottom: -1px;
	}

	a {
		text-decoration: none;
		padding: 1em 0.5em;
		display: block;
	}
</style>`

func TestQuoted(t *testing.T) {
	defer errors.Assert()

	c := parseComponent(t)
	qs := c.QuotedStyle()
	errors.T(strings.Compare(qs[0:1], "`") != 0, "expected style to start with backtick, got: %s", qs[0:1])

	qt := c.QuotedTemplate()
	errors.T(strings.Compare(qt[len(qt)-1:], "`") != 0, "expected template to start with backtick, got: %s", qt[len(qt)-1:])
}
func parseComponent(t *testing.T) *Component {

	good := goodTpl + "\n" + goodStyle
	b := bytes.NewBuffer([]byte(good))
	return Parse(b, "Good")
}

func TestTransformUnparsed(t *testing.T) {
	defer errors.Assert()

	c := &Component{}
	b := new(bytes.Buffer)
	c.Transform(b)
}

func TestTransform(t *testing.T) {
	defer errors.Assert()

	c := parseComponent(t)
	c.Package = "mypackage"
	b := new(bytes.Buffer)
	c.Transform(b)
	// TODO: Compare against golden
}
