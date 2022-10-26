package person

import (
	"github.com/1995parham-teaching/cloud-fall-2022/internal/model"
	"github.com/1995parham-teaching/cloud-fall-2022/internal/store/metric"
	"go.uber.org/zap"
)

type InMemory struct {
	data   map[string]model.Person
	logger *zap.Logger
	usage  metric.Usage
}

func NewInMemory(logger *zap.Logger) *InMemory {
	return &InMemory{
		data:   make(map[string]model.Person),
		logger: logger,
		usage:  metric.NewUsage("person"),
	}
}

func (pim *InMemory) Get() ([]model.Person, error) {
	ps := make([]model.Person, 0)

	for _, p := range pim.data {
		ps = append(ps, p)
	}

	return ps, nil
}

func (pim *InMemory) GetByName(name string) (model.Person, error) {
	p, ok := pim.data[name]
	if !ok {
		return model.Person{}, ErrPersonNotFound
	}

	return p, nil
}

func (pim *InMemory) Save(p model.Person) error {
	pim.logger.Debug("store new person", zap.Any("person", p))

	pim.data[p.Name] = p

	pim.usage.SuccessCount.Add(1)

	return nil
}
