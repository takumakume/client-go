package dtrack

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Policy struct {
	UUID             uuid.UUID
	Name             string            `json:"name"`
	Operator         string            `json:"operator"`
	ViolationState   string            `json:"violationState"`
	PolicyConditions []PolicyCondition `json:"policyConditions"`
}

type PolicyService struct {
	client *Client
}

func (ps PolicyService) Get(ctx context.Context, policyUUID uuid.UUID) (p Policy, err error) {
	req, err := ps.client.newRequest(ctx, http.MethodGet, fmt.Sprintf("/api/v1/policy/%s", policyUUID))
	if err != nil {
		return
	}

	_, err = ps.client.doRequest(req, &p)
	return
}

func (ps PolicyService) GetAll(ctx context.Context, po PageOptions) (p Page[Policy], err error) {
	req, err := ps.client.newRequest(ctx, http.MethodGet, "/api/v1/policy", withPageOptions(po))
	if err != nil {
		return
	}

	res, err := ps.client.doRequest(req, &p.Items)
	if err != nil {
		return
	}

	p.TotalCount = res.TotalCount
	return
}

func (ps PolicyService) Create(ctx context.Context, policy Policy) (p Policy, err error) {
	req, err := ps.client.newRequest(ctx, http.MethodPut, "/api/v1/policy", withBody(policy))
	if err != nil {
		return
	}

	_, err = ps.client.doRequest(req, &p)
	return
}

func (ps PolicyService) Delete(ctx context.Context, policyUUID uuid.UUID) (err error) {
	req, err := ps.client.newRequest(ctx, http.MethodDelete, fmt.Sprintf("/api/v1/policy/%s", policyUUID))
	if err != nil {
		return
	}

	_, err = ps.client.doRequest(req, nil)
	return
}

func (ps PolicyService) Update(ctx context.Context, policy Policy) (p Policy, err error) {
	req, err := ps.client.newRequest(ctx, http.MethodPost, "/api/v1/policy", withBody(policy))
	if err != nil {
		return
	}

	_, err = ps.client.doRequest(req, &p)
	return
}

func (ps PolicyService) AddProject(ctx context.Context, policyUUID, projectUUID uuid.UUID) (p Policy, err error) {
	req, err := ps.client.newRequest(ctx, http.MethodPost, fmt.Sprintf("/api/v1/policy/%s/project/%s", policyUUID, projectUUID))
	if err != nil {
		return
	}

	_, err = ps.client.doRequest(req, &p)
	return
}

func (ps PolicyService) DeleteProject(ctx context.Context, policyUUID, projectUUID uuid.UUID) (p Policy, err error) {
	req, err := ps.client.newRequest(ctx, http.MethodDelete, fmt.Sprintf("/api/v1/policy/%s/project/%s", policyUUID, projectUUID))
	if err != nil {
		return
	}

	_, err = ps.client.doRequest(req, &p)
	return
}

func (ps PolicyService) AddTag(ctx context.Context, policyUUID uuid.UUID, tagName string) (p Policy, err error) {
	req, err := ps.client.newRequest(ctx, http.MethodPost, fmt.Sprintf("/api/v1/policy/%s/project/%s", policyUUID, tagName))
	if err != nil {
		return
	}

	_, err = ps.client.doRequest(req, &p)
	return
}

func (ps PolicyService) DeleteTag(ctx context.Context, policyUUID uuid.UUID, tagName string) (p Policy, err error) {
	req, err := ps.client.newRequest(ctx, http.MethodDelete, fmt.Sprintf("/api/v1/policy/%s/project/%s", policyUUID, tagName))
	if err != nil {
		return
	}

	_, err = ps.client.doRequest(req, &p)
	return
}
