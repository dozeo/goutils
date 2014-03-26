package goutils

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	DATE   = "auth[date]="
	SIG    = "&auth[signature]="
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
	r, get := h.canonicalRepresentation(parms, q)
	mac := hmac.New(sha1.New, []byte(secret))
	mac.Write([]byte(r))
	expectedMAC := mac.Sum(nil)
	date := strings.Replace(url.QueryEscape(parms["date"]), "+", "%20", -1)
	x := ""
	if len(get) > 0 {
		x = "&"
	}
	f := ""
	if len(u.Fragment) > 0 {
		f = "#" + u.Fragment
	}
	start := ""
	if u.Scheme != "" {
		start = u.Scheme + "://" + u.Host
	}
	return fmt.Sprintf("%s%s?%s%s%s%x%s%s%s", start, u.Path, DATE, date, SIG, expectedMAC, x, get, f)
}

func (h *Hmac) Validate(urlp string, secret string) bool {
	t, s := h.ValidateTime(urlp, secret)
	fmt.Printf("Validate: " + strconv.Itoa(s) + "\n")
	return t
}

func (h *Hmac) ValidateTime(urlp string, secret string) (bool, int) {
	p := strings.LastIndex(urlp, DATE)
	if p < 0 {
		fmt.Printf("<%s> does not contain <%s>\n", urlp, DATE)
		return false, -1
	}
	s := strings.LastIndex(urlp, SIG)
	if s < 0 {
		return false, -2
	}
	times := urlp[p+len(DATE) : s]
	tt, err := url.QueryUnescape(times)
	if err != nil {
		return false, -3
	}
	t, terr := time.Parse(FORMAT, tt)
	if terr != nil {
		return false, -4
	}
	if t.Unix() <= time.Now().UTC().Unix() {
		return false, int(t.Unix() - time.Now().UTC().Unix())
	}
	newu := h.signRequest(urlp, secret, t)
	if urlp == newu {
		ts, _ := strconv.Atoi(t.Format(time.RFC850))
		return true, ts
	} else {
		return false, -6
	}
}

func (h *Hmac) SignUrl(url string, secret string, ttl int) string {
	var t time.Time
	//       1395756739
	if ttl > 1000000000 {
		t = time.Unix(int64(ttl), 0)
	} else {
		t = time.Now().UTC()
		t = t.Add(time.Second * time.Duration(ttl))
	}
	return h.signRequest(url, secret, t)
}

func (h *Hmac) canonicalRepresentation(parms map[string]string, query map[string][]string) (string, string) {
	var rep string
	rep += strings.ToUpper(parms["method"]) + "\n"
	delete(parms, "method")
	rep += "date:" + parms["date"] + "\n"
	rep += "nonce:" + parms["nonce"] + "\n"
	rep += parms["path"]
	get := ""
	if query != nil && len(query) > 0 {
		mk := make([]string, len(query))
		var i uint32
		i = 0
		for key, _ := range query {
			mk[i] = key
			i++
		}
		sort.Strings(mk)
		first := true
		for _, key := range mk {
			if key == "auth[date]" {
				continue
			}
			if key == "auth[signature]" {
				continue
			}
			if first {
				first = false
				rep += "?"
			} else {
				rep += "&"
				get += "&"
			}
			rep += key
			rep += "="
			t, _ := url.QueryUnescape(query[key][0])
			rep += t
			get += key
			get += "="
			get += url.QueryEscape(query[key][0])
		}
	}
	return rep, get
}
