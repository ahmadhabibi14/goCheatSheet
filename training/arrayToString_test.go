package training

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

const (
	InvitationStateInvited    = `invited`
	InvitationStateRevoked    = `revoked`
	InvitationStateAccepted   = `accepted`
	InvitationStateRejected   = `rejected`
	InvitationStateTerminated = `terminated`
	InvitationStateLeft       = `left`
)

type (
	InviteState struct {
		TenantCode string
		State      string
		Date       string
	}
	StateMap map[string]InviteState
)

func (s StateMap) ModifyState(tenantCode, newState string) {
	if sn, ok := s[tenantCode]; ok {
		if sn.State != newState {
			sn.State = newState
			sn.Date = time.Now().GoString()
			s[tenantCode] = sn
		}
	}
	for _, st := range s {
		if newState == InvitationStateAccepted { // Change Accepted to Left
			if st.TenantCode != tenantCode {
				if st.State == InvitationStateAccepted {
					st.State = InvitationStateLeft
					s[st.TenantCode] = st
				} else if st.State == InvitationStateInvited {
					st.State = InvitationStateRevoked
					s[st.TenantCode] = st
				}
			}
		}
	}
}

func ToStateMap(states string) (out StateMap) {
	out = StateMap{}
	statesArray := strings.Split(states, ` `)
	fmt.Println(`Len state = `, len(statesArray))
	for _, state := range statesArray {
		parts := strings.Split(state, `:`)
		out[parts[1]] = InviteState{
			TenantCode: parts[1],
			State:      parts[2],
			Date:       parts[3],
		}
	}
	return
}

func TestArrayToString(t *testing.T) {
	stateINv := `tenant:habi_9123:invited:2023-10-02 tenant:dyan_4823:invited:2023-10-02 tenant:downy_9223:rejected:2023-11-20`
	invs := ToStateMap(stateINv)
	fmt.Printf("Tenant Code: \t\t%s\n", invs[`habi_9123`].TenantCode)
	fmt.Printf("Invitation State: \t%s\n", invs["habi_9123"].State)
	fmt.Printf("Date: \t\t\t%s\n", invs["habi_9123"].Date)

	fmt.Println()

	invs.ModifyState(`dyan_4823`, `accepted`)
	fmt.Printf("Modified: \n")
	for _, v := range invs {
		fmt.Println("======================================")
		fmt.Printf("Tenant Code: \t\t%s\n", v.TenantCode)
		fmt.Printf("Invitation State: \t%s\n", v.State)
		fmt.Printf("Date: \t\t\t%s\n", v.Date)
		fmt.Println("======================================")
	}
}
