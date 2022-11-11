// This file is automatically generated. Do not modify it manually.

package product

import (
	"encoding/json"
	"strings"

	"github.com/mattermost/mattermost-server/v6/model"
)

var manifest *model.Manifest

const manifestStr = `
{
  "id": "playbooks",
  "name": "Playbooks",
  "description": "Mattermost Playbooks enable reliable and repeatable processes for your teams using checklists, automation, and retrospectives.",
  "homepage_url": "https://github.com/mattermost/mattermost-plugin-playbooks/",
  "support_url": "https://github.com/mattermost/mattermost-plugin-playbooks/issues",
  "release_notes_url": "https://github.com/mattermost/mattermost-plugin-playbooks/releases/tag/v1.33.0+alpha.4",
  "icon_path": "assets/plugin_icon.svg",
  "version": "1.33.0+alpha.4+8f592436",
  "min_server_version": "6.3.0",
  "server": {
    "executables": {
      "darwin-amd64": "server/dist/plugin-darwin-amd64",
      "darwin-arm64": "server/dist/plugin-darwin-arm64",
      "linux-amd64": "server/dist/plugin-linux-amd64",
      "linux-arm64": "server/dist/plugin-linux-arm64",
      "windows-amd64": "server/dist/plugin-windows-amd64.exe"
    },
    "executable": ""
  },
  "webapp": {
    "bundle_path": "webapp/dist/main.js"
  },
  "settings_schema": {
    "header": "",
    "footer": "",
    "settings": [
      {
        "key": "EnableExperimentalFeatures",
        "display_name": "Enable Experimental Features:",
        "type": "bool",
        "help_text": "Enable experimental features that come with in-progress UI, bugs, and cool stuff.",
        "placeholder": "",
        "default": null
      }
    ]
  }
}
`

func init() {
	_ = json.NewDecoder(strings.NewReader(manifestStr)).Decode(&manifest)
}
