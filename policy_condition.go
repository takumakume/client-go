package dtrack

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type PolicyCondition struct {
	Policy   *Policy `json:"policy,omitempty"`
	Operator string  `json:"operator"`
	Subject  string  `json:"subject"`
	Value    string  `json:"value"`
	UUID     uuid.UUID               `json:"uuid,omitempty"`
}

type PolicyConditionService struct {
	client *Client
}

func (pcs PolicyConditionService) Create(ctx context.Context, policyUUID uuid.UUID, policyCondition PolicyCondition) (p PolicyCondition, err error) {
	req, err := pcs.client.newRequest(ctx, http.MethodPut, fmt.Sprintf("/api/v1/policy/%s/condition", policyUUID), withBody(policyCondition))
	if err != nil {
		return
	}

	_, err = pcs.client.doRequest(req, &p)
	return
}

func (pcs PolicyConditionService) Update(ctx context.Context, policyCondition PolicyCondition) (p PolicyCondition, err error) {
	req, err := pcs.client.newRequest(ctx, http.MethodPost, "/api/v1/policy/condition", withBody(policyCondition))
	if err != nil {
		return
	}

	_, err = pcs.client.doRequest(req, &p)
	return
}

func (pcs PolicyConditionService) Delete(ctx context.Context, policyConditionUUID uuid.UUID) (err error) {
	req, err := pcs.client.newRequest(ctx, http.MethodDelete, fmt.Sprintf("/api/v1/policy/condition/%s", policyConditionUUID))
	if err != nil {
		return
	}

	_, err = pcs.client.doRequest(req, nil)
	return
}
