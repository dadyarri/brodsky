package resume_json

import "brodsky/pkg/config"

type ResumeJsonPlugin struct {
}

func (plugin *ResumeJsonPlugin) Name() string {
	return "resume_json"
}

func (plugin *ResumeJsonPlugin) Init(site config.Site) error {
	return nil
}

func (plugin *ResumeJsonPlugin) Execute() error {
	return nil
}
