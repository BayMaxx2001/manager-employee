package persistence

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

func (repo *mongoEmployeeRepository) AddEmployeeToTeam(ctx context.Context, employee model.Employee, tid string) error {
	_, err := repo.collection.UpdateOne(
		ctx,
		bson.M{"uid": employee.UID},
		bson.M{"$push": bson.M{"listteams": tid}},
	)

	return err
}

func (repo *mongoEmployeeRepository) DeleteEmployeeToTeam(ctx context.Context, employee model.Employee, tid string) error {
	_, err := repo.collection.UpdateOne(
		ctx,
		bson.M{"uid": employee.UID},
		bson.M{"$pull": bson.M{"listteams": tid}},
	)

	return err
}

func (repo *mongoEmployeeRepository) filterEmployee(ctx context.Context, filter interface{}) ([]model.Employee, error) {
	var Employees []model.Employee

	cur, err := repo.collection.Find(ctx, filter)
	if err != nil {
		return Employees, err
	}

	for cur.Next(ctx) {
		var e model.Employee
		err := cur.Decode(&e)
		if err != nil {
			return Employees, err
		}

		Employees = append(Employees, e)
	}

	if err := cur.Err(); err != nil {
		return Employees, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	if len(Employees) == 0 {
		return Employees, mongo.ErrNoDocuments
	}

	return Employees, nil
}

type employeeDocument struct {
	Name   string `bson:"name"`
	DOB    string `bson:"dob"`
	Gender int    `bson:"gender"`
}

func toEmployeeDocument(e model.Employee) employeeDocument {
	return employeeDocument{
		Name:   e.Name,
		DOB:    e.DOB,
		Gender: e.Gender,
	}
}
