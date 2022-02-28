// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------
package google

import (
	"encoding/json"
	"fmt"
	"time"
)

type TagsOperationWaiter struct {
	Config    *Config
	UserAgent string
	CommonOperationWaiter
}

func (w *TagsOperationWaiter) QueryOp() (interface{}, error) {
	if w == nil {
		return nil, fmt.Errorf("Cannot query operation, it's unset or nil.")
	}
	// Returns the proper get.
	url := fmt.Sprintf("https://cloudresourcemanager.googleapis.com/v3/%s", w.CommonOperationWaiter.Op.Name)

	return sendRequest(w.Config, "GET", "", url, w.UserAgent, nil)
}

func createTagsWaiter(config *Config, op map[string]interface{}, activity, userAgent string) (*TagsOperationWaiter, error) {
	w := &TagsOperationWaiter{
		Config:    config,
		UserAgent: userAgent,
	}
	if err := w.CommonOperationWaiter.SetOp(op); err != nil {
		return nil, err
	}
	return w, nil
}

// nolint: deadcode,unused
func tagsOperationWaitTimeWithResponse(config *Config, op map[string]interface{}, response *map[string]interface{}, activity, userAgent string, timeout time.Duration) error {
	w, err := createTagsWaiter(config, op, activity, userAgent)
	if err != nil {
		return err
	}
	if err := OperationWait(w, activity, timeout, config.PollInterval); err != nil {
		return err
	}
	return json.Unmarshal([]byte(w.CommonOperationWaiter.Op.Response), response)
}

func tagsOperationWaitTime(config *Config, op map[string]interface{}, activity, userAgent string, timeout time.Duration) error {
	if val, ok := op["name"]; !ok || val == "" {
		// This was a synchronous call - there is no operation to wait for.
		return nil
	}
	w, err := createTagsWaiter(config, op, activity, userAgent)
	if err != nil {
		// If w is nil, the op was synchronous.
		return err
	}
	return OperationWait(w, activity, timeout, config.PollInterval)
}
