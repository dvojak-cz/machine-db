package service


type MongoManager interface {
	Connect() error
	Disconnect() error
	
}
