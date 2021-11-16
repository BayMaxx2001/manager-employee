package persistence

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/BayMaxx2001/manager-employee/employee/internal/config"
	"github.com/BayMaxx2001/manager-employee/employee/internal/model"
)

type mongoEmployeeRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func newMongoEmployeeRepository() (repo EmployeeRepository, err error) {
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

	repo = &mongoEmployeeRepository{
		client:     client,
		collection: client.Database(config.Get().Database).Collection(config.Get().Collection),
	}
	return
}

func (repo *mongoEmployeeRepository) FindByUID(ctx context.Context, uid string) (model.Employee, error) {
	result := repo.collection.FindOne(ctx, bson.M{"uid": uid})
	var employee model.Employee
	if err := result.Decode(&employee); err != nil {
		return model.Employee{}, err
	}
	return employee, nil
}

func (repo *mongoEmployeeRepository) Save(ctx context.Context, employee model.Employee) error {
	_, err := repo.collection.InsertOne(ctx, employee)
	return err
}

func (repo *mongoEmployeeRepository) Update(ctx context.Context, uid string, employee model.Employee) error {
	_, err := repo.collection.UpdateOne(
		ctx,
		bson.M{"uid": uid},
		bson.M{
			"$set": toEmployeeDocument(employee),
		},
	)
	return err
}

func (repo *mongoEmployeeRepository) Remove(ctx context.Context, uid string) error {
	_, err := repo.collection.DeleteOne(ctx, bson.M{"uid": uid})
	return err
}

func (repo *mongoEmployeeRepository) GetAll(ctx context.Context) (ls []model.Employee, err error) {
	filter := bson.D{{}}
	return repo.filterEmployee(ctx, filter)
}

func (repo *mongoEmployeeRepository) filterEmployee(ctx context.Context, filter interface{}) ([]model.Employee, error) {
	var lsEmployees []model.Employee

	cur, err := repo.collection.Find(ctx, filter)
	if err != nil {
		return lsEmployees, err
	}

	for cur.Next(ctx) {
		var e model.Employee
		err := cur.Decode(&e)
		if err != nil {
			return lsEmployees, err
		}

		lsEmployees = append(lsEmployees, e)
	}

	if err := cur.Err(); err != nil {
		return lsEmployees, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	if len(lsEmployees) == 0 {
		return lsEmployees, mongo.ErrNoDocuments
	}

	return lsEmployees, nil
}

type employeeDocument struct {
	Name   string             `bson:"name"`
	DOB    primitive.DateTime `bson:"dob"`
	Gender int                `bson:"gender"`
}

func toEmployeeDocument(e model.Employee) employeeDocument {
	return employeeDocument{
		Name:   e.Name,
		DOB:    primitive.NewDateTimeFromTime(e.DobFormat("2006-02-01")),
		Gender: e.Gender,
	}
}
