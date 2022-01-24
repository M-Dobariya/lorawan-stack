// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.3.1
// - protoc             v3.9.1
// source: lorawan-stack/api/identityserver.proto

package ttnpb

import (
	gogo "github.com/TheThingsIndustries/protoc-gen-go-json/gogo"
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
)

// MarshalProtoJSON marshals the AuthInfoResponse_APIKeyAccess message to JSON.
func (x *AuthInfoResponse_APIKeyAccess) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.ApiKey != nil || s.HasField("api_key") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("api_key")
		x.ApiKey.MarshalProtoJSON(s.WithField("api_key"))
	}
	if x.EntityIds != nil || s.HasField("entity_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("entity_ids")
		// NOTE: EntityIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.EntityIds)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the AuthInfoResponse_APIKeyAccess to JSON.
func (x AuthInfoResponse_APIKeyAccess) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(&x)
}

// UnmarshalProtoJSON unmarshals the AuthInfoResponse_APIKeyAccess message from JSON.
func (x *AuthInfoResponse_APIKeyAccess) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "api_key", "apiKey":
			if s.ReadNil() {
				x.ApiKey = nil
				return
			}
			x.ApiKey = &APIKey{}
			x.ApiKey.UnmarshalProtoJSON(s.WithField("api_key", true))
		case "entity_ids", "entityIds":
			s.AddField("entity_ids")
			if s.ReadNil() {
				x.EntityIds = nil
				return
			}
			// NOTE: EntityIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v EntityIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.EntityIds = &v
		}
	})
}

// UnmarshalJSON unmarshals the AuthInfoResponse_APIKeyAccess from JSON.
func (x *AuthInfoResponse_APIKeyAccess) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the AuthInfoResponse message to JSON.
func (x *AuthInfoResponse) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.AccessMethod != nil {
		switch ov := x.AccessMethod.(type) {
		case *AuthInfoResponse_ApiKey:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("api_key")
			ov.ApiKey.MarshalProtoJSON(s.WithField("api_key"))
		case *AuthInfoResponse_OauthAccessToken:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("oauth_access_token")
			ov.OauthAccessToken.MarshalProtoJSON(s.WithField("oauth_access_token"))
		case *AuthInfoResponse_UserSession:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("user_session")
			// NOTE: UserSession does not seem to implement MarshalProtoJSON.
			gogo.MarshalMessage(s, ov.UserSession)
		}
	}
	if x.UniversalRights != nil || s.HasField("universal_rights") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("universal_rights")
		x.UniversalRights.MarshalProtoJSON(s.WithField("universal_rights"))
	}
	if x.IsAdmin || s.HasField("is_admin") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("is_admin")
		s.WriteBool(x.IsAdmin)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the AuthInfoResponse to JSON.
func (x AuthInfoResponse) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(&x)
}

// UnmarshalProtoJSON unmarshals the AuthInfoResponse message from JSON.
func (x *AuthInfoResponse) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "api_key", "apiKey":
			ov := &AuthInfoResponse_ApiKey{}
			x.AccessMethod = ov
			if s.ReadNil() {
				ov.ApiKey = nil
				return
			}
			ov.ApiKey = &AuthInfoResponse_APIKeyAccess{}
			ov.ApiKey.UnmarshalProtoJSON(s.WithField("api_key", true))
		case "oauth_access_token", "oauthAccessToken":
			ov := &AuthInfoResponse_OauthAccessToken{}
			x.AccessMethod = ov
			if s.ReadNil() {
				ov.OauthAccessToken = nil
				return
			}
			ov.OauthAccessToken = &OAuthAccessToken{}
			ov.OauthAccessToken.UnmarshalProtoJSON(s.WithField("oauth_access_token", true))
		case "user_session", "userSession":
			s.AddField("user_session")
			ov := &AuthInfoResponse_UserSession{}
			x.AccessMethod = ov
			if s.ReadNil() {
				ov.UserSession = nil
				return
			}
			// NOTE: UserSession does not seem to implement UnmarshalProtoJSON.
			var v UserSession
			gogo.UnmarshalMessage(s, &v)
			ov.UserSession = &v
		case "universal_rights", "universalRights":
			if s.ReadNil() {
				x.UniversalRights = nil
				return
			}
			x.UniversalRights = &Rights{}
			x.UniversalRights.UnmarshalProtoJSON(s.WithField("universal_rights", true))
		case "is_admin", "isAdmin":
			s.AddField("is_admin")
			x.IsAdmin = s.ReadBool()
		}
	})
}

// UnmarshalJSON unmarshals the AuthInfoResponse from JSON.
func (x *AuthInfoResponse) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}
