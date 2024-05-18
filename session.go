package csvq

import (
	"context"
	"io"

	"github.com/mithrandie/csvq/lib/query"
)

func getSession() *query.Session {
	session := query.NewSession()
	session.SetStdout(&query.Discard{})
	session.SetStderr(&query.Discard{})

	return session
}

func SetStdin(r io.ReadCloser) error {
	return SetStdinContext(context.Background(), r)
}

func SetStdinContext(ctx context.Context, r io.ReadCloser) error {
	return getSession().SetStdinContext(ctx, r)
}

func SetStdout(w io.WriteCloser) {
	getSession().SetStdout(w)
}

func SetOutFile(w io.Writer) {
	getSession().SetOutFile(w)
}
