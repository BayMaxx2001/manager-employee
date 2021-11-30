package persistence

import (
	"context"
	"fmt"
	"time"

	"github.com/BayMaxx2001/manager-employee/team/internal/config"
	"github.com/BayMaxx2001/manager-employee/team/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoTeamRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func newMongoTeamRepository() (repo TeamRepository, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Get().MongoDbUrl))
	if err != nil {
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return
	}

	repo = &mongoTeamRepository{
		client:     client,
		collection: client.Database(config.Get().Database).Collection(config.Get().Collection),
	}
	return
}

func (repo *mongoTeamRepository) FindByUID(ctx context.Context, uid string) (model.Team, error) {
	result := repo.collection.FindOne(ctx, bson.M{"uid": uid})

	var team model.Team
	if err := result.Decode(&team); err != nil {
		return model.Team{}, err
	}
	return team, nil
}

func (repo *mongoTeamRepository) Save(ctx context.Context, team model.Team) error {
	_, err := repo.collection.InsertOne(ctx, team)
	return err
}

func (repo *mongoTeamRepository) Update(ctx context.Context, uid string, team model.Team) error {
	_, err := repo.collection.UpdateOne(
		ctx,
		bson.M{"uid": uid},
		bson.M{
			"$set": toTeamDocument(team),
		},
	)
	return err
}

func (repo *mongoTeamRepository) Remove(ctx context.Context, uid string) error {
	_, err := repo.collection.DeleteOne(ctx, bson.M{"uid": uid})
	return err
}

func (repo *mongoTeamRepository) GetAll(ctx context.Context) (ls []model.Team, err error) {
	filter := bson.D{{}}
	return repo.filterTeam(ctx, filter)
}

func (repo *mongoTeamRepository) AddTeamToEmplopyee(ctx context.Context, team model.Team, eid string) error {
	fmt.Println(eid, "addTeamToemployee")
	_, err := repo.collection.UpdateOne(
		ctx,
		bson.M{"uid": team.UID},
		bson.M{"$push": bson.M{"listemployees": eid}},
	)
	return err
}

func (repo *mongoTeamRepository) DeleteTeamToEmployee(ctx context.Context, team model.Team, eid string) error {
	_, err := repo.collection.UpdateOne(
		ctx,
		bson.M{"uid": team.UID},
		bson.M{"$pull": bson.M{"listemployees": eid}},
	)
	return err

}
func (repo *mongoTeamRepository) filterTeam(ctx context.Context, filter interface{}) ([]model.Team, error) {
	var lsTeams []model.Team

	cur, err := repo.collection.Find(ctx, filter)
	if err != nil {
		return lsTeams, err
	}

	for cur.Next(ctx) {
		var e model.Team
		err := cur.Decode(&e)
		if err != nil {
			return lsTeams, err
		}
		lsTeams = append(lsTeams, e)
	}

	if err := cur.Err(); err != nil {
		return lsTeams, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)
	if len(lsTeams) == 0 {
		return lsTeams, mongo.ErrNoDocuments
	}

	return lsTeams, nil
}

type teamDocument struct {
	Name string `bson:"name"`
}

func toTeamDocument(e model.Team) teamDocument {
	return teamDocument{
		Name: e.Name,
	}
}
