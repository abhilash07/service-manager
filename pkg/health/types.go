/*
 * Copyright 2018 The Service Manager Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package health

type Status string

const (
	StatusUp      Status = "UP"
	StatusDown           = "DOWN"
	StatusUnknown        = "UNKNOWN"
)

// Health contains information about the health of a component.
type Health struct {
	Status  Status                 `json:"status"`
	Details map[string]interface{} `json:"details,omitempty"`
}

// New returns a new Health with an unknown status an empty details.
func New() *Health {
	return &Health{
		Status:  StatusUnknown,
		Details: make(map[string]interface{}),
	}
}

// WithStatus sets the status of the health
func (h *Health) WithStatus(status Status) *Health {
	h.Status = status
	return h
}

// WithError sets the status of the health to DOWN and adds an error detail
func (h *Health) WithError(err error) *Health {
	h.Status = StatusDown
	return h.WithDetail("error", err)
}

// WithDetail adds a detail to the health
func (h *Health) WithDetail(key string, val interface{}) *Health {
	h.Details[key] = val
	return h
}

// Up sets the health status to up
func (h *Health) Up() *Health {
	h.Status = StatusUp
	return h
}

// Down sets the health status to down
func (h *Health) Down() *Health {
	h.Status = StatusDown
	return h
}

// WithDetails adds the given details to the health
func (h *Health) WithDetails(details map[string]interface{}) *Health {
	for k, v := range details {
		h.Details[k] = v
	}
	return h
}

// Indicator is an interface to provide the health of a component
type Indicator interface {
	// Name returns the name of the component
	Name() string
	// Health returns the health of the component
	Health() *Health
}

// Aggregator is an interface to provide aggregate health information
type Aggregator interface {
	// Aggregate processes the given healths to build a single health
	Aggregate(healths map[string]*Health) *Health
}

// Registry is an interface to store and fetch health indicators
type Registry interface {
	// AddHealthIndicators registers a new health indicator
	AddHealthIndicator(indicator Indicator)
	// HealthIndicators returns the currently registered health indicators
	HealthIndicators() []Indicator
}