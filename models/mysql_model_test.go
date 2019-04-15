package models

import (
	"reflect"
	"testing"
)

// Test_sort_reference
// user has one user_profile
// user has many post
// post has many post_category
// post_category has one category
func Test_MysqlModel(t *testing.T) {
	user := Table("user")
	post := Table("post")
	userProfile := Table("user_profile")
	category := Table("category")
	postCategory := Table("post_category")

	tables := Tables{
		post,
		userProfile,
		postCategory,
		category,
		user,
	}

	references := References{
		Reference{
			TableName:           post,
			ReferencedTableName: user,
		},
		Reference{
			TableName:           userProfile,
			ReferencedTableName: user,
		},
		Reference{
			TableName:           postCategory,
			ReferencedTableName: post,
		},
		Reference{
			TableName:           postCategory,
			ReferencedTableName: category,
		},
	}

	dump := &MysqlModel{
		Tables:     tables,
		References: references,
	}
	expected := Tables{user, userProfile, post, category, postCategory}

	t.Run("MysqlModel Sort", func(t *testing.T) {
		dump := &MysqlModel{
			SortedTables: Tables{
				postCategory,
				category,
			},
		}

		tables := Tables{postCategory, category}
		if !dump.existsSorted(tables) {
			t.Fatalf("failed test")
		}

		tables2 := Tables{post, category}
		if dump.existsSorted(tables2) {
			t.Fatalf("failed test")
		}
	})

	t.Run("MysqlModel Sort", func(t *testing.T) {
		dump.Sort()
		if !reflect.DeepEqual(dump.SortedTables, expected) {
			t.Fatalf("failed test %#v . expected %#v", dump.SortedTables, expected)
		}
	})
}
