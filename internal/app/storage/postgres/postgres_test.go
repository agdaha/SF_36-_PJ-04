package postgres

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"skillfactory/SF_36-_PJ-04/internal/app/model"
	"skillfactory/SF_36-_PJ-04/internal/app/storage"
)

type CustomerRepoTestSuite struct {
	suite.Suite
	pgContainer *PostgresContainer
	store       storage.Store //проверка интерфейса
	ctx         context.Context
}

func (suite *CustomerRepoTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	pgContainer, err := CreatePostgresContainer(suite.ctx)
	if err != nil {
		log.Fatal(err)
	}
	suite.pgContainer = pgContainer

	store, err := New(suite.pgContainer.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	suite.store = store
}

func (suite *CustomerRepoTestSuite) TestAddPosts() {
	t := suite.T()

	err := suite.store.UpdatePosts([]model.Post{
		{
			Title:       "First title",
			Description: "First Description",
			Link:        "link_123",
			PubDate:     111,
			Author:      "Author 1",
		},
		{
			Title:       "Second title",
			Description: "Second Description",
			Link:        "link_122",
			PubDate:     222,
			Author:      "Author 2",
		},
	})
	assert.NoError(t, err)

	posts, err := suite.store.Posts(10)
	assert.NoError(t, err)
	assert.NotNil(t, posts)
	assert.Equal(t, 2, len(posts))

	err = suite.store.UpdatePosts([]model.Post{
		{
			Title:       "Third title",
			Description: "Third Description",
			Link:        "link_125",
			PubDate:     333,
			Author:      "Author 3",
		},
		{
			Title:       "Second title update",
			Description: "Second Description update",
			Link:        "link_122",
			PubDate:     555,
			Author:      "Author 2",
		},
	})

	posts, err = suite.store.Posts(10)
	assert.NoError(t, err)
	assert.NotNil(t, posts)
	assert.Equal(t, 3, len(posts))

	assert.Equal(t, "Second title update", posts[0].Title)
	assert.Equal(t, "Second Description update", posts[0].Description)
	assert.Equal(t, "link_122", posts[0].Link)
	assert.Equal(t, "Author 2", posts[0].Author)
}

func TestCustomerRepoTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerRepoTestSuite))
}
