package filesys_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/HibiscusCollective/go-toolbox/cmd/hookgen/internal/filesys"
	"github.com/HibiscusCollective/go-toolbox/pkg/must"
	"github.com/onsi/gomega"
	"github.com/spf13/afero"
)

func TestCreateFile(t *testing.T) {
	t.Parallel()

	scns := map[string]func(t testing.TB, g gomega.Gomega){
		"should create a writeable file": func(t testing.TB, g gomega.Gomega) {
			const (
				path    = "test.txt"
				content = "hello, world!"
			)

			afs := afero.NewMemMapFs()

			f, err := filesys.New(afs).Create(path)
			defer f.Close()
			g.Expect(err).To(gomega.BeNil())

			fmt.Fprintf(f, content)

			got, err := afs.Open(path)
			g.Expect(err).To(gomega.BeNil())
			g.Expect(string(must.GetOrFailTest(t, func() ([]byte, error) { return io.ReadAll(got) }))).To(gomega.Equal(content))
		},
		"should create a writeable file in a directory": func(t testing.TB, g gomega.Gomega) {
			const (
				path    = "dir/test.txt"
				content = "hello, world!"
			)

			afs := afero.NewMemMapFs()

			f, err := filesys.New(afs).Create(path)
			defer f.Close()
			g.Expect(err).To(gomega.BeNil())

			fmt.Fprintf(f, content)

			got, err := afs.Open(path)
			g.Expect(err).To(gomega.BeNil())
			g.Expect(string(must.GetOrFailTest(t, func() ([]byte, error) { return io.ReadAll(got) }))).To(gomega.Equal(content))
		},
	}

	for name, scn := range scns {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			scn(t, gomega.NewWithT(t))
		})
	}
}
