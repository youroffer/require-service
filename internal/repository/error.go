package repository

import "errors"

// DB
const (
	// foreign key violation: 23503
	FKViolation = "23503"
	// unique violation: 23505
	UniqueConstraint = "23505"
)

// Category
const (
	ErrCategoryNotFound = errors.New("category not found")
	ErrCategoryExists   = errors.New("category exists")
)

// Post
const (
	ErrPostExists             = errors.New("post exists")
	ErrPostNotFound           = errors.New("post not found")
	ErrPostDependencyNotFound = errors.New("cannot add post because there are no record references to category")
)

// Analytic
const (
	ErrPostIDExist                = errors.New("post_id must be unique")
	ErrAnalyticDependencyNotFound = errors.New("cannot add or update analytic because there is no record reference to post")
	ErrAnalyticNotFound           = errors.New("analytic not found")
)

// Filter
const (
	ErrFilterExist    = errors.New("filter already exist")
	ErrFilterNotFound = errors.New("filter not found")
)

// CACHE
const (
	ErrKeyNotFound = errors.New("there is no current data")
)
