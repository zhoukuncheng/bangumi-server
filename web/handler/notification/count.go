// SPDX-License-Identifier: AGPL-3.0-only
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>

package notification

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/bangumi/server/internal/pkg/errgo"
)

func (h Notification) Count(c echo.Context) error {
	accessor := h.Common.GetHTTPAccessor(c)
	count, err := h.notificationRepo.Count(c.Request().Context(), accessor.ID)
	if err != nil {
		return errgo.Wrap(err, "failed to count notification")
	}
	return c.JSON(http.StatusOK, count)
}
