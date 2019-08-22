package model

type Article struct{
  ID        string
  Title     string
  Content   string
  CreatedAt string  `db:"created_at"`
  Modified  string
}
