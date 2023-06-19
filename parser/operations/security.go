package operations

import (
	oas "github.com/parvez3019/go-swagger3/openApi3Schema"
	"strings"
)

func (p *parser) parseSecurity(pkgPath, pkgName string, operation *oas.OperationObject, comment string) error {
	//sourceString := strings.TrimSpace(comment[len("@Security"):])
	sourceString := comment

	fields := strings.Split(sourceString, " ")
	security := map[string][]string{
		fields[0]: fields[1:],
	}
	operation.Security = append(operation.Security, security)

	return nil
}
