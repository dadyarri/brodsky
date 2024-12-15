package resume_json

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

func (plugin *ResumeJsonPlugin) Execute() error {
	return nil
}
