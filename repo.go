package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"time"
)

var currentId int

var agents Agents

// Give us some seed data
func init() {
	u1 := uuid.NewV4().String()
	u2 := uuid.NewV4().String()

	RepoCreateAgent(Agent{ID: 0, UUID: u1, OS: "Windows 7 Pro", Location: "china", CreatedDate: time.Now()})
	RepoCreateAgent(Agent{ID: 1, UUID: u2, OS: "Ubuntu Linux", Location: "russia", CreatedDate: time.Now()})
}

func RepoFindAgentUUID(uuid string) Agent {
	for _, t := range agents {
		if t.UUID == uuid {
			return t
		}
	}
	return Agent{}
}

func RepoFindAgent(id int) Agent {
	for _, t := range agents {
		if t.ID == id {
			return t
		}
	}
	// return empty Todo if not found
	return Agent{}
}

func RepoCreateAgent(t Agent) Agent {
	currentId += 1
	t.ID = currentId
	t.CreatedDate = time.Now()
	agents = append(agents, t)
	return t
}

func RepoDestroyAgent(id int) error {
	for i, t := range agents {
		if t.ID == id {
			agents = append(agents[:i], agents[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Agent with id of %d to delete", id)
}
