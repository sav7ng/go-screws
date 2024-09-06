package mongo

import "errors"

type MongoTool struct{}

func (t *MongoTool) IsNoDocuments(err error) bool {
	return err.Error() == "mongo: no documents in result"
}

func (t *MongoTool) IsExist(err error) bool {
	return err.Error() == "existent"
}

func (t *MongoTool) NewExistentError() error {
	return errors.New("existent")
}
