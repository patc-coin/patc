// Copyright 2015 The go-pc Authors
// This file is part of the go-pc library.
//
// The go-pc library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-pc library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-pc library. If not, see <http://www.gnu.org/licenses/>.

package runtime_test

import (
	"fmt"

	"pc.dp.tc/common"
	"pc.dp.tc/core/vm/runtime"
)

func ExampleExecute() {
	ret, _, err := runtime.Execute(common.Hex2Bytes("6060604052600a8060106000396000f360606040526008565b00"), nil, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)
	// Output:
	// [96 96 96 64 82 96 8 86 91 0]
}
