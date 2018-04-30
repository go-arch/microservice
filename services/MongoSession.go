package services

import (
	"gopkg.in/mgo.v2"
	"log"
	"students/constants"
)



type MongoSession struct {
	Collection *mgo.Collection
	Session *mgo.Session
	Database *mgo.Database
}

func NewMongoSession()  *MongoSession{
	ms := MongoSession{}
	ms.Session = ms.GetSession(constants.URL)
	ms.Database = ms.GetDatabase(ms.Session,constants.DATABASE_NAME)
	ms.Collection = ms.GetCollection(ms.Database,constants.COLLECTION_NAME)
	return &ms
}

func (ms MongoSession) GetSession(url string) *mgo.Session{
	session, err := mgo.Dial(url)
	if err != nil{
		log.Fatal(err)
	}
	return session
}

func(ms MongoSession) GetCollection(db *mgo.Database ,collectionName string) *mgo.Collection{
	return  db.C(collectionName)

}

func(ms MongoSession) GetDatabase(session *mgo.Session,databaseName string) *mgo.Database{
	return session.DB(databaseName)
}

