package maps

// Initial payload to identify yourself to discord
func GetIdentifyPayload(authToken string) map[string]interface{} {
	var IdentifyPayload = map[string]interface{}{
		"op": 2,
		"d": map[string]interface{}{
			"token":        authToken,
			"capabilities": 1734653,
			"properties": map[string]interface{}{
				"os":                          "Linux",
				"browser":                     "Chrome",
				"device":                      "",
				"system_locale":               "en-US",
				"has_client_mods":             false,
				"browser_user_agent":          "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36",
				"browser_version":             "138.0.0.0",
				"os_version":                  "",
				"referrer":                    "",
				"referring_domain":            "",
				"referrer_current":            "",
				"referring_domain_current":    "",
				"release_channel":             "stable",
				"client_build_number":         459631,
				"client_event_source":         nil,
				"client_launch_id":            "2a1b19c9-ed33-4a3d-85f9-0a418f0744d8",
				"launch_signature":            "224fefa3-539b-470c-9a18-c20b50bec7d4",
				"client_heartbeat_session_id": "59592a72-ef6e-4550-9974-46fd71fdbaf2",
				"client_app_state":            "unfocused",
				"is_fast_connect":             false,
				"gateway_connect_reasons":     "AppSkeleton",
			},
			"presence": map[string]interface{}{
				"status":     "unknown",
				"since":      0,
				"activities": []interface{}{},
				"afk":        false,
			},
			"compress": false,
			"client_state": map[string]interface{}{
				"guild_versions": map[string]interface{}{},
			},
		},
	}
	return IdentifyPayload
}
