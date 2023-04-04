package postgres

import (
	"context"
	"crud/models"
	"testing"
)

// go test -run TestGetByIDActor

func TestCreateActor(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.CreateActor
		Output  string
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.CreateActor{
				First_name: "Asadbek",
				Last_name:  "Ergashev",
			},
			WantErr: false,
		},
		{
			Name: "Case 2",
			Input: &models.CreateActor{
				First_name: "",
				Last_name:  "",
			},
			WantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			id, err := actorTestRepo.Create(context.Background(), test.Input)

			if test.WantErr {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if id == "" {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}
		})
	}
}

func TestGetByIDActor(t *testing.T) {

	tests := []struct {
		Name    string
		Input   *models.ActorPrimarKey
		Output  *models.Actor
		WantErr bool
	}{
		{
			Name: "Case 1",
			Input: &models.ActorPrimarKey{
				Id: "33c151fe-2445-4b89-954a-a0cd715af9e4",
			},
			Output: &models.Actor{
				Id:         "33c151fe-2445-4b89-954a-a0cd715af9e2",
				First_name: "Asadbek",
				Last_name:  "Ergashev",
			},
			WantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {

			actor, err := actorTestRepo.GetByPKey(context.Background(), test.Input)

			// if err == pgx.ErrNoRows {}

			if err != nil {
				t.Errorf("%s: got: %v", test.Name, err)
				return
			}

			if actor.First_name != test.Output.First_name ||
				actor.Last_name != test.Output.Last_name ||
				actor.Id != test.Output.Id {
				t.Errorf("%s: got: %+v, expacted: %+v", test.Name, *actor, *test.Output)
				return
			}
		})
	}
}
