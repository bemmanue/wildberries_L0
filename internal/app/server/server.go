package server

import (
	"github.com/bemmanue/wildberries_L0/internal/cache"
	"github.com/bemmanue/wildberries_L0/internal/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

// server ...
type server struct {
	store  store.Store
	cache  cache.Cache
	router *gin.Engine
}

// newServer ...
func newServer(store store.Store, cache cache.Cache) *server {
	s := &server{
		store:  store,
		cache:  cache,
		router: gin.Default(),
	}

	s.configureRouter()

	return s
}

// configureRouter ...
func (s *server) configureRouter() {
	s.router.Static("/web", "./web")

	s.router.GET("/order", s.getOrder)
	s.router.GET("/order/:order_uid", s.getOrderByUID)
}

// serveHTTP ...
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// getOrder ...
func (s *server) getOrder(c *gin.Context) {
	c.File("./web/templates/order.html")
}

// getOrderByUID ...
func (s *server) getOrderByUID(c *gin.Context) {
	order, err := s.cache.Order().Find(c.Param("order_uid"))
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.Header("Content-Type", "application/json")
	c.String(http.StatusOK, string(order.Data))
}

// Subscribe ...
//func (s *server) Subscribe(c *gin.Context) {
//	var order model.Order
//	err := json.NewDecoder(c.Request.Body).Decode(&order)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
//		return
//	}
//
//	if err := s.validator.Struct(order); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
//		return
//	}
//
//	data, err := json.Marshal(order)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
//		return
//	}
//
//	orderJSON := model.OrderJSON{
//		OrderUID: order.OrderUID,
//		Data:     data,
//	}
//
//	if err := s.store.Order().Create(&orderJSON); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
//		return
//	}
//
//	if err := s.cache.Order().Create(&orderJSON); err != nil {
//		s.cache, _ = mapcache.New(s.store)
//	}
//
//	c.JSON(http.StatusCreated, gin.H{"status": "created"})
//}
