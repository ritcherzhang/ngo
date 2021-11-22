// Copyright Ngo Authors
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

package protocol

import (
	"errors"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	var err error
	err = &Error{
		Code: SystemError,
		Err:  os.ErrClosed,
	}

	e := &Error{}
	assert.True(t, errors.As(err, &e))
	assert.True(t, errors.Is(err, os.ErrClosed))
}

func TestFail(t *testing.T) {
	statusCode, body := ErrorJsonBody(SystemError)
	assert.Equal(t, http.StatusInternalServerError, statusCode)
	assert.Equal(t, SystemError, body.Code)
	assert.Nil(t, body.Data)

	statusCode, body = Fail(11111, "ssss")
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, 11111, body.Code)
	assert.Equal(t, "ssss", body.Message)
}
