package repository

import (
 "errors"
 "Fintregas/domain"
)
type MemoryRepository struct {
	Encomendas map[string]domain.Encomenda
   }
   func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
	 Encomendas: make(map[string]domain.Encomenda),
	}
   }

   func (r *MemoryRepository) AddEncomenda(encomenda domain.Encomenda) error {
	
   
	if _, exists := r.Encomendas[encomenda.ID]; exists {
	 return errors.New("encomenda ja Existe")
	}
   
	r.Encomendas[encomenda.ID] = encomenda
	return nil
   }
   
   func (r *MemoryRepository) GetEncomenda(id string) (domain.Encomenda, error) {
	
   
	encomenda, exists := r.Encomendas[id]
	if !exists {
	 return domain.Encomenda{}, errors.New("ecomenda não encontrada")
	}
   
	return encomenda, nil
   }

   func (r *MemoryRepository) UpdateEncomenda(encomenda domain.Encomenda) error {
	
   
	if _, exists := r.Encomendas[encomenda.ID]; !exists {
	 return errors.New("encomenda não Encontrada")
	}
   
	r.Encomendas[encomenda.ID] = encomenda
	return nil
   }
   
   
   
   
