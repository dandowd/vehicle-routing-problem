package dispatchers

import "vehicle-routing-problem/entities"

type Dispatch interface {
  SearchForRoutes() []*entities.Driver
}
