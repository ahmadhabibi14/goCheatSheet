package cheatsheet

import (
	"fmt"
	"testing"
)

// InviteState represents the structure for invitation state.
type InviteState struct {
	TenantCode string
	State      string
	Date       string
}

// StateMap is a map with methods for handling InviteState instances.
type StateMap map[string]InviteState

// AddInviteState adds an InviteState to the StateMap.
func (m StateMap) AddInviteState(key string, inviteState InviteState) {
	m[key] = inviteState
}

// GetInviteState retrieves an InviteState from the StateMap.
func (m StateMap) GetInviteState(key string) (InviteState, bool) {
	state, ok := m[key]
	return state, ok
}

// ModifyInviteStateField modifies a specific field of an InviteState in the StateMap.
func (m StateMap) ModifyInviteStateField(key, field, value string) {
	if inviteState, ok := m[key]; ok {
		// Modify the specified field
		switch field {
		case "TenantCode":
			inviteState.TenantCode = value
		case "State":
			inviteState.State = value
		case "Date":
			inviteState.Date = value
		default:
			fmt.Println("Invalid field specified")
			return
		}

		// Store the modified InviteState back into the map
		m[key] = inviteState
	} else {
		fmt.Println("InviteState not found.")
	}
}

func TestMapV4(t *testing.T) {
	// Create a StateMap instance
	stateMap := make(StateMap)

	// Add some InviteState instances to the map
	stateMap.AddInviteState("key1", InviteState{TenantCode: "A", State: "Pending", Date: "2023-01-01"})
	stateMap.AddInviteState("key2", InviteState{TenantCode: "B", State: "Accepted", Date: "2023-01-02"})

	// Print the original map
	fmt.Println("Original StateMap:", stateMap)

	// Modify the "TenantCode" field of an InviteState in the map
	stateMap.ModifyInviteStateField("key1", "TenantCode", "NewTenantCode")

	// Print the updated map
	fmt.Println("Updated StateMap:", stateMap)
}
