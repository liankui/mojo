package scaffolding

import (
	"os"
	path2 "path"

	"github.com/mojo-lang/core/go/pkg/mojo/core"

	"github.com/mojo-lang/mojo/pkg/util"
)

const (
	HelloWorldPackageName = "hello_world"
	HelloWorldOrg         = "mojo-lang.org"
	HelloWorldRepo        = "git.company.com/examples/hello-world"
)

const HandlerMethod = `
func (s helloWorld) GetEcho(ctx context.Context, in *pb.GetEchoRequest) (*pb.Echo, error) {
	resp := pb.Echo{
		Name:    in.Name,
		Message: "Hello, " + in.Name + "!",
	}
	return &resp, nil
}
`

func GenerateHelloWorldFiles(output string) error {
	files := []*util.GeneratedFile{{
		Name:    path2.Join(output, "hello-world/mojo/hello-world/v1/echo.mojo"),
		Content: helloWorldEcho,
	}, {
		Name:    path2.Join(output, "hello-world/mojo/hello-world/v1/hello_world.mojo"),
		Content: helloWorldService,
	}}

	for _, f := range files {
		path := path2.Dir(f.Name)
		if err := core.CreateDir(path); err != nil {
			return err
		}
		if err := os.WriteFile(f.Name, []byte(f.Content), 0o666); err != nil {
			return err
		}
	}
	return nil
}
