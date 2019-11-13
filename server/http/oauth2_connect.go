// Copyright (c) 2019-present Mattermost, Inc. All Rights Reserved.
// See License for license information.

package http

import (
	"fmt"
	"net/http"

	"golang.org/x/oauth2"

	"github.com/mattermost/mattermost-server/model"

	"github.com/mattermost/mattermost-plugin-msoffice/server/msgraph"
)

func (h *Handler) OAuth2Connect(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-ID")
	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	conf := msgraph.GetOAuth2Config(h.Config)
	state := fmt.Sprintf("%v_%v", model.NewId()[0:15], userID)
	err := h.OAuth2StateStore.Store(state)
	if err != nil {
		h.jsonError(w, err)
		return
	}

	url := conf.AuthCodeURL(state, oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusFound)
}
