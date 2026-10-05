// Copyright 2026 The OpenSandbox Authors
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

package solution

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const AccessTokenHeader = "X-EXECD-ACCESS-TOKEN"

func NewRouter(accessToken string) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	r.Use(func(c *gin.Context) {
		if accessToken == "" {
			c.Next()
			return
		}
		if c.GetHeader(AccessTokenHeader) != accessToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    "UNAUTHORIZED",
				"message": "missing or invalid access token",
			})
			return
		}
		c.Next()
	})
	r.POST("/command", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"accepted": true})
	})
	return r
}
