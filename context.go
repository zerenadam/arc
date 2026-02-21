package arc

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: r,
	}
}

func (c *Context) JSON(status int, data any) error {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(status)

	if err := json.NewEncoder(c.Writer).Encode(data); err != nil {
		return err
	}
	return nil

}

func (c *Context) XML(status int, data any) error {
	c.Writer.Header().Set("Content-Type", "application/xml")
	c.Writer.WriteHeader(status)

	if err := xml.NewEncoder(c.Writer).Encode(data); err != nil {
		return err
	}
	return nil

}

func (c *Context) String(status int, data string) error {
	c.Writer.Header().Set("Content-Type", "text/plain")
	c.Writer.WriteHeader(status)

	_, err := c.Writer.Write([]byte(data))

	if err != nil {
		return err
	}
	return nil
}
