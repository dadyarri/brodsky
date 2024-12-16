package plugins

import (
	"brodsky/pkg/site"
)

type ResumeJsonPlugin struct {
}

func (plugin *ResumeJsonPlugin) Name() string {
	return "resume_json"
}

func (plugin *ResumeJsonPlugin) Init(site site.Site) error {
	return nil
}

func (plugin *ResumeJsonPlugin) Execute(Context) error {
	return nil
}
