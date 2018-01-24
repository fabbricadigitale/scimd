package server

import "github.com/gin-gonic/gin"

// Service ...
type Service interface {
	Path() string
}

// Lister ...
type Lister interface {
	Service
	List(*gin.Context)
}

// Poster ...
type Poster interface {
	Service
	Post(*gin.Context)
}

// Getter ...
type Getter interface {
	Service
	Get(*gin.Context)
}

// Putter ...
type Putter interface {
	Service
	Put(*gin.Context)
}

// Patcher ...
type Patcher interface {
	Service
	Patch(*gin.Context)
}

// Deleter ...
type Deleter interface {
	Service
	Delete(*gin.Context)
}

// Searcher ...
type Searcher interface {
	Service
	Search(*gin.Context)
}
