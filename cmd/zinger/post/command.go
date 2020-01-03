package post

import (
	"flag"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/mccanne/charm"
	"github.com/mccanne/zinger/cmd/zinger/root"
)

var Post = &charm.Spec{
	Name:  "post",
	Usage: "post [options]",
	Short: "post data to zinger",
	Long: `
The post command allows you to...
TBD.`,
	New: New,
}

func init() {
	root.Zinger.Add(Post)
}

type Command struct {
	*root.Command
	*http.Client
	zingerAddr string
}

func New(parent charm.Command, f *flag.FlagSet) (charm.Command, error) {
	c := &Command{
		Command: parent.(*root.Command),
		Client:  &http.Client{},
	}
	f.StringVar(&c.zingerAddr, "a", ":9890", "[addr]:port to send to")
	return c, nil
}

func (c *Command) Run(args []string) error {
	if len(args) == 0 {
		args = []string{"-"}
	}
	for _, fname := range args {
		f, err := os.Open(fname)
		if err != nil {
			return err
		}
		_, err = c.post(f)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Command) post(body io.Reader) ([]byte, error) {
	url := "http://" + c.zingerAddr
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	return b, err
}
