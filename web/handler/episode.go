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

package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/bangumi/server/ctrl"
	"github.com/bangumi/server/domain"
	"github.com/bangumi/server/internal/pkg/errgo"
	"github.com/bangumi/server/internal/pkg/generic/slice"
	"github.com/bangumi/server/web/req"
	"github.com/bangumi/server/web/res"
)

func (h Handler) GetEpisode(c echo.Context) error {
	u := h.GetHTTPAccessor(c)

	id, err := req.ParseEpisodeID(c.Param("id"))
	if err != nil {
		return err
	}

	e, err := h.ctrl.GetEpisode(c.Request().Context(), id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return res.ErrNotFound
		}

		return errgo.Wrap(err, "failed to get episode")
	}

	_, err = h.ctrl.GetSubject(c.Request().Context(), u.Auth, e.SubjectID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return res.ErrNotFound
		}

		return errgo.Wrap(err, "failed to find subject of episode")
	}

	return c.JSON(http.StatusOK, res.ConvertModelEpisode(e))
}

func (h Handler) ListEpisode(c echo.Context) error {
	u := h.GetHTTPAccessor(c)

	page, err := req.GetPageQuery(c, req.EpisodeDefaultLimit, req.EpisodeMaxLimit)
	if err != nil {
		return err
	}

	epType, err := req.ParseEpTypeOptional(c.QueryParam("type"))
	if err != nil {
		return err
	}

	subjectID, err := req.ParseSubjectID(c.QueryParam("subject_id"))
	if err != nil {
		return err
	}
	if subjectID == 0 {
		return res.BadRequest("missing required query `subject_id`")
	}

	_, err = h.ctrl.GetSubject(c.Request().Context(), u.Auth, subjectID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return res.ErrNotFound
		}
		return errgo.Wrap(err, "failed to get subject")
	}

	episodes, count, err := h.ctrl.ListEpisode(c.Request().Context(), subjectID, epType, page.Limit, page.Offset)
	if err != nil {
		if errors.Is(err, ctrl.ErrOffsetTooBig) {
			return res.BadRequest("offset should be less than or equal to " + strconv.FormatInt(count, 10))
		}
		return errgo.Wrap(err, "failed to list episode")
	}

	var data = make([]res.Episode, len(episodes))
	for i, episode := range episodes {
		data[i] = res.ConvertModelEpisode(episode)
	}

	return c.JSON(http.StatusOK, res.PagedG[res.Episode]{
		Limit:  page.Limit,
		Offset: page.Offset,
		Data:   slice.Map(episodes, res.ConvertModelEpisode),
		Total:  count,
	})
}
