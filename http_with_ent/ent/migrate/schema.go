// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BlogsColumns holds the columns for the "blogs" table.
	BlogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString, Size: 30},
		{Name: "description", Type: field.TypeString},
		{Name: "episode", Type: field.TypeInt, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_blogs", Type: field.TypeInt, Nullable: true},
	}
	// BlogsTable holds the schema information for the "blogs" table.
	BlogsTable = &schema.Table{
		Name:       "blogs",
		Columns:    BlogsColumns,
		PrimaryKey: []*schema.Column{BlogsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "blogs_users_blogs",
				Columns:    []*schema.Column{BlogsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "type", Type: field.TypeString, Size: 10, Default: "Common"},
		{Name: "category", Type: field.TypeEnum, Nullable: true, Enums: []string{"Hot", "Trending", "Newest", "Controversial"}},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "password", Type: field.TypeString, Default: "QWERTYUIO"},
		{Name: "age", Type: field.TypeInt, Nullable: true, Default: 1},
		{Name: "is_active", Type: field.TypeBool, Default: true},
		{Name: "created_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// TagBlogsColumns holds the columns for the "tag_blogs" table.
	TagBlogsColumns = []*schema.Column{
		{Name: "tag_id", Type: field.TypeInt},
		{Name: "blog_id", Type: field.TypeInt},
	}
	// TagBlogsTable holds the schema information for the "tag_blogs" table.
	TagBlogsTable = &schema.Table{
		Name:       "tag_blogs",
		Columns:    TagBlogsColumns,
		PrimaryKey: []*schema.Column{TagBlogsColumns[0], TagBlogsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tag_blogs_tag_id",
				Columns:    []*schema.Column{TagBlogsColumns[0]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "tag_blogs_blog_id",
				Columns:    []*schema.Column{TagBlogsColumns[1]},
				RefColumns: []*schema.Column{BlogsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserFriendsColumns holds the columns for the "user_friends" table.
	UserFriendsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "friend_id", Type: field.TypeInt},
	}
	// UserFriendsTable holds the schema information for the "user_friends" table.
	UserFriendsTable = &schema.Table{
		Name:       "user_friends",
		Columns:    UserFriendsColumns,
		PrimaryKey: []*schema.Column{UserFriendsColumns[0], UserFriendsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_friends_user_id",
				Columns:    []*schema.Column{UserFriendsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_friends_friend_id",
				Columns:    []*schema.Column{UserFriendsColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BlogsTable,
		TagsTable,
		UsersTable,
		TagBlogsTable,
		UserFriendsTable,
	}
)

func init() {
	BlogsTable.ForeignKeys[0].RefTable = UsersTable
	TagBlogsTable.ForeignKeys[0].RefTable = TagsTable
	TagBlogsTable.ForeignKeys[1].RefTable = BlogsTable
	UserFriendsTable.ForeignKeys[0].RefTable = UsersTable
	UserFriendsTable.ForeignKeys[1].RefTable = UsersTable
}