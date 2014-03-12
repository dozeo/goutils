package goutils

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	DATE   = "auth%5Bdate%5D="
	SIG    = "&auth%5Bsignature%5D="
	FORMAT = "Mon, 02 Jan 2006 15:04:05 GMT"
)

var HMAC Hmac

type Hmac struct {
}

func (h *Hmac) signRequest(urlp string, secret string, t time.Time) string {
	u, _ := url.Parse(urlp)
	q := u.Query()
	parms := make(map[string]string)
	parms["method"] = "GET"
	parms["date"] = t.Format(FORMAT)
	//parms["nonce"] = ""
	parms["path"] = u.Path
	r := h.canonicalRepresentation(parms, q)
	mac := hmac.New(sha1.New, []byte(secret))
	mac.Write([]byte(r))
	expectedMAC := mac.Sum(nil)
	date := strings.Replace(url.QueryEscape(parms["date"]), "+", "%20", -1)
	x := ""
	if len(q) > 0 {
		x = "&"
	}
	f := ""
	if len(u.Fragment) > 0 {
		f = "#" + u.Fragment
	}
	return fmt.Sprintf("%s://%s%s?%s%s%s%s%s%x%s", u.Scheme, u.Host, u.Path, u.RawQuery, x, DATE, date, SIG, expectedMAC, f)
}

func (h *Hmac) Validate(urlp string, secret string) bool {
	u, _ := url.Parse(urlp)
	p := strings.LastIndex(urlp, DATE)
	if p < 0 {
		return false
	}
	s := strings.LastIndex(urlp, SIG)
	times := urlp[p+len(DATE) : s]
	tt, err := url.QueryUnescape(times)
	if err != nil {
		return false
	}
	var baseurl string
	if urlp[p-1] == '&' {
		baseurl = urlp[:p-1]
	} else {
		baseurl = urlp[:p]
	}
	t, terr := time.Parse(FORMAT, tt)
	if terr != nil {
		return false
	}
	if t.Unix() < time.Now().UTC().Unix() {
		return false
	}
	f := ""
	if len(u.Fragment) > 0 {
		f = "#" + u.Fragment
	}
	newu := h.signRequest(baseurl+f, secret, t)
	if urlp == newu {
		return true
	} else {
		fmt.Printf("\n%s\n%s\n", urlp, newu)
		return false
	}
}

func (h *Hmac) SignUrl(url string, secret string, ttl int) string {
	t := time.Now().UTC()
	t = t.Add(time.Second * time.Duration(ttl))
	return h.signRequest(url, secret, t)
}

func (h *Hmac) canonicalRepresentation(parms map[string]string, query map[string][]string) string {
	var rep string
	rep += strings.ToUpper(parms["method"]) + "\n"
	delete(parms, "method")
	rep += "date:" + parms["date"] + "\n"
	rep += "nonce:" + parms["nonce"] + "\n"
	rep += parms["path"]
	if query != nil && len(query) > 0 {
		mk := make([]string, len(query))
		var i uint32
		i = 0
		for key, _ := range query {
			mk[i] = key
			i++
		}
		sort.Strings(mk)
		rep += "?"
		first := true
		for _, key := range mk {
			if first {
				first = false
			} else {
				rep += "&"
			}
			rep += key
			rep += "="
			rep += query[key][0]
		}
	}
	return rep
}
