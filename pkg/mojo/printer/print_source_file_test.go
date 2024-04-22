package printer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mojo-lang/mojo/pkg/context"
	"github.com/mojo-lang/mojo/pkg/mojo/parser/syntax"
)

func TestPrinter_PrintSourceFile(t *testing.T) {
	const file = `
// Copyright 2021 Mojo-lang.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/// ApplyExpr - Superclass of various function calls, which apply an argument to
/// a function to get a result.
type ApplyExpr : Expr {
    // 函数名称调用
    // lambda 匿名调用
    callee: Expression @10
}

`
	const expect = `// Copyright 2021 Mojo-lang.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/// ApplyExpr - Superclass of various function calls, which apply an argument to
/// a function to get a result.
type ApplyExpr: Expr {
    // 函数名称调用
    // lambda 匿名调用
    callee: Expression @10
}`

	parser := &syntax.Parser{}
	sourceFile, err := parser.ParseString(context.Empty(), file)
	assert.NoError(t, err)

	p := New(&Config{}).PrintSourceFile(context.Empty(), sourceFile)
	assert.Equal(t, expect, p.Buffer.String())
}
