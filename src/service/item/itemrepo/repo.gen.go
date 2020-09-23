// AUTOGENERATED FILE

// +build !codeanalysis

package itemrepo

import ( 
	"projecto/db"
	"projecto/model"
)

type iItemRepo interface {
    Find() ([]*model.Item, error)
    Create(*model.Item) error
}

func newRepoItem(db *db.Database) *repoItem {
	return &repoItem{
		db: db,
	}
}

type repoItem struct {
	db *db.Database
}

func (r *repoItem) Find() ([]*model.Item, error) {
    records := []*model.Item{}
	if err := r.db.GetDatabase().Find(&records).Error; err != nil {
      return records, err
    }
    return records, nil 
}


func (r *repoItem) Create(record *model.Item) (error) {
	return r.db.GetDatabase().Create(&record).Error
}

